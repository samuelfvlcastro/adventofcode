package utils

import (
	"bufio"
	"os"
)

func MustFetchInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	input := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
