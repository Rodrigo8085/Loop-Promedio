package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//Se abre el documento que se va a leer
	file, err := os.Open("datos.txt")
	//Si existe algun error los reporta
	if err != nil {
		log.Fatal(err)
	}
	//Crea un nuevo scanner para ese nuevo archivo
	scanner := bufio.NewScanner(file)
	// loop le linea por linea del archi y loa imprime
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	//Cierra el docuemtno para liberar recursos
	err = file.Close()
	//Si hay un error cuando cierra el documento
	if err != nil {
		log.Fatal(err)
	}
	//Si hay algun error scaneando las lineas
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
