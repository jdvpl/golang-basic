package main

import (
	"go.mod/implementation"
	"go.mod/models"
)

func main() {
	// variables.ShowIntegers()
	// variables.ShowStrings()

	// status, text := variables.ConvertToStrings(10)
	// fmt.Println(status, text)
	// operatingSystem.GetOS()

	// number, text := excersices.ConvertToNumber("110")
	// fmt.Println(number, text)
	// scannindata.GetingDataFromKeyboard()
	// text := excersices.TableMultiplication()
	// files.AddDataToFile(text)

	// files.ReadFile()
	// fmt.Println(functions.Factorial(200))
	// maps.ShowMaps()
	// users.ShowUsers()
	Juan := new(models.Man)
	Gina := new(models.Woman)
	implementation.HumanBreathing(Juan, "Juan")
	implementation.HumanBreathing(Gina, "Gina")
}
