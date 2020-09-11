package jsontype

import (
	"testing"
	"time"
)

func TestNotSetDate(t *testing.T) {
	input := `{}`

	p1, _ := parsePayload(input)
	if (p1.Birthdate.Set != false) || (p1.Birthdate.Valid != false) {
		t.FailNow()
	}
}

func TestSetNullDate(t *testing.T) {
	input := `{"birthdate": null}`

	p1, _ := parsePayload(input)
	if (p1.Birthdate.Set != true) || (p1.Birthdate.Valid != false) {
		t.FailNow()
	}
}

func TestSetValidDate(t *testing.T) {
	input := `{"birthdate": "1999-02-09"}`

	val, _ := time.Parse(dtISOFormat, "1999-02-09")

	p1, _ := parsePayload(input)

	if (p1.Birthdate.Set != true) || (p1.Birthdate.Valid != true) || (p1.Birthdate.Value != val) {
		t.FailNow()
	}
}

func TestSetInvalidDate(t *testing.T) {
	input := `{"birthdate": "1999-13-09","other_value":99}`

	_, err := parsePayload(input)
	if err == nil {
		t.Error("Expected error")
	}
}
