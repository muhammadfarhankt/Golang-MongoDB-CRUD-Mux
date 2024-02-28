package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/models"
	"github.com/muhammadfarhankt/Golang-MongoDB-CRUD-Mux/pkg/utils"
)

func GetAllInternDetails(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get all interns")
	//fmt.Fprint(w, "Get all interns")
	allInterDetails := models.GetAllInternDetails()
	res, _ := json.Marshal(allInterDetails)
	//fmt.Println("res: ", res)
	//fmt.Println("allInterDetails: ", allInterDetails)
	w.Header().Set("Content-Type", "application/json")
	if allInterDetails == nil {
		w.WriteHeader(http.StatusNotFound)
		res = []byte(`{"message": "Interns not found"}`)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}

func GetInternDetails(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get intern details")
	//fmt.Fprint(w, "Get intern details")
	vars := mux.Vars(r)
	id := vars["id"]
	//fmt.Println("id: ", id)
	InternDetails := models.GetInternDetails(id)
	res, _ := json.Marshal(InternDetails)
	w.Header().Set("Content-Type", "application/json")
	if InternDetails.InternId == "" {
		w.WriteHeader(http.StatusNotFound)
		res = []byte(`{"message": "Intern not found"}`)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}

func CreateIntern(w http.ResponseWriter, r *http.Request) {
	CreateIntern := &models.Intern{}
	// Parse the body of the request and store the result in the Intern object
	//fmt.Println("r.Body", r.Body)
	utils.ParseBody(r, CreateIntern)
	intern := CreateIntern.CreateIntern()
	res, _ := json.Marshal(intern)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	//fmt.Fprint(w, "Create intern")
}

func UpdateIntern(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Update intern")
	vars := mux.Vars(r)
	id := vars["id"]
	UpdateIntern := &models.Intern{}
	utils.ParseBody(r, UpdateIntern)
	internDetailsToUpdate := UpdateIntern.UpdateInternDetails(id)
	res, _ := json.Marshal(internDetailsToUpdate)
	w.Header().Set("Content-Type", "application/json")
	if internDetailsToUpdate == nil {
		w.WriteHeader(http.StatusNotFound)
		res = []byte(`{"message": "Intern not found"}`)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
	//fmt.Fprint(w, "Update intern")
}

func DeleteIntern(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Delete intern")
	// fmt.Fprint(w, "Delete intern")
	vars := mux.Vars(r)
	id := vars["id"]
	//fmt.Println("id: ", id)
	delete := models.DeleteIntern(id)
	res, _ := json.Marshal(delete)
	//fmt.Println("delete: ", delete)
	//fmt.Println("res: ", res)
	w.Header().Set("Content-Type", "application/json")
	if delete.InternId == "" {
		w.WriteHeader(http.StatusNotFound)
		res = []byte(`{"message": "Intern not found"}`)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}
