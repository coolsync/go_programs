package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// counts of unicode charaters
	counts := make(map[rune]int)
	// counts of len utf-8 charaters
	var utf_len [utf8.UTFMax + 1]int
	// counts of invalid uff-8 charaters
	invalid := 0

	// str := "ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes."
	// reader := bufio.NewReader(strings.NewReader(str))
	reader := bufio.NewReader(os.Stdin)

	// ReadRune reads a single UTF-8 encoded Unicode character and returns the rune and its size in bytes.
	//  If the encoded rune is invalid,
	//  it consumes one byte and returns unicode.ReplacementChar (U+FFFD) with a size of 1.
	for {
		r, size, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && size == 1 {
			invalid++
			continue
		}

		counts[r]++
		utf_len[size]++

	}

	fmt.Print("\nrune\tcounts:\n")

	for c, n := range counts {
		fmt.Printf("\n%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcounts:\n")
	for i, size := range utf_len {
		if i > 0 {
			fmt.Printf("\n%d\t%d\n", i, size)
		}
	}

	fmt.Print("\ninvalid\tcounts:\n")

	if invalid > 0 {
		fmt.Printf("\ninvalid of unicode char counts: %d\n", invalid)
	}
}
