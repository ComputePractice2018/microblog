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

//GetProfiles обрабатывает запросы на получения списка профилей
func GetProfiles(cl data.EditableProfile) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetProfiles())
		if err != nil {
			massege := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, massege, http.StatusInternalServerError)
			log.Printf(massege)
			return
		}
	}
}

//AddProfile обрабатывает POST запрос
func AddProfile(cl data.EditableProfile) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var profile data.Profile
		err := json.NewDecoder(r.Body).Decode(&profile)
		if err != nil {
			massege := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, massege, http.StatusUnsupportedMediaType)
			log.Printf(massege)
			return
		}
		id := cl.AddProfile(profile)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditProfile обрабатыает PUT запрос
func EditProfile(cl data.EditableProfile) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var profile data.Profile
		err := json.NewDecoder(r.Body).Decode(&profile)
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

		err = cl.EditProfile(id, profile)
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)
	}
}

//DeleteProfiles обрабатывает DELETE запрос
func DeleteProfile(cl data.EditableProfile) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}

		err = cl.RemoveProfile(id)
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}

//GetPublications обрабатывает запросы на получения списка публикаций
func GetPublications(cl data.EditablePub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetPublications())
		if err != nil {
			massege := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, massege, http.StatusInternalServerError)
			log.Printf(massege)
			return
		}
	}
}

//AddPublication обрабатывает POST запрос
func AddPublication(cl data.EditablePub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var publication data.Publication
		err := json.NewDecoder(r.Body).Decode(&publication)
		if err != nil {
			massege := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, massege, http.StatusUnsupportedMediaType)
			log.Printf(massege)
			return
		}

		idpub := cl.AddPublication(publication)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(idpub))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditPublication обрабатыает PUT запрос
func EditPublication(cl data.EditablePub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var publication data.Publication
		err := json.NewDecoder(r.Body).Decode(&publication)
		if err != nil {
			massege := fmt.Sprintf("Unable to decode PUT data: %v", err)
			http.Error(w, massege, http.StatusUnsupportedMediaType)
			log.Printf(massege)
			return
		}

		varsp := mux.Vars(r)
		idpub, err := strconv.Atoi(varsp["idpub"])
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID publication: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}

		err = cl.EditPublication(idpub, publication)
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID publication: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)
	}
}

//DeletePublication обрабатывает DELETE запрос
func DeletePublication(cl data.EditablePub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		varsp := mux.Vars(r)
		idpub, err := strconv.Atoi(varsp["idpub"])
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID publication: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}

		err = cl.RemovePublication(idpub)
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID publication: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}

//GetComments обрабатывает запросы на получения списка комментариев
func GetComments(cl data.EditableComment) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetComments())
		if err != nil {
			massege := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, massege, http.StatusInternalServerError)
			log.Printf(massege)
			return
		}
	}
}

//AddComment обрабатывает POST запрос
func AddComment(cl data.EditableComment) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment data.Comment
		err := json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			massege := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, massege, http.StatusUnsupportedMediaType)
			log.Printf(massege)
			return
		}
		idcom := cl.AddComment(comment)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(idcom))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditComment обрабатыает PUT запрос
func EditComment(cl data.EditableComment) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment data.Comment
		err := json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			massege := fmt.Sprintf("Unable to decode PUT data: %v", err)
			http.Error(w, massege, http.StatusUnsupportedMediaType)
			log.Printf(massege)
			return
		}

		varspc := mux.Vars(r)
		idcom, err := strconv.Atoi(varspc["idcom"])
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID commets: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}

		err = cl.EditComment(idcom, comment)
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID comments: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)
	}
}

//DeleteComment обрабатывает DELETE запрос
func DeleteComment(cl data.EditableComment) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		varspc := mux.Vars(r)
		idcom, err := strconv.Atoi(varspc["idcom"])
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID comments: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}

		err = cl.RemoveComment(idcom)
		if err != nil {
			massege := fmt.Sprintf("Incorrect ID comments: %v", err)
			http.Error(w, massege, http.StatusBadRequest)
			log.Printf(massege)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}
