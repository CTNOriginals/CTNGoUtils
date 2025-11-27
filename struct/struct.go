package ctnstruct

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func Keys(obj any) (keys []string) {
	var val = reflect.Indirect(reflect.ValueOf(obj))
	var length = val.Type().NumField()
	keys = make([]string, length)

	for i := range length {
		keys[i] = val.Type().Field(i).Name
	}

	return keys
}
func Values(obj any) (vals []any) {
	var val = reflect.Indirect(reflect.ValueOf(obj))
	var length = val.Type().NumField()
	vals = make([]any, length)

	for i := range length {
		vals[i] = val.Field(i)
	}

	return vals
}

func ToString(obj any, excludeKeys ...string) string {
	var lines []string
	var values = Values(obj)
	for i, key := range Keys(obj) {
		if slices.Contains(excludeKeys, key) {
			continue
		}
		lines = append(lines, fmt.Sprintf("%s: %v", key, values[i]))
	}

	return strings.Join(lines, "\n")
}

func Compare(a any, b any) bool {
	var aValues = Values(a)
	var bValues = Values(b)

	for i, val := range aValues {
		if val != bValues[i] {
			return false
		}
	}

	return true
}
