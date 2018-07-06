package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/microblog/backend/server"
)

func main() {
	//name := flag.String("name", "Александр", "имя для преветствия")
	//flag.Parse()

	http.HandleFunc("/api/microblog/profiles/1/publications", server.GetPublications)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
