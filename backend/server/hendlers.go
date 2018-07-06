package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ComputePractice2018/microblog/backend/data"
	"github.com/gorilla/mux"
)

//GetPublications обрабатывает запросы на получения списка контактов
func GetPublications(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data.GetPublications())
	if err != nil {
		massege := fmt.Sprintf("Unable to encode data: %v", err)
		http.Error(w, massege, http.StatusInternalServerError)
		log.Printf(massege)
		return
	}
}

//AddPublication обрабатывает POST запрос
func AddPublication(w http.ResponseWriter, r *http.Request) {
	var publication data.Publication
	err := json.NewDecoder(r.Body).Decode(&publication)
	if err != nil {
		massege := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, massege, http.StatusUnsupportedMediaType)
		log.Printf(massege)
		return
	}
	id := data.AddPublication(publication)
	w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

//EditPublication обрабатыает PUT запрос
func EditPublication(w http.ResponseWriter, r *http.Request) {
	var publication data.Publication
	err := json.NewDecoder(r.Body).Decode(&publication)
	if err != nil {
		massege := fmt.Sprintf("Unable to decode PUT data: %v", err)
		http.Error(w, massege, http.StatusUnsupportedMediaType)
		log.Printf(massege)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		massege := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, massege, http.StatusBadRequest)
		log.Printf(massege)
		return
	}

	err = data.EditPublication(publication, id)
	if err != nil {
		massege := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, massege, http.StatusBadRequest)
		log.Printf(massege)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusAccepted)
}

//DeletePublication обрабатывает DELETE запрос
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		massege := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, massege, http.StatusBadRequest)
		log.Printf(massege)
		return
	}

	err = data.RemovePublication(id)
	if err != nil {
		massege := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, massege, http.StatusBadRequest)
		log.Printf(massege)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)
}
