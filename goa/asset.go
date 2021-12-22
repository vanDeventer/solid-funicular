// The asset is the nuleous of a system while the shell is its interface to the world.
package goa

import (
	"fmt"
	"net/http"
	"time"
)

// An asset's capabilities are exposed (and registered) as micro-services through the Arrowhead framework
type Asset struct {
	AssetName string `json:"assetName"`
}

// Default asset configuration used when populating the systemconfig.json file
func AssetDefaultConfig(asset *Asset) {
	asset.AssetName = "clock"
}

// Function Current Time responds to a service request
func CurrTime(w http.ResponseWriter, re1 *http.Request) {
	dt := time.Now()
	_, x1 := fmt.Fprintf(w, dt.String())
	if x1 != nil {
		fmt.Println(x1)
	}
}
