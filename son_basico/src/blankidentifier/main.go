package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	res, _ := http.Get("http://google.com.br")
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s --", body)

}
