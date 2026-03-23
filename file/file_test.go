package ctnfile

import (
	"os"
	"testing"
)

func TestGetFileRunes(t *testing.T) {
	var tmpdir = os.TempDir()
	var tmpfile, err = os.CreateTemp(tmpdir, "TestGetFileRunes")
	defer func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}()

	if err != nil {
		t.Errorf("Unable to create tmp file: %v", err)
	}

	var str = "éëïäöüËÏÖÜÄ你好Go"
	_, err = tmpfile.WriteString(str)

	if err != nil {
		t.Errorf("Unable to write to file (%s): %v", tmpfile.Name(), err)
	}

	var result = GetFileRunes(tmpfile.Name())
	var expect = []rune(str)

	if len(result) != len(expect) {
		t.Errorf("\nReceived: %s %v\nExpected: %s %v", string(result), result, string(expect), expect)
	}

	for i, char := range result {
		var exp = expect[i]

		if char != exp {
			t.Errorf("\nReceived: %q %d #%x\nExpected: %q %d #%x", char, char, char, exp, exp, exp)
		}
	}
}
