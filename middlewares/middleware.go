package middlewares

import "fmt"

func multiply(a, b int) int {
	return a * b
}
func sum(a, b int) int {
	return a + b
}
func rest(a, b int) int {
	return a - b
}

func CustomMiddleware() {
	fmt.Println("Executing custom middleware")
	resutl := operationMiddleware(sum)(10, 5)
	fmt.Println(resutl)
	resutl = operationMiddleware(multiply)(9, 5)
	fmt.Println(resutl)
	resutl = operationMiddleware(rest)(6, 5)
	fmt.Println(resutl)
}

func operationMiddleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Operation started")
		return f(a, b)
	}
}
