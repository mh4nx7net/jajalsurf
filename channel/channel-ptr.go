package main
import "fmt"


// default of Creatures specification
type Dog struct{
	Name string
	Age int
}

// what DogCollection are ?
type DogCollection struct{
	Data []*Dog // that form by list of `struct Dog ptr`
}

// create Init() method
// that use *DogCollection as struct with parse their pointer
func (this *DogCollection) Init(){
	cloey := &Dog{"cloey", 1}
	ralph := &Dog{"ralph", 1}
	jakie := &Dog{"jakie", 1}
	sammy := &Dog{"sammy", 1}
	this.Data = []*Dog{cloey, ralph, jakie, sammy}
	// `this.Data` same as --> `*DogCollection.Data`
	// then `.Data` itself same as --> `[]*Dog`
	// or can be verb as: this.[]*Dog = []*Dog{}
	// remember brohhh.. it's functional language !!
	// any math law's here !!
}

// create CollectionChannel() method
// that use *DogCollection as struct with parse their pointer
func (this *DogCollection) CollectionChannel() chan *Dog{ //return channel buffer

	// create empty list (channel buffer)
	// based on *Dog (ptr)struct field data type's
	// that's called as chan
	dataChannel := make(chan *Dog, len(this.Data))
	fmt.Println(make( chan *Dog, len(this.Data))) // 0xc0000b4180
	fmt.Printf("%s\n", dataChannel) // 0xc0000b4120 or ptr of `channel buffer`

	// iterrate *DogCollection.Data as dog
	// and parsing the collection of list `[]*Dog`
	// to listed of dataChannel[]
	for _, dog := range this.Data{
		dataChannel <- dog // import dog to `channel buffer`
	}
	close(dataChannel)
	return dataChannel // dataChannel(%T: channel buffer) -> chan *Dog
}

func main() {
	/*
	ppp := make([]int, 10) // create new map literal
	ppo := new([10]int) // create new map literal; but by refrence
	fmt.Println(&ppp, ppo)
	*/
	dc := DogCollection{} //init struct to dc.units
	dc.Init() // struct of DogCollection that have method on it

	// fmt.Println(dc.CollectionChannel()) //0xc0000b4120
	for dog := range dc.CollectionChannel(){
		fmt.Printf("Channel Name: %s\n", dog.Name)
	}
}
