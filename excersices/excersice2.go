package excersices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TableMultiplication() {
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var err error
	for {
		fmt.Println("Enter the number")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}

	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
