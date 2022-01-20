package main

import (
	"fmt"
)



func main() {
	favoriteSport := "Box"
	switch  favoriteSport {
		case "Soccer":
			fmt.Printf("Case %v", favoriteSport)
		case "Box":
			fmt.Printf("Case %v", favoriteSport)
		case "Surf":
			fmt.Printf("Case %v", favoriteSport)
		default:
			fmt.Printf("Not my thing")
	}
}