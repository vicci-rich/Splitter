package pathutils

import (
	"os/exec"
	"path/filepath"
)

func GetCurrentPath(filePath string) (string, error) {
	var path string
	file, err := exec.LookPath(filePath)
	if err != nil {
		return "", err
	}
	path, err = filepath.Abs(file)
	return path, nil
}

func GetParentPath(filePath string) (string, error) {
	s, err := GetCurrentPath(filePath)
	if err != nil {
		return "", err
	}
	path := filepath.Dir(s)
	return path, nil
}
