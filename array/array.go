package main

import "fmt"

func main() {
	cobaArr := [3]int{1, 2, 3}
	satu := cobaArr[0]
	dua := cobaArr[1]
	tiga := cobaArr[2]
	fmt.Println(satu, dua, tiga)

	s := []string{"hello", "world", ".."}
	s = []string{"1", "2", "4", "5"}
	fmt.Println(s)

	fmt.Println([]string{"coba"})

	var baru1 [5]int
	baru1[1] = 154000000000000
	fmt.Println(baru1)

	// create array with no data signed yet
	// keyw nameofar[leng]Type

	multidimAr := [3][3]string{{"coba1", "coba2", "coba3"},
		{"bisa1", "bisa2", "bisa3"},
		{"mantul1", "mantul2", "mantul3"}}
	fmt.Println(multidimAr[1])

	// dedicated var with fullfiled assigned value
	// namevar = [length]Type{value1,value2,..}
	baruu := []byte{1, 2, 3, 4}
	fmt.Println(baruu)

	// initialize multidim arr
	var arrX [3][2]int
	// arrX[...][0] = {1,2,3}
	// arrX[1] = {2,3,4}
	fmt.Println(arrX)

	arr1 := [3]int{9, 7, 6}
	arr2 := [...]int{9, 7, 6, 4, 5, 3, 2, 4} //with size determined by elipsis(created value after)
	arr3 := [3]int{9, 3, 5}

	fmt.Println(len(arr1), len(arr2), len(arr3))

	arr11 := [...]int{9, 7, 6}
	var cobaz = (arr1 == arr11)
	fmt.Println(cobaz)

	// try to iterate array ?
	for x := 0; x < len(arr2); x++ {
		fmt.Printf("%d\t", arr2[x])
		// println(`oke`)
	}
} //0  ==> equal to; return 0
