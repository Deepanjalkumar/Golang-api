package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type All struct {
	Passive_dns []Domain `json:"Passive_dns"`
	Count       int      `json:"Count"`
}

type Domain struct {
	Address        string `json:"Address"`
	First          string `json:"First"`
	Last           string `json:"Last"`
	Hostname       string `json:"Hostname"`
	Record_type    string `json:"Record_type"`
	Indicator_link string `json:"Indicator_link"`
	Flag_url       string `json:"Flag_url"`
	Flag_title     string `json:"Flag_title"`
	Asset_type     string `json:"Asset_type"`
	Asn            string `json:"Asn"`
}

func Alienvault_enum() []Domain {
	r, err := http.Get("https://otx.alienvault.com/api/v1/indicators/domain/owasp.org/passive_dns")

	if err != nil {
		fmt.Println(err)
	}

	file_byte, _ := ioutil.ReadAll(r.Body)

	var Src All

	json.Unmarshal(file_byte, &Src)

	return Src.Passive_dns
}
