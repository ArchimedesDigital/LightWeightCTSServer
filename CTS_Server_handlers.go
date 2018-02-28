// Handlers used by LightWeightCTSServer
// TODO: move other handlers here to make the entrypoint(CTS_Server) lighter
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// API handlers that follows CTS standard
// Retrieve some text from static XML served locally
func APICTSTextRetrieve(w http.ResponseWriter, r *http.Request) {
	// prepare filters
	params := mux.Vars(r)
	urn := NewURN(params["urn"])
	ctsParames := CTSParams{
		Sourcetext: urn.WorkFileName,
		StartID:    urn.PassageStart,
		EndID:      urn.PassageEnd,
	}

	// get the text
	textContent := ParseCTS(ctsParames)
	resp := Text{
		RawURN:             urn.rawURN,
		MapCitationPassage: textContent.MapCitationPassage,
	}

	// prepare json resp TODO: encapsulation
	respJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJson)
}
