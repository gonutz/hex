package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func printHelp() {
	appName := filepath.Base(os.Args[0])
	fmt.Printf(`Usage: %[1]s [N] [int]

N    a number >= 1 is how many bytes are displayed per line.
int  means that base numbers are printed instead of hex characters.

%[1]s reads from standard input and outputs it as
hexadecimal (0..9, A..F) characters.
Arguments can be given in any order.
Unknown arguments lead to this help message.`,
		appName,
	)
}

func main() {
	format := "%02X "
	bytesPerLine := 8

	for _, arg := range os.Args[1:] {
		if isWord(arg, "int") {
			format = "%4d "
		} else if n, ok := isPositiveInt(arg); ok {
			bytesPerLine = n
		} else {
			printHelp()
			return
		}
	}

	buf := make([]byte, bytesPerLine)

	for {
		n, err := os.Stdin.Read(buf[:])
		for i := 0; i < n; i++ {
			fmt.Printf(format, buf[i])
		}
		fmt.Println()
		if err != nil {
			break
		}
	}
}

func isWord(s, word string) bool {
	return strings.ToLower(s) == word
}

func isPositiveInt(s string) (n int, ok bool) {
	x, err := strconv.Atoi(s)
	return x, err == nil && x >= 1
}
