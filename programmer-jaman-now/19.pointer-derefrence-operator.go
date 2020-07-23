package main

import "fmt"

func changeVal(str *string){
	*str = "change new"
}

//func changeVal2(str string){
//	str = "change new"
//}
func main() {
	// & => stand for get the address/map pointer of memory
	// * => stand for pointer based/derefrences

	toChange := "coba"
	fmt.Println(toChange)
	changeVal(&toChange)
	fmt.Println(toChange)
	//changeVal2(toChange)


	/*
	x := 7 // create init 7 of new clean mem addr
	y := &x // y is address / refrences to `x memory location`
	fmt.Println(x,y)
	*y = 11 // leads to value of declared `y => x memory location`
	fmt.Println(x,y)
	*/
}
