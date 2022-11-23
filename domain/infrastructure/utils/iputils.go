package utils

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

func GetIpLocation(ip string) (country string, regiom string, city string, err error) {
	type ipapiResponse struct {
		XMLName xml.Name `xml:"root"`
		City    string   `xml:"city"`
		Region  string   `xml:"region"`
		Country string   `xml:"country_name"`
	}

	resp, err := http.Get("https://ipapi.co/xml")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	xmlBytes, err := ioutil.ReadAll(resp.Body)
	var data ipapiResponse
	xml.Unmarshal(xmlBytes, &data)
	log.Println("Defining location of ip")
	log.Println(string(xmlBytes))
	log.Println(data)
	return data.Country, data.Region, data.City, nil

}
