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
	//var oAuthScopes []string
	appConfig := config.New("")
	appConfig.Services = services
	//appConfig.OAuth2Config = &oauth2.Config{
	//	ClientID:     appConfig.ClientID,
	//	ClientSecret: appConfig.ClientSecret,
	//	Scopes:       oAuthScopes,
	//	Endpoint: oauth2.Endpoint{
	//		AuthURL:  appConfig.AuthURL,
	//		TokenURL: appConfig.TokenURL,
	//	},
	//}
	s := server.New(appConfig)
	err = s.Start()
	if err != nil {
		log.Fatalln("An error with the web s occurred:", err)
	}
}
