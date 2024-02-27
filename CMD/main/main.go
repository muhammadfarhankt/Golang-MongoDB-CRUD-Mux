package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.InternRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
