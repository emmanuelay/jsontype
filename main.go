package jsontype

import (
	"reflect"
	"strings"
)

const jsonStructTagName = "json"

// StructToMap aggregates a Struct containing jsontype structs into a map
func StructToMap(input interface{}) map[string]interface{} {
	v := reflect.ValueOf(input)
	typeOfS := v.Type()

	out := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		field := typeOfS.Field(i)
		jsonTag := field.Tag.Get(jsonStructTagName)

		if len(jsonTag) > 0 {
			if strings.Contains(jsonTag, ",") {
				jsonTags := strings.Split(jsonTag, ",")
				jsonTag = jsonTags[0]
			}
		}

		fieldValue := v.Field(i).Interface()
		fieldType := v.Field(i).Type()

		switch fieldType.Name() {
		case "String":
			strVal := fieldValue.(String)
			if (strVal.Set == true) && (strVal.Valid == true) {
				out[jsonTag] = strVal.Value
			}
			break
		case "Int":
			intVal := fieldValue.(Int)
			if (intVal.Set == true) && (intVal.Valid == true) {
				out[jsonTag] = intVal.Value
			}
			break
		case "Int64":
			intVal := fieldValue.(Int64)
			if (intVal.Set == true) && (intVal.Valid == true) {
				out[jsonTag] = intVal.Value
			}
			break
		case "Bool":
			boolVal := fieldValue.(Bool)
			if (boolVal.Set == true) && (boolVal.Valid == true) {
				out[jsonTag] = boolVal.Value
			}
			break

		default:
			break
		}
	}

	return out
}
