package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	person := make(map[string]string)
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	text, _ := scanner.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	person["name"] = text
	fmt.Print("Enter your address: ")
	text, _ = scanner.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	person["address"] = text
	jsonPerson, _ := json.Marshal(person)
	fmt.Println(string(jsonPerson))

}
