package models

type ResponseProduct struct {
	Products []*Products `json:"products"`
}

func NewResponseProduct(products []*Products) *ResponseProduct {
	return &ResponseProduct{Products: products}
}

type Products struct {
	Proveedor string  `json:"proveedor"`
	Lista     []Lista `json:"lista"`
}

func NewProducts(proveedor string, lista []Lista) *Products {
	return &Products{Proveedor: proveedor, Lista: lista}
}

type Lista struct {
	ProdId      string `json:"prod_id"`
	ProdDescrip string `json:"prod_descrip"`
}
