package controllers

import (
	"net/http"
	"goapi/db"
	"encoding/json"
	"goapi/models"
	"github.com/gorilla/mux"
	"fmt"
	// "log"
)

func ObtenerClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := db.Connection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM clientes")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	var clientes []models.Cliente

	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.Apellidos, &cliente.CreatedAt, &cliente.Email);err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientesJSON, err := json.Marshal(clientes)

	if err != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(clientesJSON)

}

func ObtenerCliente( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	db := db.Connection()
	defer db.Close()

	resultado, err := db.Query("SELECT * FROM clientes where id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resultado.Close()

	var clientes []models.Cliente

	if resultado.Next() {
		var cliente models.Cliente
		if err := resultado.Scan(&cliente.ID, &cliente.Nombre, &cliente.Apellidos, &cliente.CreatedAt, &cliente.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	if err := resultado.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientesJSON, err := json.Marshal(clientes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(clientesJSON)


}

func BorrarCliente( w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	// var mensaje models.Mensaje

	fmt.Println(id)

	db := db.Connection()
	defer db.Close()

	// resultado, err := db.Query("DELETE FROM clientes where id = ?", id)

	// if err != nil {
	// 	log.Printf("Error al ejecutar la consulta: %v\n", err)
	// 	return
	// }

	

}