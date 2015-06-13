package util

import (
	"bufio"
	"os"
	"strings"
)

// ReadFileAsLines reads an entire file into memory as a slice of its non-empty lines. From: http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write
func ReadFileAsLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " \t\r\n")
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines, scanner.Err()
}
