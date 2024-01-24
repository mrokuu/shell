package command

import "os"

func RemoveFileOrDirectory(path string) error {
	return os.RemoveAll(path)
}
