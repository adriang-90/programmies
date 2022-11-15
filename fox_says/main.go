package main

import "fmt"

// function return a sound that given pet is doing
func petSound(pet string) string {
	var sound string
	switch {
	case pet == "dog":
		sound = "woof"
	case pet == "cat":
		sound = "meow"
	case pet == "snake":
		sound = "ssss"
	default:
		sound = "No idea what that pet is"
	}
	return sound
}

func main() {
	fmt.Println("Dog says ", petSound("dog"))
	fmt.Println("Cat says ", petSound("cat"))
	fmt.Println("Snake says ", petSound("snake"))
	fmt.Println("Fox says ", petSound("fox"))

}
