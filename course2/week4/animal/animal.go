package main

import "fmt"

type Animal interface {
	Move()
	Eat()
	Speak()
	GetName() string
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Cow) GetName() string {
	return c.name
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (b Bird) GetName() string {
	return b.name
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func (s Snake) GetName() string {
	return s.name
}

func main() {

	var animals []Animal
	for {
		var commandOne string
		var param1, param2, param3 string
		fmt.Print(">")
		fmt.Scanf("%s", &commandOne)
		if commandOne == "newanimal" {
			fmt.Scanf("%s %s %s", &param1, &param2, &param3)
			fmt.Println(param3)
			switch param3 {
			case "cow":
				var cow Cow
				cow.name = param2
				animals = append(animals, cow)
			case "snake":
				var snake Snake
				snake.name = param2
				animals = append(animals, snake)
			case "bird":
				var bird Bird
				bird.name = param2
				animals = append(animals, bird)
			}
			fmt.Println(animals)
			fmt.Println("Created it!")
		} else {
			fmt.Scanf("%s %s %s", &param1, &param2, &param3)
			for _, animal := range animals {
				if animal.GetName() == param2 {
					switch param3 {
					case "eat":
						animal.Eat()
					case "speak":
						animal.Speak()
					case "move":
						animal.Move()
					}
					break
				}
			}
		}
	}

}
