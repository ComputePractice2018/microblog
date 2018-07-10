package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/microblog/backend/data"

	"github.com/ComputePractice2018/microblog/backend/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	connection := flag.String("connection", "microblog:SecretPassword@tcp(db:3306)/microblog", "mysql connection string")
	flag.Parse()

	publicationList, err := data.NewDBPublicationList(*connection)
	profileList, err := data.NewDBProfileList(*connection)
	commentList, err := data.NewDBCommentList(*connection)

	defer profileList.CloseProfile()
	defer publicationList.ClosePublication()
	defer commentList.CloseComment()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(publicationList, profileList, commentList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
