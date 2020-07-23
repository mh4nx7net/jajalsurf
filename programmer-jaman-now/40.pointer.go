// secara dasar, seperti pemrograman umumnya
// golang menggunakan variable passing by value, not by refrence
// artinya ketika mendapatkan value dari variable pd fungsi.
// golang akan menduplikasikan, bukan menunjukkan map memori tertaut

package main

import "fmt"

type Address struct {
	Kota, Provinsi, Country string
}

func main() {
	addr1 := Address{
		"subang",
		"jabar",
		"indonesia",
	}

	// var addr2 *Address = &address1 //addr2 adalah pointer menuju map addr1
	addr2 := &addr1        //take a refrence and go to addr1 map memory
	addr2.Kota = "bandung" //addr2 => &addr1 >> mengganti apapun yg merujuk ke map addr1

	addr2 = &Address{"demak", "jateng", "indonesia"} //membuat map baru dengan struct-Address sbg ref-map

	addr3 := &addr1
	gantiKabeh := &addr1

	*gantiKabeh = Address{"jogja", "diy", "indonesia"}
	// var memAddr1 *Address =  addr1
	// fmt.Println(memAddr1)
	fmt.Println(addr1)
	fmt.Println(addr2)
	fmt.Println(addr3)
	fmt.Println(gantiKabeh)

	// pointer-new // pointer untuk map-memori yang baru
	// konventional: var addrX *Address = &Address{"dat1","dat2","dat3"}
	addrX := new(Address)
	addrX.Kota = "kudus"
	fmt.Println(addrX)
}
