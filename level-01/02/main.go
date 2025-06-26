package main

import "fmt"

var x int
var y string
var z bool

func main() {
	// Valores Zero
	fmt.Println("x = ", x) // 0
	fmt.Println("y = ", y) // vazio
	fmt.Println("z = ", z) // false

	// Solução com Print Format:
	fmt.Printf("x = %d | y = %s | z = %t\n", x, y, z)
}
