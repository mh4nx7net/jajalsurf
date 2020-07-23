// golang sebagai function PL
// dalam OOP, memiliki `the most root of` data parent
// dan dapat terbaca/diimplementasikan oleh branch object dibawahnya
// contoh java:="java.lang.Object" atau kotlin:="Any"
// dalam golang maka digunakan empty interface sebagai pengganti/alternative
// untuk menangani masalah ini

// empty interface is interface's that dont have any declared method
// thatould make all of a data type would be/as implemented on them

// contoh interface kosong:
	// fmt.Println(a ...interface{}) //fungsi .Println merupakan empty interface yang dpt diisi mcm2 tipedata
	// panic(v interface{})
	// recover() interface{}
	// dan lain2

package main

import "fmt"

// type IfaceKosong interface{
// 	// interface kosong adalah kontrak yang didalamnya tidak ada apa2
// 	// seakan2 semua tipe data golang mengikuti kontrak interface kosong
// }


// fungsi Ups menggunakan return value `interface kosong`
// meskipun data1 null. maka tetap akan berlaku
// karena data1-int mengikuti kontrak interface terkait
// mengapa spt tidak terjadi apa2, karena kosong
func Ups(data1 int, Apasaja interface{}) interface{}{
	fmt.Println(Apasaja)
	if data1==1{
		return 1
	}else if data1==2{
		return true
	}else{
		return "gagal"
	}
}

func main(){
	// var data int = Ups(1) //error: interface kosong tidak bisa ditugaskan sbg int pd var data
	var data interface{} = Ups(1,"benar") //tidak lagi error karena nilai return sm2 sesuai, yaitu iface-kosong

	fmt.Println(data)
	// Ups()
}