package fileops

import (
	"errors"
	"fmt"
	"os"
)

func GetStringFromFile(fileName string) (string, error) {
	result, err := os.ReadFile(fileName)
	if err != nil {
		return "0", errors.New("Eror reading file.")
	}
	return string(result), nil
}

func SaveTwoIntsToFile(fileName string, result1 int, result2 int) {
	resultText := fmt.Sprintf("%d , %d", result1, result2)
	os.WriteFile(fileName, []byte(resultText), 0644)
}
