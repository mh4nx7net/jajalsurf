// interfaces: abstraksi tipe data yang telah dipesan/kontrak/teken/signatured
// tujuan: struct berbeda akan tetapi dapat menggunakan kontrak yang sama
// agar penggunaan ruang(rom/ram_ lebih efieien

package main

{
"fmt",
"math"
}
type geometri interface {
	// initialize type of identified function *later
	// kumpulan signature/kontrak/teken dari metode(yang telah dipesan)
	// collection of method signature

	luas() float64     // aria() sebagai method atau didalam interface disebut contract/fungsi yang diteken
	keliling() float64 // float64 sebagai return value atau tipe data yang berlaku
}

type persegi struct {
	width, height float64
}

type lingkaran struct {
	radius float64
}

func (c circle) luas() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) keliling() float64 {
	return 2 * math.Pi * c.
}

func measure(g geometry) {
	// melakukan beberapa return
	// atau job kalkulasi
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	// inisialisasi nilai
	r := rect{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}

	// fungsi yang dipanggil
	measure(r)
	measure(c)
}
