package main

import "fmt"

type rectangle struct {
	length float64
	breadth float64
	color string
}


func main() {
	r1 := rectangle{10,20,"green"}
	fmt.Println(r1)

	r2 := r1 // it will coppie from origin `r1`
	r2.color = "pink"
	fmt.Println(r1,r2)

	var rect1 = rectangle{
		10, 20, "blue",
	}
	rect2 := rectangle{length: 20, breadth: 10, color: "red"}

	if rect1 == rect2{ // compare by pointer/mem
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}

	rect3 := new(rectangle) // mean of refrence
	var rect4 = &rectangle{} // same, but differ map address(origin)
	
	if *rect3 == *rect4 { // compare by value
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
