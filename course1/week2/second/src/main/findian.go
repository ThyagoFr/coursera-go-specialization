package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text: ")
	text, _ := scanner.ReadString('\n')
	var found bool
	if string(text[0]) == "i" && string(text[len(text)-2]) == "n" {
		for _, c := range text {
			if string(c) == "a" {
				found = true
				break
			}
		}
	}
	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
