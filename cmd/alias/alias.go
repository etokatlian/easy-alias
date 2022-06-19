package alias

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"zsh-alias/pkg/helpers"

	"github.com/fatih/color"
)

func Add(scanner *bufio.Scanner) {
	f := helpers.GetUserInput(scanner, "Enter shell config file name: ")
	a := helpers.GetUserInput(scanner, "Enter alias name: ")
	c := helpers.GetUserInput(scanner, "Enter command: ")

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

		lines := helpers.GetLinesFromFile(path)

		lines[len(lines)-1] = lines[len(lines)-1] + "\n"

		output := strings.Join(lines, "\n")

		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := file.WriteString(fmt.Sprintf("alias %s='%s'", a, c)); err != nil {
			log.Println(err)
		}
	}
}

func Remove(scanner *bufio.Scanner, fileName string) {
	a := helpers.GetUserInput(scanner, "Enter alias name: ")

	// stop removal of all aliases if no user entry
	if len(a) == 0 {
		color.Yellow("Nothing Deleted")
		os.Exit(1)
	}

	// set default
	if len(fileName) == 0 {
		fileName = ".zprofile"

	}

	user, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	path := user + "/" + fileName
	lines := helpers.GetLinesFromFile(path)

	for i, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("alias %s", a)) {
			color.Red("Removed %s", line)
			lines[i] = ""
		}
	}

	clean := helpers.DeleteEmpty(lines)

	output := strings.Join(clean, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func List(fileName string) {
	fmt.Println("Remove an alias by name:")

	user, err := os.UserHomeDir()

	if len(fileName) == 0 {
		fileName = ".zprofile"
	}

	if err != nil {
		log.Fatal(err)
	}

	path := user + "/" + fileName
	lines := helpers.GetLinesFromFile(path)

	for _, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("alias")) {
			color.Green(line)
		}
	}
}

func TestFunc(x, y int) int {
	return x + y
}
