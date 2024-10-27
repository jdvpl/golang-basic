package functions

import "fmt"

func Calculate() {
	// Ejemplo de función anónima para sumar dos números
	sum := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println("Suma:", sum(5, 5))

	// Ejemplo de función anónima para restar dos números
	subtract := func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Println("Resta:", subtract(10, 3))

	// Ejemplo de función anónima para multiplicar dos números
	multiply := func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Println("Multiplicación:", multiply(4, 7))

	// Ejemplo de función anónima para dividir dos números
	divide := func(num1 int, num2 int) float64 {
		if num2 == 0 {
			return 0
		}
		return float64(num1) / float64(num2)
	}
	fmt.Println("División:", divide(20, 4))
}