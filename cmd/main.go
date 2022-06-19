package main

import (
	"bufio"
	"os"
	"zsh-alias/cmd/alias"
	"zsh-alias/pkg/helpers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	action := helpers.GetUserInput(scanner, "(A)dd or (R)emove alias: ")

	if action == string('A') {
		alias.Add(scanner)
	} else if action == string('R') {
		f := helpers.GetUserInput(scanner, "Enter shell config file name: ")
		alias.List(f)
		alias.Remove(scanner, f)
	}
}
