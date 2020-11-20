package main

import "fmt"

func main() {
	var number int = 42
	number += 5
	n := 42
	fmt.Println(number, n)

	floatSum := float32(number) + 5.5 // Go is strongly typed
	fmt.Printf("%T\n", floatSum)      // prints float32

	a := [4]string{"Hello world", "I am", "a string", "array"} // cannot be resized
	fmt.Println(len(a), a)

	s := []string{"Hello world", "I am", "a string", "slice"} // resized automatically
	fmt.Println(len(s), cap(s), s)
	s = append(s, "nice")
	fmt.Println(len(s), cap(s), s) // capacity doubled

	greet("Master Kenobi")

	greetAll("Master Kenobi", "Anakin Skywalker", "Chancellor Palpatine")

	palpatine, vader := sith()
	fmt.Println(palpatine, vader)

	anakin := Jedi{
		Rank: "Jedi Knight",
		Age:  23,
		Name: "Anakin Skywalker",
	}
	obiWan := Jedi{
		Rank:    "Jedi Master",
		Age:     38,
		Name:    "Obi-Wan Kenobi",
		Padawan: &anakin,
	}

	fmt.Println(obiWan, anakin)

	obiWan.SayHello(anakin)
}

func greet(name string) int {
	fmt.Printf("Greetings %s!\n", name)
	return 0
}

func greetAll(names ...string) {
	for _, name := range names {
		fmt.Printf("Greetings %s!\n", name)
	}
}

// go supports multiple return values
func sith() (string, string) {
	return "Darth Sidious", "Darth Vader"
}

type Jedi struct {
	Rank    string
	Age     int
	Name    string
	Padawan *Jedi
}

func (j Jedi) SayHello(other Jedi) {
	fmt.Printf("%s says: \"Hello there, %s!\"", j.Name, other.Name)
}
