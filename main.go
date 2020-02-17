package main

import "fmt"

func main() {
	//Declarado el array
	numeros := [3]float64{71.8, 56.2, 89.5}
	//Se declara el sumatorio
	var sum float64 = 0
	//Se declara el for, donde para el for...range necesita devolver dos valores que pueden ser ignorados o no
	for _, numero := range numeros {
		//Se sumn todos los números del arreglo
		fmt.Println(numero)
		sum += numero
	}
	//Realizar el promedio
	numeroElementos := float64(len(numeros))
	fmt.Printf("el promedio es: %0.2f\n", sum/numeroElementos)
}
