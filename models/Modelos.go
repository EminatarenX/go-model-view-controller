package models

import (
	"time"
)


type Cliente struct {
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	CreatedAt time.Time `json:"created_At"`
	Email string `json:"email"`
}

type Producto struct {
	ID int `json:"id"`
	Nombre string `json:"nombre"`
	Precio float64 `json:"precio"`
	CantidadEnStock int `json:"cantidad_en_stock"`
	MarcaId int `json:"marca_id"`
}

type Mensaje struct {
	Mensaje string `json:"mensaje"`
	Error error 
}