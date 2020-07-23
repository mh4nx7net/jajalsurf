package main

import (
	"fmt"
)


// multiple closure iteration function :v cool, without passing any param btw
// fungsi ini akan mengembalikan ((int, bool), bool)
func IntClosureIterator(int_data []int) (func() (int, bool), bool){
	// fungsi ini akan terus dijalankan melalui for loop
	// sampai sehabisnya int_data

	var idx int = 0 // default is null
	var data_len = len(int_data)
	return func() (int, bool){
		prev_idx := idx // kind like: []index because idx++
		idx++
		//fmt.Println(prev_idx)
		return int_data[prev_idx], (idx < data_len)
	}, (idx < data_len)
}
func IntCallbackIterator(cb func(int), pdata []int){
	for _, val := range pdata{ // int_data harus berupa list
		cb(val)
	}
}
func main(){

	var sum, from_return_int int = 0, 0

	// iki lumayan mbingungi :v
	// IntClosureIterator() will return both `(int,bool)` and `bool`

	/*argsX []string{x:= 0, x<10, x++}
	for argsX[0] argsX[1]; argsX[2]{
		fmt.Println(x)
	}*/

	for from_return_func, from_return_bool := IntClosureIterator([]int{1,2,3,4,5,6}); from_return_bool; from_return_int, from_return_bool = from_return_func(){
		// pada param terakhir pun(from_return_func)
		// masih diiterasi kembali(golang: mantap bukan ?) dan tentunya
		// memiliki return `from_return_int` dan `from_return_bool`
		// sampai habisnya increment dari fungsi `from_return_func()`
		// sum += from_return_int akan terus dijalankan
		
		sum += from_return_int //same as to `increment` or `x++`
		fmt.Println(sum,from_return_bool,from_return_func)
	}
	reconvene := func(val int){
		sum += val
	}

	IntCallbackIterator(reconvene, []int{2,10})

	//fmt.Println(sum)
}
