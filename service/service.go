package service

import (
	b64 "encoding/base64"

	"github.com/go-resty/resty/v2"
)

type client struct {
	c *resty.Client
}

func newService() *client {
	return &client{c: resty.New()}
}

// ยังไม่ได้เทสเพราะไม่มีเบอร์ DTAC *-*
func (c *client) isDTAC(number string) bool {
	resp, err := c.c.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Cookie", `v10DeviceID=be82b083f640c6fbfbac70448d59ea54; v10UA=DTAC%2F10.0%2Fweb%2F1.0%2FNetworkType%2Fbe82b083f640c6fbfbac70448d59ea54%2FDesktop-Edge; fromApp=dtacLite; v10Lang=th; _ga=GA1.3.1027016751.1646647562; _gid=GA1.3.1799594427.1647101862; _gat_UA-16732483-1=1; OptanonConsent=isIABGlobal=false&datestamp=Sat+Mar+12+2022+23%3A17%3A43+GMT%2B0700+(Indochina+Time)&version=6.15.0&hosts=&consentId=0b8892be-fff8-41a5-a922-c093fb7ea4d0&interactionCount=1&landingPath=https%3A%2F%2Fapp.dtac.co.th%2Fhome&groups=C0001%3A1%2CC0003%3A1%2CC0002%3A1%2CC0004%3A1; keepMeSignIn=true`).
		SetBody(map[string]interface{}{"msisdn": number}).
		Post("https://app.dtac.co.th/api/auth/local-auth/otp?lang=th&uid=guest")

	if err != nil {
		return false
	}

	if resp.StatusCode() != 200 {
		return false
	}

	if resp.StatusCode() == 200 {
		return true
	}
	return false
}

func (c *client) isTRUE(number string) bool {
	resp, err := c.c.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("api-version", "3").
		SetAuthToken("5aaf9ade15afe0324400bacc83067c9af5664822aaa0a739d528528b").
		SetBody(map[string]interface{}{"id": b64.StdEncoding.EncodeToString([]byte(number))}).
		Post("https://dsmapi.truecorp.co.th/iservice-subscriber/api/products/type")

	if err != nil {
		return false
	}

	if resp.StatusCode() != 200 {
		return false
	}

	if resp.StatusCode() == 200 {
		return true
	}
	return false
}

func (c *client) isAIS(number string) bool {
	resp, err := c.c.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"number": number}).
		Post("https://myais.ais.co.th/auth/mobile/" + number + "/verifyMyAIS")

	if err != nil {
		return false
	}

	if resp.StatusCode() != 200 {
		return false
	}

	if resp.StatusCode() == 200 {
		return true
	}
	return false

}

func GetIsp(number string) (service string) {

	client := newService()

	if client.isAIS(number) {
		return "AIS"
	} else if client.isTRUE(number) {
		return "TRUE"
	} else if client.isDTAC(number) {
		return "DTAC"
	} else {
		return "UNKNOWN SERVICE"
	}
}
