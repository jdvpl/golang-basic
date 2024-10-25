package main

import (
	"fmt"

	operatinSystem "go.mod/os"
	"go.mod/variables"
)

func main() {
	variables.ShowIntegers()
	variables.ShowStrings()

	status, text := variables.ConvertToStrings(10)
	fmt.Println(status, text)
	operatinSystem.GetOS()
}
