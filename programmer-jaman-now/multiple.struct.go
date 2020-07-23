package main

import "fmt"

type Animal interface {
	Speak() string
}

type Howlers struct {
}

func (h Howlers) Speak() string {
	return "HOWWWWWWWWL"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Wolf struct {
	Howlers
}

type Beagle struct {
	Howlers
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	var a Animal
	a = Wolf{}
	fmt.Println(a.Speak())
}
