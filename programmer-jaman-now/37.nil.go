// pengantar: object yang belum diinisialisasi akan bernilai null atau nil
// untuk golang. saat pertamakali variable dibuat maka akan langsung memiliki default value
// sesuai dengan assignment tipe data yang digunakan
// golang tetap memiliki data nil(data kosong)
// dan hanya bisa digunakan pada interface, function, map slice, pointer dan channel

// praktik: dapat berguna untuk mengecek data didalam map, slice, dsb

package main

import "fmt"

func NewMap(name string) map[string]string {
	// jika name kosong > return nil
	// jika terisi > membuat map baru

	// map[type-key]type-value

	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// =========[ simpel nya begini ]==========
// ========================================
// coba := {
//	 key: name,
// }
// coba := NewMap(val, key){
// 	return map[string]string{
// 		key:val,
// 	}
// }

// func cobaNil(){
// 	return nil
// }
// ========================================
// =========[ simpel nya begitu ]==========

func main() {
	// cobaNil()
	var coba, lain map[string]string

	lain = map[string]string{}
	lain["oke"] = "oke dan bisa"
	fmt.Println(lain["oke"])

	// cara vertikal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	// cara horizontal
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken1["januari"], chicken2["januari"])

	// coba = NewMap("val1", "key1")
	// coba = NewMap("val2", "key2")

	fmt.Println(coba["key1"])
	fmt.Println(coba["key2"])

	fmt.Println("sukses")

	// contoh untuk pengecekan apakah ada `value` didalam map
	var person map[string]string = NewMap("coba")

	if person == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person)
	}
}
