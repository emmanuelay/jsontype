package jsontype

import (
	"testing"
)

func TestNotSetString(t *testing.T) {
	input := `{}`

	p1 := parsePayload(input)
	if (p1.Name.Set != false) || (p1.Name.Valid != false) {
		t.FailNow()
	}
}

func TestSetNullString(t *testing.T) {
	input := `{"name": null}`

	p1 := parsePayload(input)
	if (p1.Name.Set != true) || (p1.Name.Valid != false) {
		t.FailNow()
	}
}

func TestSetValidString(t *testing.T) {
	input := `{"name": "hello, world domination"}`

	p1 := parsePayload(input)
	if (p1.Name.Set != true) || (p1.Name.Valid != true) || (p1.Name.Value != "hello, world domination") {
		t.FailNow()
	}
}
