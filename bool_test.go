package jsontype

import (
	"testing"
)

func TestNotSetBool(t *testing.T) {
	input := `{}`

	p1, _ := parsePayload(input)

	if (p1.IsActive.Set != false) || (p1.IsActive.Valid != false) {
		t.FailNow()
	}
}

func TestSetNullBool(t *testing.T) {
	input := `{"active": null}`

	p1, _ := parsePayload(input)

	if (p1.IsActive.Set != true) || (p1.IsActive.Valid != false) {
		t.FailNow()
	}
}

func TestSetValidBool(t *testing.T) {
	input := `{"active": true}`

	p1, _ := parsePayload(input)

	if (p1.IsActive.Set != true) || (p1.IsActive.Valid != true) || (p1.IsActive.Value != true) {
		t.FailNow()
	}
}
