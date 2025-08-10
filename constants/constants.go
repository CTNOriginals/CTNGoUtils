package constants

import (
	"path/filepath"
	"runtime"
	"strings"
)

// String Characters
const (
	Alphabet            = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers             = "1234567890"
	AlphaNumeric        = Alphabet + Numbers
	WordCharacters      = AlphaNumeric + "_"
	FileNameCharacters  = WordCharacters + " -."
	DirectoryCharacters = FileNameCharacters + "/\\"
)

// String Flags
const (
	StringError = "<ERROR>"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	RootPath    = filepath.Join(filepath.Dir(b), "../")
	ProjectName = strings.Split(RootPath, "\\")[len(strings.Split(RootPath, "\\"))-1]
)
