// This file contains the tools to configure the system's components
package goa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type configurationPayload struct {
	Device      Device    `json:"device"`
	Asset       Asset     `json:"asset"`
	System      System    `json:"system"`
	Services    []Service `json:"services"`
	CoreSystems []System  `json:"coreSystems"`
}

// read (or create a default) system configuration file
func Configure(device *Device, asset *Asset, system *System, services *[]Service, coreSystems *[]System) {
	systemconfigfile, errOpen := os.Open("systemconfig.json")
	// if we os.Open returns an error then handle it
	if errOpen != nil { // could not find the systemconfig.json so a default one is being created
		log.Println(errOpen)

		deviceDefaultConfig(device)
		AssetDefaultConfig(asset)
		systemDefaultConfig(system, asset.AssetName, device.NetworkInterfaces.Ipv4[0])
		serviesDefaultConfig(services)
		coreSystemsDefaultConfig(coreSystems)
		payload := configurationPayload{*device, *asset, *system, *services, *coreSystems}

		systemconfigjson, errMarshal := json.MarshalIndent(payload, "", " ")
		if errMarshal != nil {
			log.Println(errMarshal)
		}

		systemconfigfile, err := os.Create("systemconfig.json")
		if err != nil {
			log.Fatalln(err)
		}
		defer systemconfigfile.Close()

		nbytes, errWrite := systemconfigfile.Write(systemconfigjson)
		if errWrite != nil {
			log.Fatalln(errWrite)
		}
		log.Printf("wrote %d bytes\n", nbytes)
	} else { // managed to open the existing systemconfig.json file
		defer systemconfigfile.Close()
		configBytes, errRead := ioutil.ReadAll(systemconfigfile)
		if errRead != nil {
			log.Fatalln(errRead)
		}

		var payload configurationPayload
		// extract device configuration
		errUnmarshal := json.Unmarshal(configBytes, &payload)
		if errUnmarshal != nil {
			log.Fatalln(errUnmarshal)
		} else {
			*device = payload.Device
			*asset = payload.Asset
			*system = payload.System
			*services = payload.Services
			*coreSystems = payload.CoreSystems
		}
	}

}

// Default device configuration used to populate the systemconfig.json file
func deviceDefaultConfig(device *Device) {
	device.DeviceName = "Laptop"
	device.NetworkInterfaces.Ipv4 = []string{"127.0.0.1", "localhost"}
	device.NetworkInterfaces.Ipv6 = []string{"::1", "localhost"}
}

// Default system configuration used to populate the systemconfig.json file
func systemDefaultConfig(system *System, name string, address string) {
	system.SystemName = name
	system.Address = address
	system.Port = 7879
	system.Authenication = ".X509pubKey"
	system.Protocol = append(system.Protocol, "HTTP")
}

// Default array of services configuration used to populate the systemconfig.json file
func serviesDefaultConfig(Services *[]Service) {
	*Services = append(*Services, Service{
		ServiceDefinition: "Time",
		ServiceName:       "timeService",
		Path:              "/time",
		Metadata:          map[string]string{"Location": "Luleå"},
		Version:           "0.1",
	})
}

// Default array of services configuration used to populate the systemconfig.json file
func coreSystemsDefaultConfig(coreSys *[]System) {
	*coreSys = append(*coreSys, System{
		SystemName:    "ServiceRegistry",
		Address:       "http://localhost:4245/serviceregistry/register",
		Port:          4245,
		Authenication: ".X509pubKey",
		Protocol:      []string{"HTTP"},
	})
}
