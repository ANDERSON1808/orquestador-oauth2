package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Services struct {
	AuthUrl      string `json:"auth_url"`
	TokenUrl     string `json:"token_url"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	NameCompany  string `json:"name_company"`
	UrlProduct   string `json:"url_product"`
}

func (t Services) ToString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}
