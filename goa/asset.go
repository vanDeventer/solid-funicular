// The asset is the nuleous of a system while the shell is its interface to the world.
package goa

import (
	"encoding/json"
	"log"
	"net/http"
)

// An asset's capabilities are exposed (and registered) as micro-services through the Arrowhead framework
type Asset struct {
	AssetName string `json:"assetName"`
}

// Default asset configuration used when populating the systemconfig.json file
func AssetDefaultConfig(asset *Asset) {
	asset.AssetName = "serviceRegistry"
}

// Function Current Time responds to a service request
func Register(w http.ResponseWriter, re1 *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
