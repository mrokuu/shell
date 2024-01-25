package command

import (
	"os"
	"time"
)

func Touch(args []string) error {
	filePath := args[1]
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	now := time.Now()
	if err := os.Chtimes(filePath, now, now); err != nil {
		return err
	}
	return nil
}
