package cstring

import (
	"fmt"
	"strings"
)

func Repeat(str string, amount int) (result string) {
	for range amount {
		result += str
	}

	return result
}

// TODO finish the below 0 amount
//
// Returns the string with the amount of char's specified before each line.
//
// If amount is less then 0, the function will remove the amount of char that appear at the start of each line.
func Indent(str string, amount int, char string) string {
	var indent = Repeat(char, amount)

	if amount > 0 {
		return indent + strings.Join(strings.Split(str, "\n"), "\n"+indent)
	} else if amount < 0 {
		//TODO Finish this code
		fmt.Println("ALERT: utils.StringIndent() amount below 0 is not functional yet.")
		// var lines = strings.Split(str, "\n")

		// for i, line := range lines {
		// 	var lineChars = strings.Split(line, "")
		// 	var indentChars = strings.Split(char, "")
		// 	var newLine string

		// 	for j := 0; j < len(lineChars) {
		// 		if
		// 	}
		// }

		return "TODO " + str
	}

	return str
}

// Gets the first possible range of characters that that are contained in valid
//
// The return values make up the index range of the substring haysteck[start:end].
// If start and end both return 0, there was no match
func GetValidRange(haysteck string, valid string, startIndex int) (start int, end int) {
	start = -1

	for i := startIndex; i < len(haysteck); i++ {
		var char = rune(haysteck[i])

		if start == -1 && strings.ContainsRune(valid, char) {
			start = i
		} else if start != -1 && !strings.ContainsRune(valid, char) {
			return start, i
		}
	}

	if start == -1 {
		return 0, 0
	}

	return start, len(haysteck) //? Dont substract 1 to account for slices end range non-inclusivity
}

// Checks if all characters contained in haystack are also listed in valid.
func Validate(haysteck string, valid string) bool {
	start, end := GetValidRange(haysteck, valid, 0)
	return start == 0 && end == len(haysteck)
}
