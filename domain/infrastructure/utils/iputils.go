package utils

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func GetIpLocation(ip string) (country string, regiom string, city string, err error) {
	type ipapiResponse struct {
		XMLName xml.Name `xml:"root"`
		City    string   `xml:"city"`
		Region  string   `xml:"region"`
		Country string   `xml:"country_name"`
	}

	ipapiClient := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://ipapi.co/"+ip+"/xml/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := ipapiClient.Do(req)
	defer resp.Body.Close()

	xmlBytes, _ := ioutil.ReadAll(resp.Body)

	var data ipapiResponse
	xml.Unmarshal(xmlBytes, &data)
	return data.Country, data.Region, data.City, nil

}
