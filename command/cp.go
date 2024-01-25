package command

import (
	"errors"
	"io"
	"os"
)

func CpCommand(args []string) error {
	if len(args) < 3 {
		return errors.New("usage: cp <source> <destination>")
	}

	sourceFile, err := os.Open(args[1])
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(args[2])
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
