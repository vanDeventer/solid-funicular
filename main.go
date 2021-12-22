// This is a Go implementation of the Eclipse Arrowhead framework that has a service oriented architecture (SOA).
// The framework relies on a system of systems called a local cloud where systems offer and consume services from each other.
// Provided services are reregistered for that local cloud in a core system called the Service Registry.
// To obtain the location of a service, a consumer system need to query the Orchestrator system of the local cloud.
//
// The Go implementation strives to have a common structural pattern for all systems.
// A system is made up of a system shell (this application) and an asset (e.g., sensor, actuator, database, PLC, algorithm)

package main

import (
	"arrowhead/goa"
	"net/http"
	"time"
)

// The main function setup the different components of the system shell, and starts the clients and servers
func main() {
	// Define the different components that comprise an Arrowhead system
	// Define = allocate memory and initialize or configure the components of the system

	// Allocate memory
	var device goa.Device
	var asset goa.Asset
	var system goa.System
	var services []goa.Service
	var coreSystems []goa.System

	// Configure the system's components (this allows the deployment technician to configure the system at installation)
	goa.Configure(&device, &asset, &system, &services, &coreSystems)

	// Let the asset start to consume services, starting with the service registration, which is repeated every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			goa.RegisterService(goa.FillRegistrationForm(asset, system, services[0]))
		}
	}()

	// Set up request responder and start server
	http.HandleFunc("/clock/time", goa.CurrTime)
	http.ListenAndServe(":3560", nil)
}
