package rest

import "net/http"

// Rest is a struct
type Rest struct{}

// NewRest is a rest
func NewRest() *Rest {
	return &Rest{}
}

func (ri *Rest) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (ri *Rest) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (ri *Rest) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (ri *Rest) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
