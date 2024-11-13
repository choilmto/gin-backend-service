package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// integration test
func TestGetAlbum(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)
	var albums []Album
	json.Unmarshal(w.Body.Bytes(), &albums)

	assert.Equal(t, 200, w.Code)
	if len(albums) > 0 {
		assert.Equal(t, "1", albums[0].ID)
	}
}

// integration test
func TestPostAlbum(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	// Create an example album for testing
	exampleAlbum := Album{
		ID:     "test_name",
		Title:  "test_title",
		Artist: "test_artist",
		Price:  5.00,
	}
	albumJson, _ := json.Marshal(exampleAlbum)
	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(string(albumJson)))
	router.ServeHTTP(w, req)

	dst := &bytes.Buffer{}
	json.Compact(dst, w.Body.Bytes())
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, string(albumJson), dst.String())
}
