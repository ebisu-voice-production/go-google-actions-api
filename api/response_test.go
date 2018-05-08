package api

import (
	"encoding/json"
	"testing"
)

func TestResponse(t *testing.T) {
	var response AppResponse
	b, err := json.Marshal(response)
	if err != nil {
		t.FailNow()
	}
	s := string(b)
	if s != `{"expectUserResponse":false}` {
		t.FailNow()
	}
}
