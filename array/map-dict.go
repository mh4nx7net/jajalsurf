package main
import (
	"fmt"
)

type testing struct{
	T struct{
	}
}


func TestSearch(t *testing.T){
	dictionary := map[string]string{"test":"ini test"}

	got := search(dictionary,"test")
	want := "ini test"

	if got != want{
		t.Errorf()
	}
}
func main() {
	

}
