package jsontype

import (
	"encoding/json"
	"testing"
)

type testpayload struct {
	Name     String `json:"name,omitempty"`
	Age      Int    `json:"age"`
	UserID   Int64  `json:"user_id"`
	IsActive Bool   `json:"active"`
}

func parsePayload(str string) testpayload {
	var payload testpayload

	err := json.Unmarshal([]byte(str), &payload)

	if err != nil {
		panic(err)
	}

	return payload
}

func TestStructToMap(t *testing.T) {
	input := `{"name": "hello, world domination","age":43,"user_id":1342992172097472455,"untracked":true}`
	parsedPayload := parsePayload(input)
	output := StructToMap(parsedPayload)

	if len(output) != 3 {
		t.Errorf("Expected output length to be 3, got %v", len(output))
	}

	if output["age"].(int) != 43 {
		t.Errorf("Expected Age == 43, got %v", output["age"])
	}

	if output["name"].(string) != "hello, world domination" {
		t.Errorf("Expected 'name' = 'hello, world domination' got %v", output["name"])
	}

	if output["user_id"].(int64) != 1342992172097472455 {
		t.Errorf("Expected 'user_id' = 1342992172097472455, got %v", output["user_id"])
	}
}
