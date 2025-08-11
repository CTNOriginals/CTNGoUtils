package cslice

import (
	"fmt"
	"strings"

	cstring "github.com/CTNOriginals/CTNGoUtils/v2/string"
)

func Splice[T comparable](slice []T, start int, count int) (remaining []T, deleted []T) {
	remaining = make([]T, len(slice)-count)
	deleted = make([]T, count)

	deleteCount := 0
	for i, item := range slice {
		if i >= start && i < (start+count) {
			deleted[deleteCount] = item
			deleteCount++
			continue
		}

		remaining[i-deleteCount] = item
	}

	return
}

func ToString[T interface{ fmt.Stringer }](slice []T, withIndexes bool) string {
	if len(slice) == 0 {
		return "[]"
	}

	var items []string
	for i, item := range slice {
		var str = item.String()

		if withIndexes {
			str = fmt.Sprintf("%d:%s", i, str)
		}

		var lines = strings.Split(str, "\n")

		if len(lines) > 1 {
			str = lines[0] + "\n"
			str += cstring.Indent(strings.Join(lines[1:], "\n"), 1, " ")
		}

		str = cstring.Indent(str, 1, " ")
		items = append(items, str)
	}

	return fmt.Sprintf("[\n%s\n]", strings.Join(items, "\n"))
}
