package main

import (
	"fmt"

	"go.mod/variables"
)

func main() {
	variables.ShowIntegers()
	variables.ShowStrings()

	status, text := variables.ConvertToStrings(10)
	fmt.Println(status, text)
}
