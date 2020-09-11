package jsontype

import (
	"testing"
)

func TestNotSetInt(t *testing.T) {
	input := `{}`

	p1, _ := parsePayload(input)
	if (p1.Age.Set != false) || (p1.Age.Valid != false) {
		t.FailNow()
	}
}

func TestSetNullInt(t *testing.T) {
	input := `{"age": null}`

	p1, _ := parsePayload(input)
	if (p1.Age.Set != true) || (p1.Age.Valid != false) {
		t.FailNow()
	}
}

func TestSetValidInt(t *testing.T) {
	input := `{"age": 42}`

	p1, _ := parsePayload(input)
	if (p1.Age.Set != true) || (p1.Age.Valid != true) || (p1.Age.Value != 42) {
		t.FailNow()
	}
}

func TestNotSetInt64(t *testing.T) {
	input := `{}`

	p1, _ := parsePayload(input)
	if (p1.UserID.Set != false) || (p1.UserID.Valid != false) {
		t.FailNow()
	}
}

func TestSetNullInt64(t *testing.T) {
	input := `{"user_id": null}`

	p1, _ := parsePayload(input)
	if (p1.UserID.Set != true) || (p1.UserID.Valid != false) {
		t.FailNow()
	}
}

func TestSetValidInt64(t *testing.T) {
	input := `{"user_id": 1342992172097472455}`

	p1, _ := parsePayload(input)
	if (p1.UserID.Set != true) || (p1.UserID.Valid != true) || (p1.UserID.Value != 1342992172097472455) {
		t.FailNow()
	}
}
