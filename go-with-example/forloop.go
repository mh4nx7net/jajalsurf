package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
		// return
	}
	for z := 1; z <= 10; z++ {
		if z%2 == 1 {
			fmt.Println(z, "adalah ganjil")
			break //if ganjil maka langsung return main program

			// kontradiksi dengan fungsi break diatasnya..
			// program executed by order of the mem-cal-func
			continue
		}
		if z%2 == 0 {
			fmt.Println(z, "adalah genap")
			return //jika dapat genap, maka langsung prune program
		}
		fmt.Println("loop sucess")
	}
	fmt.Println("lanjut utama")
}
