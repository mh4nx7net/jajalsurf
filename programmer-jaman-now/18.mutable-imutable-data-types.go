package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}
func main() {
	var x []int = []int{3, 4, 5} //just create slice[]int on `mem-map`
	fmt.Println(x)               // pass that in to console
	changeFirst(x)               // dup `the value of` 1st x, pass to another `mem-map` bottom
	fmt.Println(x)               // here, by the mem addr. Actualy the 1st x and 2nd one are competely differ

	/*
			// SLICES => mutable actualy. but...(can't add more length inside)
			var x []int = []int{ // [] => denoted as slices, and followed by datatype's of slices
				10,
				100,
			}
			// x []= 18 // even though you add more them to indexes.. it's immutable !!
			y := x // not duplicated on differ `map-mem` actualy. it's just aliases
			y[0] = 5 // and here we go.. now you can see
			fmt.Println(x, y)


			// MAP => mutable, but... (can add more length inside)
			var x map[string]int = map[string]int{ // map[type]type => denoted as pairing of key and value
				"motor": 10,
				"mobil": 5,
			}
			x["oke"]=108 // it's Mutable data types !!, it will add to most significant index
			y:=x // no coppies actualy, :/ just aliases. dunnowhy
			y["dari-y"]=10 // will add to last significant index
			fmt.Println(x,y)



		// LIST => still mutable, but... (more friendly than two other's)
		// and you can douplicate them on differ `map-memory`
		// remember !! golang is datatype functional language
		// then, it's cannot be `var x [3]int = {3, 4,}`. so must be `var x [3]int = [3]int{3, 4,}`
		// uncomment code below. to see their interaction
		// var x [3]int = {3, 4,}
		var x [3]int = [3]int{3, 4} // yap sangat merepotkan !!
		// var x []int = []int{1,2,}
		y := x
		y[2] = 8 // will duplicate/copy and change on differ `map-mem`
		fmt.Println(x, y)
	*/

}
