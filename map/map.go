package cmap

import (
	"fmt"
	"strings"
)

func Keys[K comparable, V any](m map[K]V) (ret []K) {
	ret = make([]K, len(m))

	i := 0
	for key := range m {
		ret[i] = key
		i++
	}

	return ret
}

func Values[K comparable, V any](m map[K]V) (ret []V) {
	ret = make([]V, len(m))

	i := 0
	for key := range m {
		ret[i] = m[key]
		i++
	}

	return ret
}

func ToStringFunc[K comparable, V any](obj map[K]V, stringConvert func(val V) string) string {
	if len(obj) == 0 {
		return fmt.Sprintf("%T { }", obj)
	}

	var lines []string
	for key, item := range obj {
		var itemLines = strings.Split(stringConvert(item), "\n")
		lines = append(lines, fmt.Sprintf("%v: {\n  %s\n }", key, strings.Join(itemLines, "\n  ")))
	}

	return fmt.Sprintf("%T {\n %s\n}", obj, strings.Join(lines, "\n "))
}

func ToString[K comparable, V string](obj map[K]V) string {
	return ToStringFunc(obj, func(v V) string {
		return string(v)
	})
}

func ToStringer[K comparable, V fmt.Stringer](obj map[K]V) string {
	return ToStringFunc(obj, V.String)
}
