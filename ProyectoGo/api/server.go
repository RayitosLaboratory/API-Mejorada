package api

import "net/http"

func New(addr string) *http.Server {

	InitRoutes()

	return &http.Server{
		Addr: addr,
	}
}
