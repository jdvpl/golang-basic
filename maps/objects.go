package maps

import "fmt"

// los mapas son una estructura de datos que permite almacenar datos en pares clave-valor, se declaran con la palabra reservada map seguido de los tipos de datos de la clave y el valor
func ShowMaps() {
	countries := make(map[string]string)
	fmt.Println(countries)
	countries["colombia"] = "bogota"
	countries["argentina"] = "buenos aires"
	countries["mexico"] = "ciudad de mexico"
	fmt.Println(countries)
	fmt.Println(countries["colombia"])
	//asignando la capacidad al map
	states := make(map[string]string, 5)
	fmt.Println(states)
	states["Huila"] = "Neiva"
	states["Cundinamarca"] = "Bogota"
	states["Tolima"] = "Ibague"
	states["Caldas"] = "Manizales"
	states["Antioquia"] = "Medellin"
	states["Valle del Cauca"] = "Cali"
	fmt.Println(states)

	championShip := map[string]int{
		"Barcelona":         60,
		"Real Madrid":       48,
		"Manchester United": 61,
		"Juventus":          18,
		"Milan":             39,
	}
	fmt.Println(championShip)
	for team, points := range championShip {
		fmt.Printf("The team %s has %d points\n", team, points)
	}
	//delete a key from the map
	delete(championShip, "Milan")

	points, exist := championShip["Juventus"]
	fmt.Printf("The team Milan has %d points and the team exist = %t\n", points, exist)
}
