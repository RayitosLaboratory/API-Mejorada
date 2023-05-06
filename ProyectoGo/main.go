package main

import (
	api "ProyectoGo/api"
	"context"
	"errors"
	"log" //salida en terminal
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//nodemon --exec go run main.go

//go get init
//go get -u github.com/go-sql-driver/mysql
//go get -u github.com/gorilla/mux
//go get -u github.com/joho/godotenv

func main() {

	contex, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverChan := make(chan os.Signal, 1)
	//Signals a las que se pondra a la escucha...
	signal.Notify(serverChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := api.New(":3000")

	//Ejecucion: localhost:3000
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		} else {
			log.Println("El servidor se detuvo sin inconvenientes")
		}
	}()
	log.Println("Servidor escuchando en el Puerto: 3000")

	<-serverChan

	server.Shutdown(contex)
	log.Println("El servidor se encuentra cerrado")
}
