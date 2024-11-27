package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// TODO: find a better way to parse cli args

type FileMetadata struct {
	ByteCnt uint
	CharCnt uint
	WordCnt uint
	LineCnt uint
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "invalid argument.\n")
		printUsage()
		os.Exit(1)
	}
	option, filePath := args[0], args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file `%s`. check if the filepath is correct.\n", filePath)
		os.Exit(1)
	}
	defer file.Close()

	switch option {
	case "-c": // bytes
		byteCnt := getByteCount(file)
		fmt.Println(byteCnt, filePath)

	case "-l": // lines
		lineCnt := getLineCount(file)
		fmt.Println(lineCnt, filePath)

	case "-m": // chars
		charCnt := getCharCount(file)
		fmt.Println(charCnt, filePath)

	default:
		fmt.Fprintf(os.Stderr, "invalid command '%s'\n", option)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("\nUsage:")
	fmt.Println("./ccwc <OPTION> <FILE>")
	fmt.Println("\nOPTIONs:")
	fmt.Println("    -c    print bytes count")
	fmt.Println("    -l    print line count")
	fmt.Println("    -m    print character count")
}

func getFileMetadata(f *os.File) *FileMetadata {
	var (
		buf     = make([]byte, 20*1024)
		byteCnt = 0
		charCnt = 0
		wordCnt = 0 // TODO:
		lineCnt = 0
	)

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		byteCnt += n
		charCnt += utf8.RuneCount(buf[:n])
		lineCnt += strings.Count(string(buf[:n]), "\n")
	}

	return &FileMetadata{
		ByteCnt: uint(byteCnt),
		CharCnt: uint(charCnt),
		WordCnt: uint(wordCnt),
		LineCnt: uint(lineCnt),
	}
}

func getByteCount(f *os.File) uint {
	info, _ := f.Stat()
	return uint(info.Size())
}

func getLineCount(f *os.File) uint {
	buf := make([]byte, 30*1024)
	cnt := 0

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		cnt += strings.Count(string(buf[:n]), "\n")
	}

	return uint(cnt)
}

func getCharCount(f *os.File) uint {
	buf := make([]byte, 20*1024)
	cnt := 0

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		cnt += utf8.RuneCount(buf[:n])
	}

	return uint(cnt)
}
