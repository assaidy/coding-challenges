package main

import (
	"os"
	"testing"
)

func TestGetFileMeatadata(t *testing.T) {
	f, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	expected := FileMetadata{ByteCnt: 342190, CharCnt: 339292, WordCnt: 58164, LineCnt: 7145}
	actual := *getFileMetadata(f)

	if actual != expected {
		t.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}
