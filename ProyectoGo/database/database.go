package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conexion() (conexion *sql.DB) {

	err := godotenv.Load(".envexample")
	if err != nil {
		panic(err)
	}

	Driver := os.Getenv("Driver")
	Server := os.Getenv("Server")
	Port := os.Getenv("Port")
	Usuario := os.Getenv("Usuario")
	Password := os.Getenv("Password")
	DatabaseBD := os.Getenv("DatabaseBD")

	conexion, err2 := sql.Open(Driver,
		Usuario+":"+Password+"@tcp("+Server+":"+Port+")/"+DatabaseBD)

	if err2 != nil {
		panic(err2)
	}

	return conexion

}
