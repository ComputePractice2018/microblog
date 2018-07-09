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
	router.HandleFunc("/api/microblog/profiles", server.AddProfile).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/{id}", server.EditProfile).Methods("PUT")

	router.HandleFunc("/api/microblog/profiles/{id}/publications", server.GetPublications).Methods("GET")
	router.HandleFunc("/api/microblog/profiles/{id}/publications", server.AddPublication).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}", server.EditPublication).Methods("PUT")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}", server.DeletePublication).Methods("DELETE")

	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments", server.GetComments).Methods("GET")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments", server.AddComment).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments/{idcom}", server.EditComment).Methods("PUT")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments/{idcom}", server.DeleteComment).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
