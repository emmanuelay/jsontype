package jsontype

import (
	"encoding/json"
	"testing"
	"time"
)

type testpayload struct {
	Name      String `json:"name,omitempty"`
	Age       Int    `json:"age"`
	UserID    Int64  `json:"user_id"`
	IsActive  Bool   `json:"active"`
	Birthdate Date   `json:"birthdate"`
}

func parsePayload(str string) (testpayload, error) {
	payload := testpayload{}

	err := json.Unmarshal([]byte(str), &payload)

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func TestStructToMap(t *testing.T) {
	input := `{
				"name": "hello, world domination",
				"age": 43,
				"user_id": 1342992172097472455,
				"untracked": true,
				"birthdate": "2019-12-11"
			}`

	parsedPayload, err := parsePayload(input)
	if err != nil {
		t.FailNow()
	}

	output := StructToMap(parsedPayload)

	if len(output) != 4 {
		t.Errorf("Expected output length to be 4, got %v", len(output))
	}

	if output["age"].(int) != 43 {
		t.Errorf("Expected Age = 43, got %v", output["age"])
	}

	if output["name"].(string) != "hello, world domination" {
		t.Errorf("Expected 'name' = 'hello, world domination' got %v", output["name"])
	}

	if output["user_id"].(int64) != 1342992172097472455 {
		t.Errorf("Expected 'user_id' = 1342992172097472455, got %v", output["user_id"])
	}

	dateVal, _ := time.Parse(dtISOFormat, "2019-12-11")

	if output["birthdate"].(time.Time) != dateVal {
		t.Errorf("Expected 'birthdate' = %v, got %v", dateVal, output["birthdate"])
	}
}
