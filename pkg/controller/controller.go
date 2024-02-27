package controller

import (
	"fmt"
	"net/http"

	"github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/models"
)

func GetAllInternDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all interns")
	fmt.Fprint(w, "Get all interns")
}

func GetInternDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get intern details")
	fmt.Fprint(w, "Get intern details")
}

func CreateIntern(w http.ResponseWriter, r *http.Request) {
	Intern := &models.Intern{}
	fmt.Println("Create intern")
	fmt.Fprint(w, "Create intern")
}

func UpdateIntern(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update intern")
	fmt.Fprint(w, "Update intern")
}

func DeleteIntern(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete intern")
	fmt.Fprint(w, "Delete intern")
}
