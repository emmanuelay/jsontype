package jsontype

import "encoding/json"

type Payload struct {
	Name   String `json:"name"`
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
