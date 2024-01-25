package command

import (
	"errors"
	"fmt"
	"os"
)

func CatCommand(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: cat <filename>")
	}

	fileContent, err := os.ReadFile(args[1])
	if err != nil {
		return err
	}
	fmt.Println(string(fileContent))

	return nil
}
