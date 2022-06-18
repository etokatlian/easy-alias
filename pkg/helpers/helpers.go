package helpers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetLinesFromFile(path string) []string {
	fBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fBytes), "\n")
	return lines
}

func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func GetUserInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	r := scanner.Text()
	return r
}
