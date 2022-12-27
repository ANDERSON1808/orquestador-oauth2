package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"orquestador/basic/models"
	"os"
)

func GetServices() (services []*models.Services, err error) {
	services = make([]*models.Services, 3)
	raw, err := ioutil.ReadFile("./configClients/integraciones.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(raw, &services)
	if err != nil {
		return
	}
	return
}
