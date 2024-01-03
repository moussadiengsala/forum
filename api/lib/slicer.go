package lib

import (
	"reflect"
)

func Slicer(data interface{}, isIDFieldNeeded bool) []interface{} {
	val := reflect.ValueOf(data)
	var result []interface{}
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		if !isIDFieldNeeded && fieldName == "ID" {
			continue
		}
		result = append(result, val.Field(i).Interface())
	}

	return result
}

func ExtractFields(obj interface{}) []interface{} {
	val := reflect.ValueOf(obj)

	// Check if obj is a pointer and dereference it if needed
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	numFields := val.NumField()

	fields := make([]interface{}, numFields)
	for i := 0; i < numFields; i++ {
		fields[i] = val.Field(i).Addr().Interface()
	}

	return fields
}
