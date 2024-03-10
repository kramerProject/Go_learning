package scope

import "fmt"

// var b int
// var c, d string = "Var C", "Var D"
var y int = 20

// func main() {
// 	a := 10
// 	b = 22
// 	fmt.Printf("a: %v \n b: %v \n", a, b)
// 	fmt.Printf("c: %v \n d: %v \n", c, d)
// 	fmt.Printf("Z: %v \n", z)
// }

func PrintY() {
	fmt.Println(y)
}
