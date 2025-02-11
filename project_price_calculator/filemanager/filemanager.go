package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (f FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(f.OutputFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	time.Sleep(time.Second * 3)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("error encoding JSON")
	}
	return nil
}

func (f FileManager) ReadLines() (*[]string, error) {
	lines := []string{}
	file, err := os.Open(f.InputFilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textPrice := scanner.Text()
		lines = append(lines, textPrice)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return &lines, nil
}
