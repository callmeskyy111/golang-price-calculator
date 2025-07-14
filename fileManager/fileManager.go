package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutPutFilePath string
}

func(fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	scanner := bufio.NewScanner(file)

	var lines []string //helper variab;e

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

// Storing JSON Data in Files
func(fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutPutFilePath)
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

// Constructir Fx
func New(inputPath, outputPath string)FileManager{
return FileManager{
	InputFilePath: inputPath,
	OutPutFilePath: outputPath,
}
}