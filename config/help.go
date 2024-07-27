package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetCurrentPathRelative(path string) string {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("error getting current working directory: %v", err))
	}

	return filepath.Join(rootDir, path)
}
