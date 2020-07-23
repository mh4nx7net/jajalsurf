package main

import (
	"bytes"
	"fmt"
)

func main(){
	var slice1 [4]int
	slice1 = [...]int{1,2,3,0} // `...` -> denote as: origin of any value
	var slice2 [4]int

	fmt.Printf("%p\n",&slice1)
	fmt.Printf("%p\n",&slice1)
	alamat := &slice1
	var thisCompareByValue bool = (slice1 == slice2)
	var thisCompareByRefrence bool = (&slice1 == &slice2)
	fmt.Println(thisCompareByValue)
	fmt.Println(thisCompareByRefrence)
	fmt.Printf("%T\n",alamat)
	// array diatas terlihat kurang fleksibel; dan golang jarang menggunakannya
	// sebagai gantinya golang menyediakan slices
	// pendeklarasian mirip array, hanya saja tanpa length `[length]` eksplisit
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)
	// dengan fungsi make() juga dapat digunakan untuk membuat slice
	// func make([]T, len, cap) []T {}

	//var s []byte
	//s = make([]byte, 5, 10)
	// s == []byte{0, 0, 0, 0, 0}
	//s[0] = 100

	s := make([]int, 1, 13)
	fmt.Printf("%v\n",cap(s))
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %v\n", cap(s), len(s), s)
	}
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	is_same := bytes.Compare([]byte{'g','o'}, b[0:1])
	fmt.Println(is_same)
	cek_byte0 := &[]byte{'g'}
	cek_byte1 := &[]byte{'g'}
	fmt.Printf("%v -> %p\n\n",string(b),&b)
	fmt.Printf("compare ptr origin:\n byte0= %p byte1= %p\n\n",cek_byte0, cek_byte1)
	fmt.Printf("compare ptr var   :\n %p <--> %p\n\n",cek_byte1, &cek_byte1)


	// FUN FACT!!!! :V
	// Slicing does not copy the slice's data.
	// It creates a new slice value that points to the original array.
	// This makes slice operations as efficient as manipulating array indices.
	// Therefore, modifying the elements (not the slice itself)
	// of a re-slice modifies the elements of the original slice:

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:] // e == []byte{'a', 'd'}
	e[1] = 'm' // e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
}
