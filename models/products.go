package models

type ResponseProduct struct {
	Products []Products `json:"products"`
}
type Products struct {
	Proveedor string  `json:"proveedor"`
	Lista     []Lista `json:"lista"`
}
type Lista struct {
	ProdId      string `json:"prod_id"`
	ProdDescrip string `json:"prod_descrip"`
}
