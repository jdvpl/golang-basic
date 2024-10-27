package arrays_slices

import "fmt"

// un Slice es una estructura de datos que almacena una colección de elementos del mismo tipo en go. A diferencia de los arrays, la longitud de un slice es dinámica y puede ser modificada en tiempo de ejecución.

var tableSlice []int = []int{2, 5, 4}
var array [10]int = [10]int{8, 4, 6, 9, 12, 68, 9, 0}

func ShowSlice() {
	fmt.Println(tableSlice)
	portion := array[2:5]
	fmt.Println(portion)
	portion = array[5:]
	fmt.Println(portion)
	portion = array[:5]
	fmt.Println(portion)
	portion = array[:]
	fmt.Println(portion)
}

// La función append() se utiliza para agregar elementos a un slice. Si el slice tiene capacidad suficiente, los elementos se agregan al slice original. Si no, se crea un nuevo slice con una capacidad mayor y se copian los elementos del slice original al nuevo slice.
// la funcion make reserva memoria para un slice y devuelve un slice con la longitud y la capacidad especificadas.
func CapacitySlice() {
	elements := make([]int, 8, 10)
	fmt.Printf("Len %d, Capacity %d \n", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Len %d, Capacity %d, Elements %d \n", len(nums), cap(nums), nums)
}
