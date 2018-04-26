package api

import (
	"encoding/json"
	"os"
	"testing"
)

func TestRequest(t *testing.T) {
	file, _ := os.Open("./examples/request.json")
	var request AppRequest
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&request)
	if err != nil {
		t.FailNow()
	}
}
