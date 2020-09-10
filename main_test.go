package jsontype

import (
	"encoding/json"
	"testing"
)

type Payload struct {
	Name   String `json:"name,omitempty"`
	Age    Int    `json:"age"`
	UserID Int64  `json:"user_id"`
}

func parsePayload(str string) Payload {
	var payload Payload

	err := json.Unmarshal([]byte(str), &payload)

	if err != nil {
		panic(err)
	}

	return payload
}

func TestStructToMap(t *testing.T) {
	input := `{"name": "hello, world domination","age":43,"user_id":1342992172097472455}`
	p1 := parsePayload(input)

	output := StructToMap(p1)
	if output["age"] != 43 {
		t.FailNow()
	}

	if output["name"] != "hello, world domination" {
		t.FailNow()
	}

}
