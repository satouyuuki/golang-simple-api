package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

// 初期化
func TestMain(m *testing.M) {
	router = setupRouter()
	os.Exit(m.Run())
}

func TestGetAlbums(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetAlbumByID(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body.String())
}

func TestDeleteAlbumByID(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	fmt.Println(w.Body.String())
}

func TestPostAlbums(t *testing.T) {
	jsonBody := []byte(`{
		"id": "4",
		"title": "The Modern Sound of Betty Carter",
		"artist": "Betty Carter",
		"price": 49.99
	}`)
	bodyReader := bytes.NewReader(jsonBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	fmt.Println(w.Body.String())
}

func TestPutAlbumByID(t *testing.T) {
	jsonBody := []byte(`{
		"id": "3",
		"title": "title 3 example update",
		"artist": "artist 3 update",
		"price": 55.99
	}`)
	bodyReader := bytes.NewReader(jsonBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/albums/3", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	fmt.Println(w.Body.String())
}
