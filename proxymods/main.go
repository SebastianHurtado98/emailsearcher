package main

import (
	"fmt"
    "net/http"
    "github.com/go-chi/chi"
)

func main() {
	fmt.Println("Iniciando servidor")
    r := chi.NewRouter()

    r.Get("/emails/", proxyToZincsearch)

    http.ListenAndServe(":8000", r)
}