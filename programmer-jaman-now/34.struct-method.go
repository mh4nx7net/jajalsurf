// STRUCT FUNCTION atau STRUCT METHOD
// struct atau tipe data yang terdiri dari kumpulan data
// bisa digunakan sebagai param pada function
// menambahkan method kedalam struct, guna membuat struct memiliki modul function

package main

import "fmt"

type Pelanggan struct {
	Name, Address string
	Age           int
}

// conventional way
func sayHi(pelanggan Pelanggan, name string) {
	fmt.Println("paKabar", name, ". Namaku adalah", pelanggan.Name)
}

// struct-method way
func (pelanggan Pelanggan) bukaSahur(name string) { // `bukaSahur('param1')` adalah new entitled func
	fmt.Println("Sahur gan", name, "apakah agan", pelanggan.Name, "mau sahur")
}

func main() {
	var samsul Pelanggan
	samsul.Name = "Samsul"
	samsul.Address = "Jumoyo"
	samsul.Age = 31

	fmt.Println(samsul)

	sayHi(samsul, "samsul-1")       //conventional way ==> struct dimasukan kdlm function dan dibypass
	samsul.bukaSahur("samsul-punk") //struct-method way ==> struct seolah2 memiliki function

	// pendeklarasian lebih efisien
	pungki := Pelanggan{
		Name: "Pungki",
	}

	pungki.bukaSahur("pungki-rock")
}
