package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4/server"
	"io"
	"io/ioutil"
	"net/http"
)

func GetAllProductForCompany(w http.ResponseWriter, r *http.Request, s *server.Server) (err error) {
	fmt.Println("Start consult product for company")

	_, err = s.ValidationAuthorizeRequest(r)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Authorize fails"))
		return
	}

	products, err := readDataProducts()
	if err != nil {
		_, _ = w.Write([]byte("Error consult products"))
		w.WriteHeader(400)
		return err
	}
	crearJson, _ := json.Marshal(products)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(crearJson)
	if err != nil {
		return err
	}
	return
}

func readDataProducts() (response []Data, err error) {
	url := "https://run.mocky.io/v3/7917c86f-8ef2-449b-a78e-5c55225ff326"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response = make([]Data, 3)

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}
	return
}

type Data struct {
	ProdId      string `json:"prod_id"`
	ProdDescrip string `json:"prod_descrip"`
}
