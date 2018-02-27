// tests for LightWeightCTSServer
// TODO: cover endpoints other than these APIs
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// TODO: decouple main and make Router public so test can use it instead of duplicating it here
func testServer() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/cts/text/{urn}", APICTSTextRetrieve).Methods("GET")
	return r
}

// TODO: test real data
func TestGetTextByURN(t *testing.T) {

	// prepare request
	uri := "/api/cts/text/"
	urnInRequest := "tlg0012.tlg001.perseus-grc2:1.1"
	req, err := http.NewRequest(http.MethodGet, uri+urnInRequest, nil)
	log.Println(req)
	if err != nil {
		t.Fatal(err)
	}

	// prepare response
	resp := httptest.NewRecorder()
	testServer().ServeHTTP(resp, req)

	// check expectations
	expectedResponse, _ := json.Marshal("tlg0012.tlg001.perseus-grc2:1.1")
	log.Println(resp.Body)
	acturalResponse := resp.Body.String()
	if string(expectedResponse) != acturalResponse {
		t.Fatalf("Expected %s got %s", expectedResponse, acturalResponse)
	}

}
