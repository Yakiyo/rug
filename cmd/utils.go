package cmd

import (
	"errors"
	"os"
)

// Check if file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
