package controllers

import (
	"os"

	"github.com/ShadowApex/meraki-go-sdk/meraki"
)

const header = "X-Cisco-Meraki-API-Key"

var merakiAPI *meraki.APIClient
var apiKey = os.Getenv("MERAKI_API_KEY")
var debugStr = os.Getenv("DEBUG")
var debug = debugStr == "1"

func init() {
	// Build the meraki config
	config := meraki.NewConfiguration()
	config.Host = "api.meraki.com"
	config.Debug = debug
	config.Scheme = "https"
	config.DefaultHeader[header] = apiKey

	// Build the meraki API client
	merakiAPI = meraki.NewAPIClient(config)
}
