package api

import (
	clientes "ProyectoGo/clientes"
	BD "ProyectoGo/database"
	"encoding/json"
	"fmt"      //formatear salida
	"log"      //salida en terminal
	"net/http" //peticiones http
	"strconv"

	"github.com/gorilla/mux"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusBadRequest) //codido de estado 400
		fmt.Fprint(res, "Solo valido con el METODO GET")
		return
	}

	res.WriteHeader(http.StatusOK) //codigo de estado 200

	cliente := clientes.Cliente{
		Id_cliente: 6,
		Nombre:     "Juan",
		Direccion:  "Calle 123, Ciudad",
		Telefono:   "555-1234",
	}

	//opcion 1
	//fmt.Fprint(res, "Hola Mundo")

	//opcion 2 ***
	json.NewEncoder(res).Encode(cliente)

	//opcion 3
	//json_cliente, _ := json.Marshal(cliente)
	//fmt.Fprint(res, string(json_cliente))

	log.Println("Hola!! en la Terminal del Servidor! :D")

}

func GetAll(res http.ResponseWriter, req *http.Request) {

	cn := BD.Conexion()

	//PARA PROBAR LA CONEXION
	//cmd, err := cn.Prepare("SELECT NOW()")
	//fmt.Fprintf(res, "Conexion establecida")

	datosDevueltos, err := cn.Query("SELECT * FROM CLIENTES")

	if err != nil {
		panic(err)
	}

	ObjCliente := clientes.Cliente{}
	arrayCliente := []clientes.Cliente{}

	for datosDevueltos.Next() {
		var idcliente int
		var nombre, direccion, telefono string
		err = datosDevueltos.Scan(&idcliente, &nombre, &telefono, &direccion)
		ObjCliente.Id_cliente = idcliente
		ObjCliente.Nombre = nombre
		ObjCliente.Direccion = direccion
		ObjCliente.Telefono = telefono

		arrayCliente = append(arrayCliente, ObjCliente)
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(arrayCliente)

}

func GetOne(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id_cliente, err := strconv.Atoi(params["id_cliente"])

	//fmt.Fprintf(res, "%v", params["id_cliente"])

	cn := BD.Conexion()

	//op1
	//datosDevueltos, err := cn.Query("SELECT * FROM CLIENTES WHERE ID=?", id_cliente)

	//op2
	datosDevueltos, err := cn.Query("call sp_Select_OneClient(?)", id_cliente)

	if err != nil {
		panic(err)
	}

	ObjCliente := clientes.Cliente{}

	for datosDevueltos.Next() {
		var idcliente int
		var nombre, direccion, telefono string
		err = datosDevueltos.Scan(&idcliente, &nombre, &telefono, &direccion)
		ObjCliente.Id_cliente = idcliente
		ObjCliente.Nombre = nombre
		ObjCliente.Direccion = direccion
		ObjCliente.Telefono = telefono
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ObjCliente)

}

func InsertOne(res http.ResponseWriter, req *http.Request) {

	cliente := &clientes.Cliente{}

	err := json.NewDecoder(req.Body).Decode(cliente)

	cn := BD.Conexion()
	cmd, err := cn.Prepare("INSERT INTO clientes (id, nombre, direccion, telefono) VALUES (10, Alejandro, Tampico, 8331623453)")

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Datos de entrada en formato incorrecto: %v", err)
		return
	}

	cmd.Exec(cliente.Id_cliente, cliente.Nombre, cliente.Direccion, cliente.Telefono)

	res.WriteHeader(http.StatusOK) //codigo de estado 200
	fmt.Fprint(res, "Se ha almacenado el cliente correctamente")

}

func UpdateOne(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id_cliente, errP := strconv.Atoi(params["id_cliente"])

	if errP != nil {
		panic(errP)
	}

	//fmt.Fprintf(res, "%v", params["id_cliente"])

	cliente := &clientes.Cliente{}

	err := json.NewDecoder(req.Body).Decode(cliente)

	cn := BD.Conexion()
	cmd, err := cn.Prepare("UPDATE clientes SET nombre=?, direccion=?, telefono=? WHERE id=?")

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Datos de entrada en formato incorrecto: %v", err)
		return
	}

	cmd.Exec(cliente.Nombre, cliente.Direccion, cliente.Telefono, id_cliente)

	res.WriteHeader(http.StatusOK) //codigo de estado 200
	fmt.Fprint(res, "Se ha actualizado el cliente correctamente")

}

func DeleteOne(res http.ResponseWriter, req *http.Request) {

}
