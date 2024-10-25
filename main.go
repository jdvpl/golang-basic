package main

import (
	"fmt"

	"go.mod/excersices"
	"go.mod/operatingSystem"
	"go.mod/variables"
)

func main() {
	variables.ShowIntegers()
	variables.ShowStrings()

	status, text := variables.ConvertToStrings(10)
	fmt.Println(status, text)
	operatingSystem.GetOS()

	number, text := excersices.ConvertToNumber("110")
	fmt.Println(number, text)
}
