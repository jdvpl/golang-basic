package main

import (
	"fmt"

	"go.mod/variables"
)

func main() {
	variables.ShowIntegers()
	fmt.Println(variables.Amount)
	variables.ShowStrings()
	fmt.Println(variables.Amount)
}
