package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func getNameInput() string {
start:
	fmt.Print("Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	r, _ := regexp.Compile("^[\\w\\.,\\ ]+$")
	if !r.MatchString(name) {
		fmt.Println("Name must be letters, numbers, period, comma, and not empty")
		goto start
	}
	return name
}
