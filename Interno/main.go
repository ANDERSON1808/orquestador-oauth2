package main

import (
	"log"
	"orquestador/basic/config"
	"orquestador/basic/server"
)

func main() {
	services, err := server.GetServices()
	if err != nil {
		log.Fatalln("An erro whit file service", err)
	}
	appConfig := config.New("")
	appConfig.Services = services
	s := server.New(appConfig)
	err = s.Start()
	if err != nil {
		log.Fatalln("An error with the web s occurred:", err)
	}
}
