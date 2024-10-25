package bucles

import (
	"fmt"
)

func For() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("--------")
	for j := 0; j < 40; j += 2 {
		if j%3 == 0 {
			fmt.Println(j)
		}
		if j == 20 {
			continue
		}
		if j == 30 {
			break
		}
		fmt.Println(j)
		fmt.Println("------------")
	}
}
