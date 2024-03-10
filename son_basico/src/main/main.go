package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u := uuid.NewV4()
	fmt.Println("Hello World %s\n", u)
}
