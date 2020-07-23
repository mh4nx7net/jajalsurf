package main

import (
	"fmt"
	"time"
)

func main() {
	expression := 2

	switch expression {
	case 1:
		fmt.Println("diskip")
	case 2:
		fmt.Println("ini yang nampak")
	case 3:
		fmt.Println("ketiga")
	}

	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday() { //expression with some `time` function
	case time.Saturday, time.Sunday:
		fmt.Println("akhir pekan bos")
	case time.Monday:
		fmt.Println("buset andang masuk kerja")
	default:
		fmt.Println("tidak antar keduanya")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("sebelum siang")
	default:
		fmt.Println("setelah siang")
	}

	// try to implement switch inside lambda function as an variable
	// then use them as simple function that would return a some string

	// baca selengkapnya mengenai tipe data interface
	// https://dasarpemrogramangolang.novalagung.com/A-interface-kosong.html
	// https://gobyexample.com/interfaces
	// Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun.
	// Tipe data dengan konsep ini biasa disebut dengan dynamic typing.

	whoAmI := func(i interface{}) { //whoAmI adalah fungsi lambda/temporary, dengan argumen mengambil dari interface/ bisa disbut sbg bypass args
		switch t := i.(type) {
		case bool: //tipe data dapat digunakan sebagai ekspresi
			fmt.Println("ini adalah boolean")
		case int:
			fmt.Println("ini adalah interger")
		default:
			fmt.Println("tipe datamu adalah %T", t)
		}
	}
	whoAmI(true)
}
