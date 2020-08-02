package main

import (
	"fmt"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func (a *Animal) New(eat, move, speak string) {
	a.eat = eat
	a.move = move
	a.speak = speak
}

func (a *Animal) Eat() {
	fmt.Println(a.eat)
}

func (a *Animal) Move() {
	fmt.Println(a.move)
}

func (a *Animal) Speak() {
	fmt.Println(a.speak)
}

func main() {

	var cow Animal
	var snake Animal
	var bird Animal

	cow.New("grass", "walk", "moo")
	bird.New("worms", "fly", "peep")
	snake.New("mice", "slither", "hsss")

	animals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	for {
		var animal, info string
		fmt.Print(">")
		fmt.Scanf("%s %s", &animal, &info)
		animalSelected := animals[animal]
		switch string(info) {
		case "eat":
			animalSelected.Eat()
		case "move":
			animalSelected.Move()
		case "speak":
			animalSelected.Speak()
		}
	}
}
