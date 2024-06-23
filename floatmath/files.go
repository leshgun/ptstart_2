package floatmath

import (
	"bufio"
	"os"
)

func ReadFile(path string) ([]string, error) {
	// Open file for reading
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	// Convert lines into array
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {	
		lines = append(lines, scanner.Text())
	}
	// The file is no longer needed.
	file.Close()
	return lines, nil
}
