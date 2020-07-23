package main

import (
	"fmt"
)
type user struct{}
func main() {
	// declared nil slices string
	var data []struct{}
	var data1 []struct{}
	var data2 []struct{}

	datas := []string{}
	datas1 := []string{}
	fmt.Printf("%p\n%p\n%p\n%p\n%p\n",data,data1,data2,datas,datas1) //would redirect to the same ptr


	// var u user
	u:=user{}
	//data = []struct{"oke","okk"}
	fmt.Println(data, datas,u)

	var emptyStruct struct{}
	fmt.Println(emptyStruct)

}

