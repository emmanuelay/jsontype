package jsontype

import "encoding/json"

// Int type
type Int struct {
	Value int
	Valid bool
	Set   bool
}

// UnmarshalJSON ...
func (i *Int) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// Int64 type
type Int64 struct {
	Value int64
	Valid bool
	Set   bool
}

// UnmarshalJSON ...
func (i *Int64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}
