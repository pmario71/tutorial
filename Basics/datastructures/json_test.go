package datastructures

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GreetingRequest struct {
	Name string `json:"name"`
}

func Test_decoding_simple_object(t *testing.T) {
	const msg = `{ "name": "Bert" }`

	var req GreetingRequest
	d := json.NewDecoder(strings.NewReader(msg))

	d.Decode(&req)

	assert.Equal(t, "Bert", req.Name, "Name not decoded correctly")
}
