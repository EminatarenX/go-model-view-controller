package controllers

import (
	"net/http"
	"goapi/db"
	"encoding/json"
	"goapi/models"
)

func ObtenerProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := db.Connection()
	defer db.Close()

	resultado, err := db.Query("SELECT * FROM productos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer resultado.Close()

	var productos []models.Producto

	for resultado.Next() {
		var producto models.Producto
		if err := resultado.Scan(&producto.ID, &producto.Nombre, &producto.Precio, &producto.CantidadEnStock, &producto.MarcaId); err != nil {
			
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		productos = append(productos, producto)
	}

	if err := resultado.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productosJSON, err := json.Marshal(productos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(productosJSON)

	
}