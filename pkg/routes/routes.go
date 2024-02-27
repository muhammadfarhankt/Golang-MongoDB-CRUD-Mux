package routes

import (
	"github.com/gorilla/mux"
	"github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/controller"
)

var InternRoutes = func(r *mux.Router) {
	r.HandleFunc("/interns", controller.GetAllInternDetails).Methods("GET")
	r.HandleFunc("/intern/{id}", controller.GetInternDetails).Methods("GET")
	r.HandleFunc("/intern/{id}", controller.UpdateIntern).Methods("PUT")
	r.HandleFunc("/intern/{id}", controller.DeleteIntern).Methods("DELETE")
}
