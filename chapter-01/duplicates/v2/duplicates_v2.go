package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	occurrences := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {

			file, err := os.Open(arg)

			// Error handling.
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplicates_v2.go: %v\n", err)
				// Try next file
				continue
			}

			// No error. Try to count lines.
			countLines(file, counts, occurrences)
			file.Close()
		}
	}

	for line, count := range counts {
		if count > 1 {
			if occurrences != nil {
				files := occurrences[line]
				fmt.Printf("Found duplicate: %d\t%q\t found in files:%s\n", count, line, files)
			} else {
				fmt.Printf("Found duplicate: %d\t%q\n", count, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, occurrences map[string][]string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if occurrences != nil {
			rememberOccurrence(occurrences, line, f.Name())
		}
		counts[line]++
	}
	// Normally, we should check for potential erros as soon as Scan() returns
	// false by calling Err() (will be nil in case of EOF).
	// We ignore this here for simplicity.
}

func rememberOccurrence(occurrences map[string][]string, line string, fileName string) {
	if !lineAlreadyMetInFile(fileName, occurrences[line]) {
		occurrences[line] = append(occurrences[line], fileName)
	}
}

func lineAlreadyMetInFile(currentFileName string, files []string) bool {
	for _, fileName := range files {
		if currentFileName == fileName {
			return true
		}
	}
	return false
}
