package routes

import (
	"github.com/gorilla/mux"
	"goapi/controllers"

)

func ClientesRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/clientes", controllers.ObtenerClientes).Methods("GET")
	r.HandleFunc("/clientes/{id}", controllers.ObtenerCliente).Methods("GET")
	r.HandleFunc("/clientes/{id}", controllers.BorrarCliente).Methods("DELETE")

	return r
}