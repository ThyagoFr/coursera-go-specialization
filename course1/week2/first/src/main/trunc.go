package main

import (
	"fmt"
)

func main() {
	var floatNumber float64
	fmt.Print("Insert a float number: ")
	_, _ = fmt.Scan(&floatNumber)
	fmt.Printf("%d \n", int(floatNumber))
}
