package main

import "fmt"
/*
type rectangle struct{
	length int
	breadth	int
	color string

	Geometry struct{
		area int
		perimeter int
	}
}
type rectElse struct {
	length int
	breadth int
	color string
}*/

type nestedStruct struct{
	length int
	color string
	nestedInside struct{
		fillA string
		fillB string
	}
}
type MoreDeep3 struct{
		putA string
		putB string
}

type MoreDeep struct{
	MoreDeep2 struct{
		putA string
		putB string
	}
	takeDeep []MoreDeep3
}
type Salary struct{ //but `Salary` not just an array, it'll become map[key]value
	Basic, HRA, TA float64 //define multiple datafield at once
}
type Employee struct{
	FirstName, Lastname, Email string //define multiple datafield at once
	Age int
	MonthlySalary []Salary // take from Salary and set as Array
	hexAddr []int
	moreNested []MoreDeep
}

func main() {
	a := &MoreDeep{
		takeDeep: []MoreDeep3{
			MoreDeep3{
				putA: "oke1",
				putB: "bisa1",
			},
		},
	}
	(*a).MoreDeep2.putA ="oke"
	(*a).MoreDeep2.putB ="bisa"
	fmt.Println(a.takeDeep[0])

	e := Employee{
		FirstName: "Herman",
		Lastname: "Abdurrahman",
		Email: "herman@gmail.com",
		Age: 20,
		hexAddr: []int{1,2,3},
		moreNested: []MoreDeep{
			MoreDeep{
				/*
				MoreDeep2{
					putA: "halo",
					putB: "dunia",
				},*/
				takeDeep: []MoreDeep3{
					MoreDeep3{
						putA: "oke",
						putB: "bisa",
					},
				},
			},
		},
		MonthlySalary: []Salary{
			Salary{
				Basic: 20.000,
				HRA: 15.000,
				TA: 29.000,
			},
			Salary{
				Basic: 22.000,
				HRA: 5.000,
				TA: 9.000,
			},

		},
	}
	changeNestedValue := &e.moreNested[0]
	(*changeNestedValue).MoreDeep2.putA = "haloo"
	(*changeNestedValue).MoreDeep2.putB = "duniaa"

	fmt.Println(e.moreNested[0])
	fmt.Println(e.FirstName,e.Lastname)
	fmt.Println(e.MonthlySalary[0].Basic)
	fmt.Println(e.hexAddr[1])
	fmt.Println(*changeNestedValue)
	fmt.Println(e)

	//===============================================================
	//===============================================================


	/*
	var rectA1 = &rectElse{1,2,"tiga"}
	fmt.Println("rect A1:\t",rectA1)

	var rectA2 = &rectElse{} // clean initiate
	rectA2.length = 10
	rectA2.color = "red"
	fmt.Println("rect A2:\t",rectA2)
	(*rectA2).breadth = 10 //edit value of refrence occured
	fmt.Println("rect A2:\t",rectA2)


	var rectA3 = &rectElse{} //as default actually it's clean struct!!
	(*rectA3).breadth = 100 //edit value of refrence occured
	(*rectA3).color = "green" //edit value of refrence occured
	(*rectA3).length = 99 //edit value of refrence occured
	fmt.Println("rect A3:\t", rectA3)

	var cekStruct = &nestedStruct{}
	(*cekStruct).length = 10
	(*cekStruct).color = "kelip - kelip"
	(*cekStruct).nestedInside.fillA = "oke"
	(*cekStruct).nestedInside.fillB = "oke"
	fmt.Println("cek Struct:\t",cekStruct)

	//===============================================================
	//===============================================================


	var rectA2 = &rectA1
	var rectA3 = &rectA2
	fmt.Printf("%v",rectA2) // 0xc00000e028
	fmt.Printf("%v",*rectA2) // &{1 2 tiga}
	fmt.Printf("%v",**rectA2) // {1 2 tiga}
	fmt.Printf("%v",***rectA3) // {1 2 tiga}

	//===============================================================
	//===============================================================

	var rect1 = rectElse{10, 20, "yellow"}
	fmt.Println(rect1)
	var rect2 = rectElse{length: 20, color: "black"}
	fmt.Println(rect2)

	var rect3 = new(rectElse) //pointer to instance
	rect3.length = 10
	rect3.breadth = 10
	rect3.color = "purple"
	fmt.Println(rect3)
	var rect4 = new(rectElse)
	rect4.length = 5
	rect4.breadth = 6
	rect4.color = "red"
	fmt.Printf("%v",rect4)

	var rectBlue rectangle //initialization purpose

	rectBlue.length = 10
	rectBlue.breadth = 20
	rectBlue.color = "blue"
	rectBlue.Geometry.area = rectBlue.length * rectBlue.breadth
	rectBlue.Geometry.perimeter = 2 * (rectBlue.length + rectBlue.breadth)

	//fmt.Println(rectBlue)
	//fmt.Println("area:\t", rectBlue.Geometry.area)


	var cleanStruct struct{ //initiate struct{}
		panjang int
		lebar int
		luas int
	}
	var cleanStruct1 struct{}

	var cleanMap map[string]string
	var cleanSlice []int
	var cleanArr [5]string

	cleanStruct.panjang= 10
	cleanStruct.lebar = 20
	cleanStruct.luas = cleanStruct.panjang * cleanStruct.lebar
	fmt.Println(cleanStruct.luas, cleanStruct1)

	fmt.Println(cleanMap, cleanSlice, cleanArr)

	*/
	//===============================================================
	//===============================================================

	/*
	oke := struct{
		"coba":10.5,
		"bisa":24.10,
		"tam":"red",
	}
	fmt.Println(oke)
	*/
}
