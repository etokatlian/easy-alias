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
		alias.AddAlias(scanner)
	} else if action == string('R') {
		alias.RemoveAlias(scanner)
	}

}
