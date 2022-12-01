package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestTrimmer(t *testing.T) {
	filler := strings.Repeat("abc", 2000)
	afterTrimmed := "[{" + filler + "}]"
	testInput := bytes.NewBufferString(filler + afterTrimmed)
	trimmer := NewTrimmer(testInput, '[')
	trimmed, err := io.ReadAll(trimmer)
	if err != nil {
		t.Error(err)
	}
	if string(trimmed) != afterTrimmed {
		t.Errorf("Expected %s, received %s", afterTrimmed, trimmed)
	}

}
