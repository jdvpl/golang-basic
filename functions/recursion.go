package functions

import (
	"fmt"
	"math/big"
)

// recursion es una función que se llama a sí misma hasta que se cumple una condición de salida. En este caso, la función se llama a sí misma para calcular el factorial de un número.
func Pow(value int) {
	if value > 1000000000 {
		return
	}
	fmt.Println(value)
	Pow(value * 2)

}

func Factorial(value int64) *big.Int {
	if value == 0 {
		return big.NewInt(1)
	}
	n := big.NewInt(value)
	return n.Mul(n, Factorial(value-1))
}
