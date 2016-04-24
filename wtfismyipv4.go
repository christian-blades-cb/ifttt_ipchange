package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
)

const MY_FING_IPV4_URL = "https://ipv4.wtfismyip.com/json"

type FingIpV4 struct {
	IPAddress string `json:"YourFuckingIPAddress"`
	Location  string `json:"YourFuckingLocation"`
	HostName  string `json:"YourFuckingHostname"`
	Isp       string `json:"YourFuckingISP"`
}

func GetIpV4(ctx context.Context, client *http.Client, url string) (ipv4 FingIpV4, err error) {
	resp, err := ctxhttp.Get(ctx, client, url)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	err = decoder.Decode(&ipv4)

	return
}
