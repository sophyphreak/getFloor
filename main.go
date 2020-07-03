package main

import (
	"fmt"

	"./userinterface"
)

func main() {
	date := userinterface.GetDate()
	fmt.Println(date)
}
