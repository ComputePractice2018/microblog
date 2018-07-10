package server

import (
	"github.com/ComputePractice2018/microblog/backend/data"
	"github.com/gorilla/mux"
)

//NewRouter
func NewRouter(publicationList data.EditablePub, profileList data.EditableProfile, commentList data.EditableComment) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/microblog/profiles", AddProfile(profileList)).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/{id}", EditProfile(profileList)).Methods("PUT")

	router.HandleFunc("/api/microblog/profiles/{id}/publications", GetPublications(publicationList)).Methods("GET")
	router.HandleFunc("/api/microblog/profiles/{id}/publications", AddPublication(publicationList)).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}", EditPublication(publicationList)).Methods("PUT")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}", DeletePublication(publicationList)).Methods("DELETE")

	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments", GetComments(commentList)).Methods("GET")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments", AddComment(commentList)).Methods("POST")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments/{idcom}", EditComment(commentList)).Methods("PUT")
	router.HandleFunc("/api/microblog/profiles/{id}/publications/{idpub}/comments/{idcom}", DeleteComment(commentList)).Methods("DELETE")
	return router
}
