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
	fmt.Printf(`Usage: %[1]s [N] [int] [ASCII]

N      This number >= 1 is how many bytes are displayed per line. Default = 8.
int    Prints bytes as base 10 numbers.
ASCII  Prints ASCII characters in range [32..126] as characters.

%[1]s reads from standard input and outputs it as hexadecimal
characters (0-9,A-F).
Arguments can be given in any order.
Options int and ASCII can be combined.
Unknown arguments lead to this help message.`,
		appName,
	)
}

func main() {
	bytesPerLine := 8
	showInts := false
	showASCII := false

	for _, arg := range os.Args[1:] {
		if isWord(arg, "int") {
			showInts = true
		} else if isWord(arg, "ascii") {
			showASCII = true
		} else if n, ok := isPositiveInt(arg); ok {
			bytesPerLine = n
		} else {
			printHelp()
			return
		}
	}

	printByte := printHex
	if showInts && showASCII {
		printByte = printASCIIOrInt
	} else if showInts {
		printByte = printInt
	} else if showASCII {
		printByte = printASCIIOrHex
	}

	buf := make([]byte, bytesPerLine)

	for {
		n, err := os.Stdin.Read(buf[:])
		for i := 0; i < n; i++ {
			fmt.Print(printByte(buf[i]))
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

func isPrintableASCII(b byte) bool {
	return 32 <= b && b <= 126
}

func printASCII(b byte) string {
	return "'" + string(rune(b)) + "' "
}

func printHex(b byte) string {
	return fmt.Sprintf("%02X ", b)
}

func printASCIIOrHex(b byte) string {
	if isPrintableASCII(b) {
		return printASCII(b)
	} else {
		return fmt.Sprintf("#%02X ", b)
	}
}

func printInt(b byte) string {
	return fmt.Sprintf("%4d ", b)
}

func printASCIIOrInt(b byte) string {
	if isPrintableASCII(b) {
		return printASCII(b)
	} else {
		return fmt.Sprintf("%3d ", b)
	}
}
