package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/microblog/backend/data"
)

var testPublications = "{\"namepub\":\"Backend\",\"time\":\"00:00\",\"publication\":\"Run backend test...\"}"

func TestCrudHandlers(t *testing.T) {
	cpub := data.NewPublicationList()
	cpro := data.NewProfileList()
	ccom := data.NewCommentList()
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

	testData := strings.NewReader(testPublications)
	req, err = http.NewRequest("POST", "/api/microblog/profiles/0/publications", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0/publications/0" {
		t.Error("Expected another Location")
	}

	if len(cpub.GetPublications()) != 1 {
		t.Error("Expected new value")
	}

	testData = strings.NewReader(testPublications)
	req, err = http.NewRequest("PUT", "/api/microblog/profiles/0/publications/0", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten:%d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/microblog/profiles/0/publications/0" {
		t.Error("Expected another Location")
	}

	if len(cpub.GetPublications()) != 1 {
		t.Error("Expected old value")
	}

	testData = strings.NewReader(testPublications)
	req, err = http.NewRequest("DELETE", "/api/microblog/profiles/0/publications/0", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten:%d)", resp.StatusCode)
	}

	if len(cpub.GetPublications()) != 0 {
		t.Error("Expected old value")
	}
}
