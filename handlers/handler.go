package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/beerproto/beerjson.go"
)

func BeerJSON(w http.ResponseWriter, r *http.Request, beer *beerjson.Beerjson) {
	defer r.Body.Close()

	str :=  &struct {
		Beer *beerjson.Beerjson `json:"beerjson"`
	}{
		Beer: beer,
	}

	err := json.NewDecoder(r.Body).Decode(&str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
