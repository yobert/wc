package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, name := range os.Args[1:] {
		if err := do(name); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func do(name string) error {
	fh, err := os.Open(name)
	if err != nil {
		return err
	}
	defer fh.Close()

	lines := 0
	words := 0
	chars := 0

	lastword := 0

	var buf [4096]byte

	for {
		n, err := fh.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		for i := 0; i < n; i++ {
			ch := buf[i]
			if ch == '\n' {
				lines++
			}
			if ch == '\n' || ch == '\r' || ch == '\t' || ch == '\v' || ch == ' ' {
				if chars-lastword > 0 {
					words++
				}
				lastword = chars + 1
			}
			chars++
		}
	}

	fmt.Printf(" %d %d %d %s\n", lines, words, chars, name)
	return nil
}
