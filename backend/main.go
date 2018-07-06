package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/microblog/backend/server"
	"github.com/gorilla/mux"
)

func main() {
	//name := flag.String("name", "Александр", "имя для преветствия")
	//flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/microblog/profiles/1/publications", server.GetPublications).Methods("GET")
	router.HandleFunc("/api/microblog/profiles/1/publications", server.AddPublication).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/1/publications/{id}", server.EditPublication).Methods("PUT")
	router.HandleFunc("/api/microblog/profiles/1/publications/{id}", server.DeletePublication).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
