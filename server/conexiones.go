package server

import (
	"log"
	"net/http"
	"orquestador/basic/config"

	"golang.org/x/oauth2"
)

func (s *Server) handleOAuthCallback(w http.ResponseWriter, req *http.Request) {
	s.oAuthAuthorisationCode = req.URL.Query().Get("code")
	if config.DebugMode {
		log.Println("Received authorisation code:", s.oAuthAuthorisationCode)
	}
	tok, err := s.config.OAuth2Config.Exchange(
		s.context,
		s.oAuthAuthorisationCode,
		oauth2.SetAuthURLParam(s.getAuthorisationHeader()),
		oauth2.SetAuthURLParam("grant_type", "client_credentials"),
	)

	if err != nil {
		log.Println("An error occurred while trying to exchange the authorisation code with the Xero API.")
		log.Fatalln(err)
	}
	s.oAuthToken = tok
	if config.DebugMode {
		log.Println("Got OAuth2 Token from API.")
		log.Println("Token expiry:", tok.Expiry.String())
	}
	s.httpClient = *s.config.OAuth2Config.Client(s.context, tok)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) refreshAccessToken() error {
	// We create a new token source that only has the refresh token, to force the OAuth2 client to retrieve a new access
	// token.
	src := s.config.OAuth2Config.TokenSource(s.context, &oauth2.Token{RefreshToken: s.oAuthToken.RefreshToken})
	newToken, err := src.Token()
	if err != nil {
		return err
	}
	// Also update the Server struct properties
	s.oAuthToken = newToken
	return nil
}

func (s *Server) preFlightCheck() bool {
	return !s.oAuthToken.Valid()
}
