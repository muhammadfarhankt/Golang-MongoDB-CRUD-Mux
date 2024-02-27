package controller

import (
	"fmt"
	"net/http"
)

func GetAllInternDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all interns")
	fmt.Fprint(w, "Get all interns")
}

func GetInternDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get intern details")
	fmt.Fprint(w, "Get intern details")
}

func UpdateIntern(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update intern")
	fmt.Fprint(w, "Update intern")
}

func DeleteIntern(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete intern")
	fmt.Fprint(w, "Delete intern")
}
