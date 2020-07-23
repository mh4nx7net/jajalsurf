package main

import(
	"fmt"
	"reflect"
)

type rectangle struct {
	length float64
	breadth float64
	color string
}

func main() {
	var rect1 = rectangle{
		10,
		20,
		"red",
	}
	fmt.Println(reflect.TypeOf(rect1)) // main.rectangle
	fmt.Println(reflect.ValueOf(rect1)) // {{{10 20 red}

	rect2 := rectangle{
		length: 10.1,
		breadth: 20.1,
		color: "blue",
	}
	fmt.Println(reflect.TypeOf(rect2))
	fmt.Println(reflect.ValueOf(rect2).Kind()) // .Kind() with type bunch of value


	rect3 := new(rectangle)
	fmt.Printf("%T\n",rect3) // *main.rectangle
	fmt.Println(reflect.TypeOf(rect3)) // *main.rectangle
	fmt.Println(reflect.ValueOf(rect3).Kind()) // pointer
	fmt.Println(reflect.ValueOf(&rectangle{}).Kind()) // same kind as -> top of


}
