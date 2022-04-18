// Exercise 1.4
// Dup4 prints the count and text of any lines that appear more than once
// and the files that the line appears in. It takes input from Stdin or a
// list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts_per_file := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts_per_file, "os.Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts_per_file, arg)
			f.Close()
		}
	}

	for line, per_file := range counts_per_file {
		files := ""
		sep := ""
		total := 0
		for filename, count := range per_file {
			files += sep + filename
			sep = " "
			total += count
		}

		if total > 1 {
			fmt.Printf("%d\t%s\t`%s`\n", total, files, line)
		}
	}
}

func countLines(f *os.File, counts_per_file map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts_per_file[line] == nil {
			counts_per_file[line] = make(map[string]int)
		}
		counts_per_file[line][filename]++
	}
	// NOTE: Ignoring potential errors from input.Err()
}
