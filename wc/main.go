package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type FileMetadata struct {
	ByteCnt uint
	CharCnt uint
	WordCnt uint
	LineCnt uint
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fmt.Fprintf(os.Stderr, "invalid argument.\n")
		printUsage()
		os.Exit(1)
	}
	if len(args) == 1 {
		filePath := args[0]

		file := mustGetFile(filePath)
		defer file.Close()

		metadata := getFileMetadata(file)
		fmt.Println(metadata.LineCnt, metadata.WordCnt, metadata.ByteCnt, filePath)

	} else if len(args) == 2 {
		option, filePath := args[0], args[1]

		file := mustGetFile(filePath)
		defer file.Close()

		switch option {
		case "-c": // bytes
			byteCnt := getFileMetadata(file).ByteCnt
			fmt.Println(byteCnt, filePath)

		case "-m": // chars
			charCnt := getFileMetadata(file).CharCnt
			fmt.Println(charCnt, filePath)

		case "-w": // words
			wordCnt := getFileMetadata(file).WordCnt
			fmt.Println(wordCnt, filePath)

		case "-l": // lines
			lineCnt := getFileMetadata(file).LineCnt
			fmt.Println(lineCnt, filePath)

		default:
			fmt.Fprintf(os.Stderr, "invalid command '%s'\n", option)
			printUsage()
			os.Exit(1)
		}
	}
}

func printUsage() {
	fmt.Println("\nUsage:")
	fmt.Println("./ccwc <OPTION> <FILE>")
	fmt.Println("   print file data for the specified option.")
	fmt.Println("./ccwc <FILE>")
	fmt.Println("   print file data equivalent to the -c, -l and -w options.")
	fmt.Println("\nOPTIONs:")
	fmt.Println("    -c    print byte count")
	fmt.Println("    -m    print char count")
	fmt.Println("    -w    print word count")
	fmt.Println("    -l    print line count")
}

func mustGetFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file `%s`. check if the filepath is correct.\n", filePath)
		os.Exit(1)
	}
	return file
}

func getFileMetadata(f *os.File) *FileMetadata {
	var (
		reader  = bufio.NewReader(f)
		byteCnt = 0
		charCnt = 0
		wordCnt = 0
		lineCnt = 0
		inWord  = false
	)

	for {
		r, sz, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		byteCnt += sz
		charCnt += 1

		if unicode.IsSpace(r) {
			if inWord {
				wordCnt++
				inWord = false
			}
			if r == '\n' {
				lineCnt++
			}
		} else {
			inWord = true
		}
	}

	return &FileMetadata{
		ByteCnt: uint(byteCnt),
		CharCnt: uint(charCnt),
		WordCnt: uint(wordCnt),
		LineCnt: uint(lineCnt),
	}
}
