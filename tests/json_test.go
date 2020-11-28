package tests

import (
	"encoding/json"
	"testing"
)

type JsonBool struct {
	Active bool   `json:"active"`
	Name   string `json:"name,omitempty"`
}

func TestJson(t *testing.T) {

	var n = new(JsonBool)
	js := `{"active":false,"name":"zhaoyunxing"}`

	if err := json.Unmarshal([]byte(js), n); err != nil {
		t.Fail()
	}

	t.Logf("%v", n)

}
