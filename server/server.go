package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"orquestador/basic/config"
	"os"
	"strconv"

	"golang.org/x/oauth2"
)

type Server struct {
	config                 *config.Config
	context                context.Context
	httpServer             *http.Server
	httpClient             http.Client
	oAuthAuthorisationCode string
	oAuthToken             *oauth2.Token
}

const callbackURI = "/product"

func New(c *config.Config) *Server {
	defaultAppPort := 8080

	if c.AppPort == 0 {
		if envAppPort := os.Getenv("APP_PORT"); envAppPort != "" {
			var err error
			c.AppPort, err = strconv.Atoi(envAppPort)
			if err != nil {
				log.Fatalln("An error occurred while trying to read the APP_PORT environment variable:", err)
			}
		} else {
			c.AppPort = defaultAppPort
		}
	}

	c.OAuth2Config.RedirectURL = fmt.Sprintf("http://localhost:%d%s", c.AppPort, callbackURI)
	if config.DebugMode {
		log.Println("RedirectURL:", c.OAuth2Config.RedirectURL)
	}

	s := &Server{
		config:     c,
		context:    context.Background(),
		httpServer: &http.Server{Addr: fmt.Sprintf(":%d", c.AppPort)},
	}

	http.HandleFunc(callbackURI, s.getAllProducts)

	return s
}

func (s *Server) Start() error {
	log.Printf("Hey there! I'm up and running, and can be accessed at: http://localhost:%d\n", s.config.AppPort)
	return s.httpServer.ListenAndServe()
}

func (s *Server) getAuthorisationHeader() (string, string) {
	return "authorization", base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("Basic %s:%s", s.config.ClientID, s.config.ClientSecret),
	))
}

func (s *Server) redirectToAuthorisationEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Location", s.config.OAuth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline))
	w.WriteHeader(http.StatusTemporaryRedirect)
}
