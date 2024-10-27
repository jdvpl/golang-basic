package functions

import "fmt"

// Un closure es una función que referencia variables libres en su entorno. En el ejemplo anterior, la función anónima retornada por counter captura la variable count y la incrementa cada vez que se llama a increment. Esto permite que la variable count mantenga su estado entre llamadas a la función.
func Clousure() {
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

	// Ejemplo de closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	increment := counter()
	fmt.Println("Closure Counter:", increment()) // 1
	fmt.Println("Closure Counter:", increment()) // 2
	fmt.Println("Closure Counter:", increment()) // 3
}

func Table(value int) func() int {
	number := value
	sequence := 0

	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	table := Table(5)
	fmt.Println(table())
	fmt.Println(table())
	fmt.Println(table())
	for i := 0; i < 10; i++ {
		fmt.Println(table())
	}
}
