package command

import "os"

func MakeDirectory(dir string) error {
	return os.Mkdir(dir, os.ModePerm)
}
