package arrays_slices

import "fmt"

var table [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var matrix [20][10]int

// un array es una estructura de datos que almacena una colección de elementos del mismo tipo en go. La longitud de un array es fija y no puede ser modificada una vez que se ha definido.
func ShowArrays() {
	table[7] = 32
	table[2] = 167

	table2 := [10]string{"jiren", "goku", "vegeta", "piccolo", "krilin", "gohan", "trunks", "bulma", "chichi", "goten"}
	fmt.Println(table)
	fmt.Println(table2)
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
	for i := 0; i < len(table2); i++ {
		fmt.Println(table2[i])
	}
	matrix[8][5] = 168
	fmt.Println(matrix)
}
