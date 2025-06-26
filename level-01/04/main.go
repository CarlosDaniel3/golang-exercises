package main

import "fmt"

type myType int
var x myType

func main() {
	fmt.Printf("x = %v\n", x)
	fmt.Printf("Tipo de x = %T\n", x)
	x = 42
	fmt.Printf("Agora o x vale %v\n", x)
}