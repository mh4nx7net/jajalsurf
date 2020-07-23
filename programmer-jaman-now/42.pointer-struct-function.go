// struct method: meskupun method dapat diakses melalui struct.
// data didalamnya tetap akan berupa `pass by value`
// penggunaan pointer pada `struct method` sangat direkomendasikan guna
// meminimalisir penggunaan memory ketika melakukan pemanggilan method berulang

package main

import "fmt"

type Information struct {
	person Person
	detail Detail
}
type Person struct {
	Name string
	Age  int
}
type Detail struct {
	Address string
	Email   string
}

func (man *Information) pushPerson() {
	man.person.Name = "Pak. " + man.person.Name
}
func (more *Information) pushDetail() {
	more.detail.Address = "Jl." + more.detail.Address
	more.detail.Email = more.detail.Email + "@kk.com"
}

func (turnMe *Information) NewPush() {
	turnMe.pushPerson()
	turnMe.pushDetail()
}

func main() {
	user1 := Information{
		Person{"Rumi", 20},
		Detail{"pandak", "rumi-id"},
	}

	user1.NewPush()
	fmt.Println("Selamat Pagi", user1.person.Name)
	fmt.Println("umur:", user1.person.Age)
	fmt.Println("addr:", user1.detail.Address)
	fmt.Println("email:", user1.detail.Email)

	// ==============================================
	// ==============================================

	emplo8 := &Empl{
		fName: "oke"
		lName: "bisa"
		age:   55,
	}

}
