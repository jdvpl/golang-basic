package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func SlowName(name string, channel1 chan bool) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
	channel1 <- true
}

func ShowGoroutines() {
	fmt.Println("Hello bugs I am here")
	var x string
	fmt.Scanln(&x)
}
