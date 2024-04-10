package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var lflag = flag.Bool("l", false, "count lines")
	var wflag = flag.Bool("w", false, "count words")
	var cflag = flag.Bool("c", false, "count bytes")
	var mflag = flag.Bool("m", false, "count characters")
	flag.Parse()

	var fp string = ""

	if flag.NArg() < 1 {
		fmt.Println("no file specified")
	} else if flag.NArg() > 1 {
		fmt.Println("too many files specified, only one file is allowed")
	} else {
		fp = flag.Arg(0)
	}

	if flag.NFlag() == 0 {
		fmt.Println(countBytes(fp), countWords(fp), countLines(fp), fp)
	} else {

		if *lflag {
			fmt.Println(countLines(fp), fp)
		}
		if *wflag {
			fmt.Println(countWords(fp), fp)
		}
		if *cflag {
			fmt.Println(countBytes(fp), fp)
		}
		if *mflag {
			fmt.Println(countChars(fp), fp)
		}
	}
}

func countLines(file string) (out string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(data[:]), "\n")
	lines := len(arr) - 1
	out = fmt.Sprintf("%d", lines)
	return
}

func countBytes(file string) (out string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	bytes := len(data)
	out = fmt.Sprintf("%d", bytes)
	return
}

func countWords(file string) (out string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	str := string(data[:])
	words := len(strings.Fields(str))
	out = fmt.Sprintf("%d", words)
	return
}

func countChars(file string) (out string) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	str := string(data[:])
	count := utf8.RuneCountInString(str)
	out = fmt.Sprintf("%d", int(count))
	return
}
