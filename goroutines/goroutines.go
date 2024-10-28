package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func SlowName(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}

func ShowGoroutines() {
	fmt.Println("Hello bugs I am here")
	var x string
	fmt.Scanln(&x)
}
