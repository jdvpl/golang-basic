package scannindata

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func GetingDataFromKeyboard() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the first number")

	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Not a number " + err.Error())
		}
	}

	fmt.Println("Enter the second number")

	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Not a number " + err.Error())
		}
	}

	fmt.Println("Enter the legend")

	if scanner.Scan() {
		legend = scanner.Text()
	}
	fmt.Println(legend, number1*number2)
}
