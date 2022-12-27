package server

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"orquestador/basic/config"
	"orquestador/basic/models"
)

func (s *Server) getAllProducts(w http.ResponseWriter, req *http.Request) {
	var listProduct []*models.Products

	for _, item := range s.config.Services {
		var oAuthScopes []string
		s.config.OAuth2Config = &oauth2.Config{
			ClientID:     item.ClientId,
			ClientSecret: item.ClientSecret,
			Scopes:       oAuthScopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  item.AuthUrl,
				TokenURL: item.TokenUrl,
			},
		}
		//Autenticacion
		s.handleOAuthCallback(w, req)
		//Solicitud productos
		invoicesRequest, err := http.NewRequest("GET", fmt.Sprintf("%s?client_id=%s&response_type=token", item.UrlProduct, item.ClientId), nil)
		if err != nil {
			errMsg := "An error occurred while trying to create a request to send to the API product."
			_, _ = w.Write([]byte("<p>" + errMsg + "</p>"))
			_, _ = w.Write([]byte("<p>Please check the log for more details.</p>"))
			log.Println(errMsg)
			log.Println(err)
			return
		}
		//?client_id=000000&response_type=token
		invoicesRequest.Header.Add(s.getAuthorisationHeader())
		invoicesResponse, err := s.httpClient.Do(invoicesRequest)
		//Validacion respuesta
		if err != nil && invoicesResponse.StatusCode != 200 {
			errMsg := "An error occurred while trying to retrieve the products."
			w.Write([]byte(errMsg + " Please check the log for details"))
			log.Println(errMsg)
			log.Println(err)
			return
		}
		respBody, err := ioutil.ReadAll(invoicesResponse.Body)
		if err != nil {
			errMsg := "An error occurred while trying to read the response from the Product"
			_, err := w.Write([]byte(errMsg + " Please check the log for details"))
			log.Println(errMsg)
			log.Fatalln(err)
			return
		}
		//Construir productos respuesta
		products, err := DecodeProducts(respBody)
		if err != nil {
			return
		}
		//Construir respuesta final
		listProduct = append(listProduct, models.NewProducts(item.NameCompany, products))
	}
	//Respuesta general de todos los productos
	crearJson, _ := json.Marshal(models.NewResponseProduct(listProduct))
	_, err := w.Write(crearJson)
	w.WriteHeader(200)
	if err != nil {
		return
	}
}

func DecodeProducts(respBody []byte) (lista []models.Lista, err error) {
	err = json.Unmarshal(respBody, &lista)
	if err != nil {
		errMsg := "There was an error attempting to unmarshal the data from the organisation endpoint."
		log.Println(errMsg)
		log.Println(err)
		if config.DebugMode {
			log.Println("Response Body:")
			fmt.Println(string(respBody))
		}
		return
	}
	return
}
