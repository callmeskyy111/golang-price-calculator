package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open("prices.txt")
	if err != nil {
		return nil, errors.New("failed to open file!")
	}
	scanner := bufio.NewScanner(file)

	var lines []string //helper variab;e

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file!")
	}

	file.Close()
	return lines, nil

}

// Storing JSON Data in Files
func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create file")
	}
	encoder:=json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("faled to convert to JSON")
	}
	file.Close()
	return  nil
}