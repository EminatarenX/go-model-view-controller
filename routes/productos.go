package routes

import (
	"github.com/gorilla/mux"
	"goapi/controllers"
)

func ProductosRouter() *mux.Router {
	r := mux.NewRouter()
	
	r.HandleFunc("/productos", controllers.ObtenerProductos).Methods("GET")

	return r
}