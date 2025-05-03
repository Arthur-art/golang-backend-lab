package model

type GetProduct struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name_product"`
	Price float64 `json:"price_product"`
}

type PostProduct struct {
	Name  string  `json:"name_product"`
	Price float64 `json:"price_product"`
}