// pointer pada param function berguna ketika kita
// ingin membuat function yang dapat mengubah data asli pada param-func
// secara default golang menggunakan `pass by value`
// artinya data yang dikirimkan adalah berupa `nilai`
// implement: gunakan pointer sbg param (*nilai_pointer) untuk melakukan `pass by mem-map`
package main

import "fmt"

type Address struct{
	Kota, Provinsi, Country string
}

func ChangeCountry(addr *Address){ //addr adalah pointer dari Address
	addr.Country = "indonesia" //as default
}
func main(){
	// initial burhan address
	burhanAddr := Address{
		"Palu", "Sulawesi","",
	}


	pointerMap_burhanAddr := &burhanAddr
	ChangeCountry(pointerMap_burhanAddr)
	fmt.Println(burhanAddr)

}