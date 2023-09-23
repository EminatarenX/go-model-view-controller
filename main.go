package main

import (
    "net/http"
	"goapi/routes"
	"github.com/gorilla/mux"
	
)

func main() {

	r := mux.NewRouter()

	r.PathPrefix("/clientes").Handler(routes.ClientesRouter())
	r.PathPrefix("/productos").Handler(routes.ProductosRouter())

	http.Handle("/", r)

	http.ListenAndServe(":4000", nil)
}