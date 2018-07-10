package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/microblog/backend/data"

	"github.com/ComputePractice2018/microblog/backend/server"
)

func main() {
	//name := flag.String("name", "Александр", "имя для преветствия")
	//flag.Parse()

	publicationList := data.NewPublicationList()
	profileList := data.NewProfileList()
	commentList := data.NewCommentList()

	router := server.NewRouter(publicationList, profileList, commentList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
