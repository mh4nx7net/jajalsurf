package main

import "fmt"

type Dog struct{
	Name string
	Age int

}

func main(){
	okk := [10]int{1,2,3,4,5,6,7,8,9}
	t := okk[6:9]
	sl1 := make([]int, 6)
	sl2 := append(sl1, t...)
	fmt.Println(sl1, sl2, okk, t)

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	z := append([]int{}, append(x, y...)...)
	fmt.Println(z)
	a := []int{1, 2}
	a = append(a, 3, 4) // a == [1 2 3 4]
	slice := append([]byte("Hello "), "world!"...)
	fmt.Println(a,string(slice))

	Jakie := Dog{
		Name: "Jakie",
		Age: 2,
	}
	//fmt.Printf("map Jakie:\t%p\n", &Jakie) // with their own ptr
	Sammy := Dog{
		Name: "Sammy",
		Age: 3,
	}
	//fmt.Printf("map Sammy:\t%p\n", &Sammy) // with their own ptr too


	// []anyType_that_can_itter_byList{option1, option2, etc}
	// followed by repetition

	// copy the value and put it on
	// to []Dog{list} struct
	// so the map memory it'self would be differ at all
	dogs_intheHouse := []Dog{Jakie,Sammy}
	// result: Addr: 0xc0000a0020 Addr: 0xc0000a0020 |-> kind as same
	// to fix the issue if dogs_intheHouse would stand with same value of ptr
	// just change the type of []Dog{} --> which stand as value based
	// to []*Dog{} --> which stand as ptr based

	// fmt.Printf("map Jakie:\t%p\n", &dogs_intheHouse[0])
	// fmt.Printf("map Sammy:\t%p\n", &dogs_intheHouse[1])


	// for-range would create clean map-memory
	// to read parsed dogs_intheHouse data's inside
	for _, dogs := range dogs_intheHouse{
		// for above would equal to:
		// var dogs = range dogs_intheHouse

		fmt.Printf("Name: %s\t Age: %d\n", dogs.Name, dogs.Age)
		fmt.Printf("Addr: %p\n\n", &dogs)
		// even though here... also diffrenent
		// not sourcing by ptr, but coppied by value
	}


	// by mean: allDogs would create a list
	// that form by struct of `Dog{}` pointer
	allDogs := []Dog{} //list of `structed ptr` of the `Dog{}`
	for _, dog := range dogs_intheHouse{
		// each of `[]dogs_intheHouse` itself is conform by same ptr
		allDogs = append(allDogs, dog)
	}
	for _, dog := range allDogs{
		fmt.Printf("Name: %s\tAge: %d\n", dog.Name, dog.Age)
	}

}
