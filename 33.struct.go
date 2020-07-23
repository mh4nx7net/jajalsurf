package main

import "fmt"

// named-struct example
type User struct {
	name        string
	age         int
	bankBalance float32
}
type coba struct {
	// type cobaa struct{
	// 	nilai int
	// }
}

func main() {

	// anonymous-struct example
	anonStruct := struct {
		nama string
		usia int8
	}{
		nama: "Andri",
		usia: 22,
	}
	fmt.Println(anonStruct)

	a := coba{}
	fmt.Println(a)

	user1 := User{
		"Nindi",
		22,
		30,
	}
	user2 := User{
		name:        "Mohit",
		age:         24,
		bankBalance: 100.0,
	}
	user3 := new(User)
	*user3 = User{"namo", 12, 11}
	fmt.Println(user1, user2, user3)
}
