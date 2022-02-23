package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	filename []string
	count    int
}

func main() {
	counts := make(map[string]pair)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%v\t%s\n", n.count, n.filename, line)
		}
	}
}

func countLines(f *os.File, counts map[string]pair) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		l := input.Text()
		there := counts[l]
		there.count += 1
		if !findInSlice(there.filename, f.Name()) {
			there.filename = append(there.filename, f.Name())
		}
		counts[l] = there
	}
}

func findInSlice(filenames []string, check string) bool {
	for _, elem := range filenames {
		if check == elem {
			return true
		}
	}
	return false
}
