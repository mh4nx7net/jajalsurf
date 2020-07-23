package main

import (
	"fmt"
	"encoding/json"
)

type Employee struct{
	FirstName string `json: "firstname"`
	LastName string `json: "lastname"`
	City string `json: "city"`
}

type StructImplementTo struct{
	// key string `key/json: "value/key_that_fit"`
	A_Key1 string `json: "for key1"`
	// 1stChar must be Capitalize,
	// kind like : Key string `json: "description"`
	// then for the byte list, no mater the char is High/Lower Case
}

func main() {
	// marshaling is => conversion of higher level object into some kind of lower level object
	// or mean as converting an object of type %T into a lower level object.

	// implement with wrap of key-pair eval str ==> called as marshall
	// kind like eval on bash, that would parse to/as json
	// using json.Unmarshal([]list_ofByte(json_string), toPaired_memory_ofStruct)
	json_string := `{
		"firstname": "Hendra",
		"lastname": "Darmawan",
		"city": "Jakarta"
	}`

	User1 := new(Employee) //by refrence-and-not value
	json.Unmarshal([]byte(json_string), User1)
	jsonUser1, _ := json.Marshal(User1) //convert struct to json.Marshal
	fmt.Printf("%s\n", jsonUser1)

	// parse each byte of char to list of arr
	// by mean, convert unicode of char to embeded byte and put it on list of array
	// (`any thing that inside on backtick`)
	// []byte(`any thing inside of this back quote mark`)
	valueOfStructToImplement := []byte(`{
		"A_Key1": "ini value 1"
	}`)
	implement1 := new(StructImplementTo) //by refrence-and-not value
	json.Unmarshal(valueOfStructToImplement, implement1)
	fmt.Println(*implement1)
	implement1.A_Key1 = "bukan value 1 lagi"
	fmt.Printf("%s\n", *implement1)


	// implement without wrap of key-pair eval str
	User2 := new(Employee)
	User2.FirstName = "Andri"
	User2.LastName = "Sulistyo"
	User2.City = "Depok"
	jsonUser2, _ := json.Marshal(User2) //convert struct to json.Marshal
	fmt.Printf("%s\n", jsonUser2)
}
