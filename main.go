package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	f := getUserInput(scanner, "Enter shell config file name: ")
	a := getUserInput(scanner, "Enter alias name: ")
	c := getUserInput(scanner, "Enter command: ")

	// set default
	if len(f) == 0 {
		f = ".zprofile"
	}

	if len(a) != 0 && len(c) != 0 {
		user, err := os.UserHomeDir()

		if err != nil {
			log.Fatal(err)
		}

		path := user + "/" + f

		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModePerm)

		if _, err := file.WriteString(fmt.Sprintf("alias %s='%s'\n", a, c)); err != nil {
			log.Println(err)
		}
	}
}

func getUserInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	r := scanner.Text()
	return r
}
