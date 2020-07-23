package main
import (
	"fmt"
	"math/rand"
	"time"
)
type Record struct{
	ID int
	Name string
	Color string
}

func main() {
	/*
	var okk []Record
	coba1 := Record{
		ID: 1619707036,
		Name: "oke",
		Color: "ok",
	}
	coba2 := Record{
		ID: 1619707036,
		Name: "oke",
		Color: "ok",
	}
	okk = append(okk, coba1, coba2)
	fmt.Println(okk)
	*/

	// Let's keep things unknow
	random := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Printf("%d\n",(*random).Intn(10)) // output decimal
	// fmt.Printf("%s",rand.New(rand.NewSource(1619707036)))

	// Create a large slice pretending
	// we will retrieved data from databases ;)

	// empty list with 1000column
	// kind like datas := { []Record, []Record, ...}
	datas := make([]Record, 1000) 


	for index, _ := range datas{ // range till 1000column
		pick := (*random).Intn(10)
		color:= "Red"

		if pick == 2 {
			color = "Blue"
		}
		datas[index] = Record{
			ID: index,
			Name: fmt.Sprintf("Rec: %d", index),
			Color: color,
		}
		// fmt.Println(datas[index])
	}

	// below contain: var [ Record1{}, Record2{}, ... ]
	var red []Record // clean list for red Record
	var blue []Record // clean list for blue record

	// datas := { []Record, []Record, ...} till 1000
	for _, record := range datas {
		if record.Color == "Red" {
			red = append(red, record)
		} else {
			blue = append(blue, record)
		}
	}

	fmt.Printf("\ntotal Red->\t[%d]\ntotal Blue->\t[%d]\n\n", len(red), len(blue))

}
