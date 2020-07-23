package main

import "fmt"

type rectangle struct{
	length float64
	breadth	float64
	color string
}

func main() {
	oke := rectangle{
		10.5,
		24.10,
		"red"}
		fmt.Println(oke)
	}
