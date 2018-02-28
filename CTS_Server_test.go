// tests for LightWeightCTSServer
// TODO: cover endpoints other than these APIs
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func Test_URN_struct_is_constructed_correctly(t *testing.T) {
	// TODO: table tests

	expectedURN := URN{
		rawURN:       "tlg0012.tlg001.perseus-grc2:1.1-1.2",
		WorkFileName: "tlg0012.tlg001.perseus-grc2.xml",
		PassageStart: "1.1",
		PassageEnd:   "1.2",
	}

	rawURN := "tlg0012.tlg001.perseus-grc2:1.1-1.2"
	urn := NewURN(rawURN)

	if !reflect.DeepEqual(urn, &expectedURN) {
		t.Errorf("Expected %s got %s", &expectedURN, urn)
	}

}

func Test_can_get_Text_by_URN(t *testing.T) {

	// prepare request
	uri := "/api/cts/text/"
	urnInRequest := "tlg0012.tlg001.perseus-grc2:1.1-1.2"
	req, err := http.NewRequest(http.MethodGet, uri+urnInRequest, nil)
	if err != nil {
		t.Fatal(err)
	}

	// prepare response
	resp := httptest.NewRecorder()
	testServer().ServeHTTP(resp, req)

	// check expectations
	expectedPassageCitationIDs := []string{"1.1", "1.2"}
	acturalResponse := resp.Body.String()
	acturalText := &Text{}
	_ = json.Unmarshal([]byte(acturalResponse), acturalText)
	for _, citationID := range expectedPassageCitationIDs {
		_, ok := acturalText.MapCitationPassage[citationID]
		if !ok {
			t.Errorf("Expected citationID %s in %s", citationID, acturalText.MapCitationPassage)
		}
	}

}
