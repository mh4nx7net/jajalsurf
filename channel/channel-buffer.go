package main

import (
	"fmt"
	"reflect"
)

func main() {
	// channel basic syntaxist
	// only_ch_that_assign_as_send_value := make(chan <- Type, buffer_size) // upstream
	// only_ch_that_assign_as_receive_value := make(<- chan Type, buffer_size) // downstream
	// this both will be usefull if you wanna use them as param in some func(upstream, downstream)

	// station with 2 slot of truck
	// you can call them as chan/bucket
	ch := make(chan []int, 2)
	fmt.Println(reflect.ValueOf(ch).Interface())

	ch <- []int{1} // truck take order and give id `1` as value
	// here for any syntax  // put any completeness between truck1 and truck2
	fmt.Println(<-ch) // pull order off. if truck `1`, been completed

	ch <- []int{2} // truck take order and give id `2` as value
	// here for any syntax  // put any completeness between truck2 and end of `station stream schema`
	fmt.Println(<-ch) // pull order off. if truck `2`, been completed

	// ============================[ new ronde ]===============================
	// now station routine was empty, any other truck can ask 1..2 order to `the station`
	ch <- []int{3}   // truck take order and give id `3` as value
	ch <- []int{123} // truck take order and give id `4` as value

	// ch <- 5 // dont put right back, goroutines will be asleep/deadlock
	// dont panic!; code above hapend
	// because you have more or equal thread's waiting/idle,
	// than your active thread's must be running or completed.

	fmt.Println(<-ch) //just 1 left
	fmt.Println(<-ch) //just 0 left

}
