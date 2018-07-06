package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/microblog/backend/data"
)

//PublicationsHandler обрабатывает все запросы к /api/microblog/profiles/1/publications/1/comments
func PublicationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetPublications(w, r)
		return
	}

	if r.Method == "POST" {
		AddPublication(w, r)
		return
	}

	massege := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, massege, http.StatusMethodNotAllowed)
	log.Printf(massege)
}

//GetPublications обрабатывает запросы на получения списка контактов
func GetPublications(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request method: %s", r.Method)
	binaryData, err := json.Marshal(data.PublicationList)
	if err != nil {
		massege := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, massege, http.StatusInternalServerError)
		log.Printf(massege)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		massege := fmt.Sprintf("Hendler write error: %v", err)
		http.Error(w, massege, http.StatusInternalServerError)
		log.Printf(massege)
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
	log.Printf("%+v", publication)
	w.WriteHeader(http.StatusCreated)
}
