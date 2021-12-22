// The registration.go address the registration use case.
// https://github.com/eclipse-arrowhead/core-java-spring#service-registry

package goa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// The Register Service function post a registration from to the Service Registry system to register one service
func RegisterService(registrationForm RegistrationRequest) {
	payload, errMarshal := json.MarshalIndent(registrationForm, "", " ")
	if errMarshal != nil {
		fmt.Println("Registration marshall error")
	}
	fmt.Printf("%+v\n", string(payload))
	serviceRegistryURL := "http://127.0.0.1:4243/serviceregistry/register"
	request, error := http.NewRequest("POST", serviceRegistryURL, bytes.NewBuffer(payload))
	if error != nil {
		fmt.Println(error)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		//panic(error)
		fmt.Println(error)
	} else {
		defer response.Body.Close()

		fmt.Println("response Status:", response.Status)
		fmt.Println("response Headers:", response.Header)
		body, errRead := ioutil.ReadAll(response.Body)
		if errRead != nil {
			fmt.Print("error in reading registration payolad")
		}
		fmt.Println("response Body:", string(body))
		// need to unmarshall the resonse with RegistrationReply
	}

}

// The Fill Registration Form function fills out the form structure that is used when registering a service
func FillRegistrationForm(asset Asset, system System, service Service) RegistrationRequest {
	var form RegistrationRequest
	form.ServiceDefinition = service.ServiceDefinition
	form.ProviderSystem.SystemName = system.SystemName
	form.ProviderSystem.Address = system.Address
	form.ProviderSystem.Port = system.Port
	form.ProviderSystem.AuthenticationInfo = system.Authenication
	form.ServiceURI = "http://" + system.Address + ":" + strconv.Itoa(system.Port) + "/" + system.SystemName + service.Path
	form.EndOfValidity = "tomorrow"
	form.Secure = "INSECURE"
	form.Metadata = service.Metadata
	form.Version = service.Version
	form.Interfaces = system.Protocol
	return form
}

type RegistrationRequest struct {
	ServiceDefinition string `json:"serviceDefinition"`
	ProviderSystem    struct {
		SystemName         string `json:"systemName"`
		Address            string `json:"address"`
		Port               int    `json:"port"`
		AuthenticationInfo string `json:"authenticationInfo"`
	} `json:"providerSystem"`
	ServiceURI    string            `json:"serviceUri"`
	EndOfValidity string            `json:"endOfValidity"`
	Secure        string            `json:"secure"`
	Metadata      map[string]string `json:"metadata"`
	Version       string            `json:"version"`
	Interfaces    []string          `json:"interfaces"`
}

type RegistrationReply struct {
	ID                int `json:"id"`
	ServiceDefinition struct {
		ID                int    `json:"id"`
		ServiceDefinition string `json:"serviceDefinition"`
		CreatedAt         string `json:"createdAt"`
		UpdatedAt         string `json:"updatedAt"`
	} `json:"serviceDefinition"`
	Provider struct {
		ID                 int    `json:"id"`
		SystemName         string `json:"systemName"`
		Address            string `json:"address"`
		Port               int    `json:"port"`
		AuthenticationInfo string `json:"authenticationInfo"`
		CreatedAt          string `json:"createdAt"`
		UpdatedAt          string `json:"updatedAt"`
	} `json:"provider"`
	ServiceURI    string `json:"serviceUri"`
	EndOfValidity string `json:"endOfValidity"`
	Secure        string `json:"secure"`
	Metadata      struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"metadata"`
	Version    int `json:"version"`
	Interfaces []struct {
		ID            int    `json:"id"`
		InterfaceName string `json:"interfaceName"`
		CreatedAt     string `json:"createdAt"`
		UpdatedAt     string `json:"updatedAt"`
	} `json:"interfaces"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
