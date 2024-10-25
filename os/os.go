package operatinSystem

import (
	"fmt"
	"runtime"
)

func GetOS() {
	if os := runtime.GOOS; os == "darwin" || os == "linux" {
		fmt.Println("OS X")
	} else {
		fmt.Println("Not OS X")
	}
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Printf("%s.\n", os)
	}
}
