package main

import (
	"golang.org/x/oauth2"
	"log"
	"orquestador/basic/config"
	"orquestador/basic/server"
)

func main() {
	var oAuthScopes []string
	appConfig := config.New("")
	appConfig.OAuth2Config = &oauth2.Config{
		ClientID:     appConfig.ClientID,
		ClientSecret: appConfig.ClientSecret,
		Scopes:       oAuthScopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  appConfig.AuthURL,
			TokenURL: appConfig.TokenURL,
		},
	}
	s := server.New(appConfig)
	err := s.Start()
	if err != nil {
		log.Fatalln("An error with the web s occurred:", err)
	}
}
