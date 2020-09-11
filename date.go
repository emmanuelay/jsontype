package jsontype

import (
	"strings"
	"time"
)

// Date represents an ISO 8601-format date (2006-01-02)
type Date struct {
	Value time.Time
	Valid bool
	Set   bool
}

const dtISOFormat = "2006-01-02"

// UnmarshalJSON ...
func (i *Date) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	dateString := strings.Trim(string(data), "\"")

	if dateString == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp time.Time

	temp, err := time.Parse(dtISOFormat, dateString)
	if err != nil {
		i.Valid = false
		return err
	}

	i.Value = temp
	i.Valid = true
	return nil
}
