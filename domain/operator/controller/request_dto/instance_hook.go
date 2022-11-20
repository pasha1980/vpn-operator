package request_dto

type InstanceHook struct {
	Action            string   `json:"action"`
	URL               string   `json:"url"`
	AvailableServices []string `json:"availableServices"`
	Secret            *string  `json:"secret"`
	Version           *string  `json:"version"`
	Country           *string  `json:"country"`
	Region            *string  `json:"region"`
}
