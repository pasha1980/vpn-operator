package job

import (
	"github.com/digitalocean/godo"
	"github.com/go-ping/ping"
	"time"
	"vpn-operator/config"
	"vpn-operator/domain/operator/model"
)

var doClient *godo.Client
var doDroplets []doDroplet

type doDroplet struct {
	Droplet  godo.Droplet
	Instance model.Instance
}

func (d *doDroplet) IsAvailable() bool {
	pinger, _ := ping.NewPinger(d.Instance.IP)
	pinger.Timeout = 5 * time.Second
	pinger.Run()

	stat := pinger.Statistics()
	if stat.PacketsRecv > 0 {
		return true
	}

	return false
}

func (d *doDroplet) Restart() error {
	actionService := doClient.DropletActions

	_, _, err := actionService.Reboot(config.Context, d.Droplet.ID)
	if err != nil {
		_, _, err = actionService.PowerCycle(config.Context, d.Droplet.ID)
		if err != nil {
			_, _, err = actionService.PowerOff(config.Context, d.Droplet.ID)
			_, _, err = actionService.PowerOn(config.Context, d.Droplet.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func InitDigitalOceanHealthCheck() {
	doClient = godo.NewFromToken(config.Config.DigitalOceanToken)

	for range time.Tick(time.Minute) {
		actualizeDropletList()
		dropletsHealthCheck()
	}
}

func actualizeDropletList() {
	doDroplets = make([]doDroplet, 0)

	responseDroplets, _, err := doClient.Droplets.List(config.Context, nil)
	if err != nil {
		return
	}

	for _, droplet := range responseDroplets {
		if !isDropletSupportHealthCheck(droplet) {
			continue
		}

		ip, _ := droplet.PublicIPv4()

		var instance model.Instance
		err = config.DB.Where(model.Instance{
			IP: ip,
		}).First(&instance).Error
		if err != nil {
			continue
		}

		doDroplets = append(
			doDroplets,
			doDroplet{
				Droplet:  droplet,
				Instance: instance,
			},
		)
	}
}

func isDropletSupportHealthCheck(droplet godo.Droplet) bool {
	for _, tag := range droplet.Tags {
		if tag == config.Config.DigitalOceanTag {
			return true
		}
	}

	return false
}

func dropletsHealthCheck() {
	for _, droplet := range doDroplets {
		go func(droplet doDroplet) {
			if !droplet.IsAvailable() {
				config.Log.Write("DO droplet #"+droplet.Droplet.Name+" not available. Restarting...", "digital ocean")
				err := droplet.Restart()
				if err != nil {
					instance := droplet.Instance
					instance.IsActive = false
					config.DB.Save(&instance)
				}
			}
		}(droplet)
	}
}
