package main

//basic concurency example

import (
	"fmt"
	"reflect"
	"time"
)

// channel in go routine is used to veryfy and acomplis task of concurent
// channel buffer and unbuffered are...
// buffered -> adding continiously within/at certain point
// it's used inalot of computer's lives
// analogy -> ordinary buffer kind like truck vessle that wait till fill'sup
// after then, it's move around and would be changed with the another one's

// unbuffered -> just can handle only onething

// basic concept of channel is `send` and `receive`
// not read, write, though assign anything it's diffrent types/mechanism !!
// so, it diffrent construct than an simple variable

/*
1. create channel that saved to some variable
2. which variable itself would contain an address of channel would be
3. send some `value` to that variable/address if it been completed
*/

// jadi, channel disamping mengirimkan data(value based) dari satu gorutin ke lainnya(rutin)
// channel juga melakukan singkronasi channel

/*
Buffered channel
With a buffered channel, a Go routine may put a value in the channel ( ck <- 42 )
and then continue on with the next instruction without having to wait
for someone to read from the channel.
This is true unless the channel buffer is full.
If a channel is full, the Go routine will wait for another Go routine to read
from the channel before being able to put its own value there.

Unbuffered channel
An unbuffered channel will have no room to store any data.
So in order for a value to be passed over an unbuffered channel,
the sending Go routine will block until the receiving Go routine has received the value.
So, there is surely a difference between a buffered and an unbuffered channel

INTINYA --> BUFFERED CHANNEL:
selalu dapat menyimpan/mengirim sinyal/push --> (send), bersamaan dg ini
program akan selalu tereksekusi dg baik. namun ketika `bucket` penuh,
kita harus melakukan penerimaan/pengelolaan sinyal/pop --> (receive).
hal ini dilakukan agar bucket dapat digunakan kembali

INTINYA --> UNBUFFERED CHANNEL:
tidak seperti buffer channel, dimana memiliki bucket(room) dan terdeklarasi eksplisit
unbuffered tidak dapat menyimpan/menyangga data apapun,
jadi untuk setiap sinyal yang sukses dikirimkan/push --> (send),
pada saat itu juga rutin harus secara praktis menerima/mengelola sinyal/pop --> (receive)
agar eksekusi block program selanjutnya dapat diteruskan
*/

func concur1() {
	// done berfungsi untuk mengumpulkan sinyal channel (send)
	done := make(chan bool) // any that can receive status of channel what it must be
	fmt.Println(done)
	// `go` keyword 'ds used as concur routin initialization
	go func() {
		fmt.Println("Job 1 -> this from earth, named as Job 1")
		fmt.Println("Job 1 -> wait and will took about 10sec")
		time.Sleep(10 * time.Second)
		fmt.Println("Job 1 -> cheers, you can see me after 10sec")
		done <- true
	}() //private func with run as `go` concur routine

	go func() {
		fmt.Println()
		fmt.Println("Guest 1 -> heyy, here'sme from another universe")
		time.Sleep(5 * time.Second)
		fmt.Println("Guest 1 -> took several sec; then you can see me from another universe")
		done <- false
	}()

	fmt.Println("Job 2 -> initialization")
	// time.Sleep(10*time.Second) // give whatever sec for routine can complete their job's
	// just change with receive signal from `done` addr ptr
	// result berfungsi untuk memutakhirkan sinyal channel (receive)
	result := <-done

	fmt.Println(result)
}

func concur2() {
	/*
		ci := make(chan int)    // unbuffered channel of integers
		cj := make(chan int, 0) // unbuffered channel of integers
		ck := make(chan int, 1) // buffered channel of integers
	*/

	done := make(chan bool)
	go func() {
		fmt.Println()
		fmt.Println("Job1 -> Hello World, ready to wait")
		time.Sleep(2 * time.Second)
		fmt.Println("Job1 -> Wait success with 2sec after")
		done <- true //call me, if you wanna resume code below
		/*
			fmt.Println()
			fmt.Println("Job2 -> try to give another signal")
			fmt.Println("Job2 -> u'll se another signal after ")
			time.Sleep(2 * time.Second)
			fmt.Println("Job2 -> here another signal with 4sec after *literally")
			done <- true
		*/
	}()

	fmt.Println("Job9 -> iam from another universe")
	<-done
	fmt.Println()
	fmt.Println("clear and complete any jobs")
}

func tryReflecChan() {
	var val chan int // chan with based on int
	var right chan []int
	// right = append(right, 1)
	var strVal = reflect.ValueOf(&val) // return; reflect.value
	var strVal1 = reflect.ValueOf(right)
	fmt.Println(strVal1)
	value1 := reflect.MakeChan(reflect.Indirect(reflect.ValueOf(&val)).Type(), 0) //same return; reflect.value

	indirectStr := reflect.Indirect(strVal)
	value2 := reflect.MakeChan(indirectStr.Type(), 2048)

	fmt.Println(value1.Kind(), value1.Cap())
	fmt.Println(value2.Kind(), value2.Cap())
}

//  [ begin ]=================[ basic concurent: unbuffered ]==========================
var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world" // assign litle pieces of value
	// c <- 0             // receiver & sender bisa dibolak balik ? ;var c = make(chan int, 1) [works!!]
	x := <-c
	fmt.Println(reflect.ValueOf(x).Kind())
}
func main() {
	go f()
	print(a) // assign yet :/
	// <-c      //var c = make(chan int, 1) [works!!]
	// c <- 1   // receiver & sender bisa dibolak balik ?
	<-c
	print(a) // value's assign complete :D
	//  [ end ]==============[ basic concurent: unbuffered ]==========================

	// call `concur1` func if you wanna see them
	// concur2()
	// tryReflecChan()
	// bufferVSunbuffer()

}
