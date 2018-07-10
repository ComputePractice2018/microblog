package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/microblog/backend/data"
)

var testPublications = "{\"namepub\":\"Backend\",\"time\":\"00:00\",\"publication\":\"Run backend test...\"}"
var testProfiles = "{\"nikname\":\"Alx\",\"name\":\"Alexandr\",\"surname\":\"Nikishin\",\"email\":\"San4a-97@mail.ru\",\"github\":\"Sanalx\"}"
var testComments = "{\"time\":\"00:00\",\"comment\":\"Upgrade\"}"
var cpub = data.NewPublicationList()
var cpro = data.NewProfileList()
var ccom = data.NewCommentList()

func TestProfileServerGET(t *testing.T) {

	router := NewRouter(cpub, cpro, ccom)

	req, err := http.NewRequest("GET", "/api/microblog/profiles", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}

	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}
}

func TestProfileServerPOST(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	testData := strings.NewReader(testProfiles)
	req, _ := http.NewRequest("POST", "/api/microblog/profiles", testData)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0" {
		t.Error("Expected another Location")
	}

	if len(cpro.GetProfiles()) != 1 {
		t.Error("Expected new value")
	}
}

func TestProfileServerPUT(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	testData := strings.NewReader(testProfiles)
	req, _ := http.NewRequest("PUT", "/api/microblog/profiles/0", testData)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0" {
		t.Error("Expected another Location")
	}

	if len(cpro.GetProfiles()) != 1 {
		t.Error("Expected old value")
	}
}

func TestProfileServerDELETE(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	req, _ := http.NewRequest("DELETE", "/api/microblog/profiles/0", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten:%d)", resp.StatusCode)
	}

	if len(cpro.GetProfiles()) != 0 {
		t.Error("Expected old value")
	}
}

func TestPublicationServerGET(t *testing.T) {

	router := NewRouter(cpub, cpro, ccom)

	req, err := http.NewRequest("GET", "/api/microblog/profiles/0/publications", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}

	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}
}

func TestPublicationServerPOST(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	testData := strings.NewReader(testPublications)
	req, _ := http.NewRequest("POST", "/api/microblog/profiles/0/publications", testData)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0/publications/0" {
		t.Error("Expected another Location")
	}

	if len(cpub.GetPublications()) != 1 {
		t.Error("Expected new value")
	}
}

func TestPublicationServerPUT(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	testData := strings.NewReader(testPublications)
	req, _ := http.NewRequest("PUT", "/api/microblog/profiles/0/publications/0", testData)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0/publications/0" {
		t.Error("Expected another Location")
	}

	if len(cpub.GetPublications()) != 1 {
		t.Error("Expected old value")
	}
}

func TestPublicationServerDELETE(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	req, _ := http.NewRequest("DELETE", "/api/microblog/profiles/0/publications/0", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten:%d)", resp.StatusCode)
	}

	if len(cpub.GetPublications()) != 0 {
		t.Error("Expected old value")
	}
}

func TestCommentServerGET(t *testing.T) {

	router := NewRouter(cpub, cpro, ccom)

	req, err := http.NewRequest("GET", "/api/microblog/profiles/0/publications/0/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}

	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}
}

func TestCommentServerPOST(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	testData := strings.NewReader(testComments)
	req, _ := http.NewRequest("POST", "/api/microblog/profiles/0/publications/0/comments", testData)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0/publications/0/comments/0" {
		t.Error("Expected another Location")
	}

	if len(ccom.GetComments()) != 1 {
		t.Error("Expected new value")
	}
}

func TestCommentServerPUT(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	testData := strings.NewReader(testComments)
	req, _ := http.NewRequest("PUT", "/api/microblog/profiles/0/publications/0/comments/0", testData)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0/publications/0/comments/0" {
		t.Error("Expected another Location")
	}

	if len(ccom.GetComments()) != 1 {
		t.Error("Expected old value")
	}
}

func TestCommentServerDELETE(t *testing.T) {
	router := NewRouter(cpub, cpro, ccom)
	req, _ := http.NewRequest("DELETE", "/api/microblog/profiles/0/publications/0/comments/0", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten:%d)", resp.StatusCode)
	}

	if len(ccom.GetComments()) != 0 {
		t.Error("Expected old value")
	}
}
