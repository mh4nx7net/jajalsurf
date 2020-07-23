package main

import "fmt"

// `MetaOwn` adalah metode yang akan diteken nantinya
// umumnya/yang paling dekat digunakan antar
// struct yang bermacam2 dg kebutuhan data yang mirip/sama

type MetaOwn interface {
	// mendaftarkan `declared lambda function` sbg kontrak interface MetaOwn
	GetName() string //dengan return string :D
	GetBorn() string

	// karena tiap2 dari kontrak bersifat dinamik(didalam sistem)
	// maka fungsi harus tetap diterapkan scr eksplisit
	// untuk fungsi spesifik yang memiliki nilai return spesifik

	// contoh
	// func *lambda(args1 typeofArg1) signature type{
	// 	return args1.obj
	// }
}

// fungsi kontrak yang turut terdeklarasi scr eksplisit
// mempergunakan/memfungsikan kontrak yang telah dideklarasikan
// dan dilanjut mempointing kontrak ke spesifik struct/method lain2

//membuat fungsi `Getname()` dan `GetBorn()` dengan return nilai `string`
func (user Person) GetName() string {
	// meta `GetName()` dipointing ke `struct Person.Name`
	return user.Name
}

func (user Person) GetBorn() string {
	// meta `GetBorn()` dipointing ke `struct Person.Lahir`
	return user.Lahir
}

// ============================================
// fungsi yang akan dipanggil
// dengan `MetaOwn` sbg args1
func ucapPagi(meta MetaOwn) {
	// mengambil/menggunakan metode `GetName()`
	// dari `meta` yang ada didalam kontrak `MetaOwn`
	fmt.Println("Pagi, Namaku adalah: ", meta.GetName())
}

func dimanaLahir(meta MetaOwn) {
	// konsep sama. cuma dari metode kontrak GetBorn dengan return nilai ttp string
	fmt.Println("Saya lahir di:", meta.GetBorn())
}

// ============================================

// ============================================
// ============================================
// COL (Collection of Field) orang
// yang berisikan data nama,lahir, alamat
type Person struct {
	Name  string
	Lahir string
	// Alamat string
}

// ============================================
// ============================================

func main() {
	var user1 Person // `user1` adalah object/data berjenis Person
	user1.Name = "budi"
	user1.Lahir = "Yogyakarta"

	// fungsi yang akan dipanggil
	ucapPagi(user1)
	dimanaLahir(user1)
}
