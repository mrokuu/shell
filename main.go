package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"shell/command"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "ls":

		return command.ListDirectory(".")
	case "rm":
		if len(args) < 2 {
			return errors.New("usage: rm <file/directory>")
		}
		return command.RemoveFileOrDirectory(args[1])
	case "mkdir":
		if len(args) < 2 {
			return errors.New("usage: mkdir <directory>")
		}
		return command.MakeDirectory(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
