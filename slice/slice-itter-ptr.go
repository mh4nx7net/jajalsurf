package main

import "fmt"


type Dog struct{
	Name string
	Age int
}
func main(){
	Jakie := Dog{
		Name: "Jakie",
		Age: 2,
	}
	fmt.Printf("%p\n", &Jakie)
	Sammy := Dog{
		Name: "Sammy",
		Age: 1,
	}
	fmt.Printf("%p\n", &Sammy)

	// dogs is literal list of `ptr dog struct`
	// that form by `ref of &Jakie and &Sammy`
	dogs := []*Dog{&Jakie, &Sammy}
	//fmt.Println(dogs)

	for _, dog := range dogs {
		fmt.Printf("Name: %s\tAge: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n\n", dog)
	}
}
