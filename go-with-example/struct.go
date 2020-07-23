// struct areform to/as grouping data together
// struct are typed collection of fields
// macam2 koleksi lapangan terdeklarasi{eksplisit}

// struct dari person akan berisi tentang tinggi,usia
// struct dari identitas akan berisi tentang alamat, nama

package main

type person struct {
	name string
	age  int
}

// metode untuk membuat objek(newPerson) baru
// dengan melakukan passing `name`` sbg args1 dan
// tipe metode menggunakan pointer *person
func newPerson(name string) *person {

}
