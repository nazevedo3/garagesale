package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/nazevedo3/garagesale/internal/product"
)

// Product has handler methods for dealing with Products
type Product struct {
	Db *sqlx.DB
}

// List gives all products as a list
func (p *Product) List(w http.ResponseWriter, r *http.Request) {

	list, err := product.List(p.Db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error quering db", err)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error marshalling", err)
		return
	}

	w.Header().Set("content-type", "application/json, charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		log.Println("error writing", err)
	}
}
