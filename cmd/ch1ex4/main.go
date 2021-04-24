package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	seenIn := make(map[string]string)
	files := os.Args[1:]
	//files = append(files, "test.txt")

	// The debugger uses a different cwd by default
	fmt.Println(os.Getwd())

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				//os.Exit(1)
			}

			countLines(f, counts)

			for item, count := range counts {
				if count > 1 {
					seenIn[item] = arg
				}
			}

			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, seenIn[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
