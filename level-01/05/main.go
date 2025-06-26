package main

import "fmt"

type myType int
var x myType
var y int

func main() {
	// Primeira Parte
	fmt.Printf("x = %v\n", x)
	fmt.Printf("O tipo de x é %T\n", x)
	x = 42
	fmt.Printf("Agora, x vale %v\n", x)

	fmt.Println("----------------------")

	// Segunda Parte
	y = int(x)
	fmt.Printf("y = %d\n", y)
	fmt.Printf("O Tipo de y é %T\n", y)
}
