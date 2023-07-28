package parser

import (
	"os"

	json "github.com/json-iterator/go"
)

// Read a rug file from a specified path
func ReadRugFile(path string) (RugFile, error) {
	content, err := readFile(path)
	if err != nil {
		return RugFile{}, err
	}
	rug := RugFile{}
	err = json.Unmarshal(content, &rug)
	if err != nil {
		return RugFile{}, nil
	}
	return rug, nil
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}