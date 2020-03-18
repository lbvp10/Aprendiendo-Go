package main

import "fmt"

func main() {
	/*********************************
	Arrays
	 **********************************/
	Colores := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(Colores);

	/*********************************
	slice
	**********************************/
	slice := make([]int, 3, 4)
	fmt.Println(slice);
	fmt.Printf("len %d - cap %d\n", len(slice), cap(slice));

	//La  capacidad de los slice se x2 cuando se añade un nuevo elemnto que desborsa la  misma
	slice = append(slice, 20)
	slice = append(slice, 21)
	slice = append(slice, 22)
	fmt.Println(slice);
	fmt.Printf("len %d - cap %d\n", len(slice), cap(slice));

	//Diferentes formas de incializar un slice
	var numeros1 []int
	numeros2 := []int{1, 2, 3, 4, 5, 6}
	numeros3 := make([]int, 5)
	fmt.Println(numeros1) // []
	fmt.Println(numeros2) // [1 2 3 4 5 6]
	fmt.Println(numeros3) // [0 0 0 0 0]

	var letras [2]string
	fmt.Println(letras) // []

	//tareas con slice
	newSlice := slice[4:6]
	fmt.Println(newSlice) // [21 22]
	//unir dos slice
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println(append(slice1, slice2...))
}
