package command

import (
	"os"
	"path/filepath"
)

func MoveFileOrDirectory(source, destination string) error {
	destInfo, err := os.Stat(destination)
	if err == nil && destInfo.IsDir() {
		destFileName := filepath.Join(destination, filepath.Base(source))
		return os.Rename(source, destFileName)
	}

	return os.Rename(source, destination)
}
