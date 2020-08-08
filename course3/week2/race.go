package main

import (
	"fmt"
	"time"
)

var value int

func increase1() {
	fmt.Println("Go Routine 1")
	value++
}

func increase2() {
	fmt.Println("Go Routine 2")
	value++
}

func main() {
	go increase1()
	go increase2()
	time.Sleep(2 * time.Second)

	// Run "go run -race race.go"
	// Race conditions = When 2 or mores threads try to access the same state(in this case the variable "value") simultaneously
	// The race conditions happens in this case because the go routines shares the same variable(value) simultaneously
}
