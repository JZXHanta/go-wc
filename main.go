package main

import (
	"flag"
	"fmt"
	"io"
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

	var fp string
	var pipe string
	var stdin []byte

	if flag.NArg() > 1 {
		fmt.Println("too many files specified, only one file is allowed")
	} else {
		fp = flag.Arg(0)
	}
	if fp == "" {
		pipe = "stdin"
		stdin = readPipe()
	} else {
		pipe = ""
	}

	if flag.NFlag() == 0 {
		fmt.Println(countLines(fp, stdin), countWords(fp, stdin), countBytes(fp, stdin), pipe)
	} else {

		if *lflag {
			fmt.Println(countLines(fp, stdin), fp, pipe)
		}
		if *wflag {
			fmt.Println(countWords(fp, stdin), fp, pipe)
		}
		if *cflag {
			fmt.Println(countBytes(fp, stdin), fp, pipe)
		}
		if *mflag {
			fmt.Println(countChars(fp, stdin), fp, pipe)
		}
	}
}

func countLines(file string, stdin []byte) (out string) {
	var data []byte
	if file == "" {
		if stdin != nil {
			data = stdin
		}
	} else {
		data = getData(file)
	}

	arr := strings.Split(string(data[:]), "\n")
	lines := len(arr) - 1
	out = fmt.Sprintf("%d", lines)
	return
}

func countBytes(file string, stdin []byte) (out string) {
	var data []byte
	if file == "" {
		if stdin != nil {
			data = stdin
		}
	} else {
		data = getData(file)
	}

	bytes := len(data)
	out = fmt.Sprintf("%d", bytes)
	return
}

func countWords(file string, stdin []byte) (out string) {
	var data []byte
	if file == "" {
		if stdin != nil {
			data = stdin
		}
	} else {
		data = getData(file)
	}

	str := string(data[:])
	words := len(strings.Fields(str))
	out = fmt.Sprintf("%d", words)
	return
}

func countChars(file string, stdin []byte) (out string) {
	var data []byte
	if file == "" {
		if stdin != nil {
			data = stdin
		}
	} else {
		data = getData(file)
	}

	str := string(data[:])
	count := utf8.RuneCountInString(str)
	out = fmt.Sprintf("%d", int(count))
	return
}

func readPipe() (file []byte) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	file = []byte(stdin)
	return
}

func getData(file string) (f []byte) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	f = data
	return
}
