package utils

import (
	"bufio"
	"os"
)

func FetchInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, nil
	}
	defer file.Close()

	input := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}
