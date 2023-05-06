package clientes

type Cliente struct {
	Id_cliente int    `json:"id_cliente"` //el nombre del campo debe empezar en mayuscula
	Nombre     string `json:"nombre"`
	Direccion  string `json:"direccion"`
	Telefono   string `json:"telefono"`
}
