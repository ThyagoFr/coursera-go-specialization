package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var persons []Person
	var fileName string
	fmt.Scan(&fileName)
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var person Person
		line := strings.Split(scanner.Text(), " ")
		person.fname = line[0][0:20]
		person.lname = line[1][0:20]
		persons = append(persons, person)
	}
	for _, p := range persons {
		fmt.Println("First name:", p.fname, "\t| Last name:", p.lname)
	}
}
