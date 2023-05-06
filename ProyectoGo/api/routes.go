package api

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func InitRoutes() {
	//http.HandleFunc("/", Index)
	//http.HandleFunc("/clientes", OptFunc1)
	//http.HandleFunc("/clientes/:idcliente", OptFunc2)
	//http.HandleFunc("/clientes/registros/id", getJoin)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/clientes", GetAll).Methods(http.MethodGet)
	router.HandleFunc("/clientes", InsertOne).Methods(http.MethodPut)
	router.HandleFunc("/clientes/{id_cliente}", GetOne).Methods(http.MethodGet)
	router.HandleFunc("/clientes/{id_cliente}", UpdateOne).Methods(http.MethodPatch)
	router.HandleFunc("/clientes/{id_cliente}", DeleteOne).Methods(http.MethodDelete)

	http.Handle("/", router)
}
