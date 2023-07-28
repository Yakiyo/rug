package parser

import (
	"os"

	json "github.com/json-iterator/go"
)

// Default rugfile values
func DefaultRug() RugFile {
	return RugFile{
		Scripts: map[string]string{
			"test": "echo Hello World",
		},
	}
}

func WriteDefault(path string) error {
	def, err := json.MarshalIndent(DefaultRug(), "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, def, 0666)
	return err
}
