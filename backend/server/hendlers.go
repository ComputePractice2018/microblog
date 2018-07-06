package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/microblog/backend/data"
)

//GetPublications обрабатывает запросы на получения списка контактов
func GetPublications(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.PublicationList)
	if err != nil {
		w.Header().Add("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON marshaling error: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		log.Printf("Hendler write error: %v", err)
	}
}
