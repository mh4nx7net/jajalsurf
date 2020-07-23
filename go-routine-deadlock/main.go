package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func simple_deadlock() {
	c := make(chan bool)
	// c <- true // fatal error: all goroutines are asleep - deadlock! --> goroutine 1 [chan send]:
	<-c // fatal error: all goroutines are asleep - deadlock! --> goroutine 1 [chan receive]:

	// mengapa ini terjadi,
	// yap dikarenakan tidak ada sinyal yang dikirimkan(send)
	// tiba2 saja melakukan penerimaan(receive)

	// deadlock-chan receive; terjadi karena utas yang dibuat dan aktif lebih banyak
	// daripada yang semestinya mengantri
	// INTINYA --> tidak ada yang mengantri; sedangkan utas dipaksa dibuat dan dipaksa aktif

	// secara dasar, program membutuhkan 3 utas, namun;
	// pada waktu yang sama, ketika deadlock terdeteksi, maka 4 utas akan dibuat;
	// 1. untuk go-routine, berjalannya program inti
	// 2. untuk monitoring system; `sysmon`
	// 3. untuk garbage collector, yang dijalankan oleh go-routine program inti
	// 4. yang terakhir dibuat ketika go-routine program utama/inti diblockir selama inisialisasi
	// 		maka ketika ini terjadi, go membuat utas baru(time-up) untuk rutin yang lain

	// setiap utas terjadi idle, detektor akan dinotis. deadlock terjadi ketika nilai utas idle sama dengan
	// banyaknya utas yang aktif dikurangi dengan jumlah utas yang dimiliki sistem.

	// INTINYA --> keseluruhan proses membutuhkan 4 utas, dan sistem-sendiri menggunakan 4 utas
	// 4utas proses - 4 utas sistem = 0 utas(habis/null)
	// disisi lain, deadlock terjadi karena utas aktif berjumlah keseluruhan 3,
	// disaat yang sama ketiga utas itu mengalami idle, hingga tejadilah buntu

	// NB. Perlu diingat, perilaku ini tetap memiliki batasan, setiap utas akan diusahakan aktif
	// sehingga kerapkali hal ini membuat deteksi deadlock
	// tidak begitu berguna pada rutin yang tengah bekerja

}

func advance_deadlock() {
	// Praktik baru, dengan melakukan improvisasi dengan menanamkan `sinyal berhenti`
	// bila `sinyal interupsi` dikirimkan(send)

	s := make(chan os.Signal, 1) //same or equal as `c chan <- os.Signal`

	// semua akan berjalan baik
	// tapi ketika mendapat "syscall.SIGINT(for interupt)" maka akan dilakukan "syscall.SIGTERM(for terminate)"
	// `cause package signal to relay "incoming signal" to "chan of os.signal"`
	// pengalihan interaksi(relay) sinyal dari behaviour syscall.SIGINT kepada syscall.SIGTERM

	signal.Notify(s, syscall.SIGINT, syscall.SIGKILL)

	// inspeksi --> signal.Notify(s, syscall.SIGINT, `what_will_do`)
	// ./main.go:55:35: cannot use what_will_do (type what_will_do_Type) as type os.Signal in argument to signal.Notify:
	// int does not implement os.Signal (missing Signal method)

	// membuat channel baru
	c := make(chan bool)

	select {
	case <-c: // menjalankan channel c --> mengelola/menerima sinyal(receive)
	// ketika menjalankan channel c berlangsung, sebenarnya program mengalami deadlock

	case <-s: // menjalankan sinyal s --> mengelola/menerima sinyal(receive)
		// tidak jadi deadlock karena,
		// masih ada utas aktif dan berlangsung, yaitu rutin oleh `channel s`
		// rutin ini berjalan dengan menunggu pengiriman sinyal(send) oleh interaksi syscall.SIGINT
		println("program stopped")

		// ketika sinyal dari `channel s` terpenuhi (syscall.SIGINT)
		// println("program stopped") // atau turunan dari blok `case <-s:` akan dijalankan scr tuntas
	}
	println("keluar\n")
	// ketika interaksi sinyal.SIGINT dijalankan maka akan mendapat/menerima sinyal(receive) yang sebenarnya
	// hingga program tidak benar-benar mengalami deadlock
	// program diselesaikan hingga baris terakhir
	// dan diteruskan dengan melakukan perintah syscall.SIGKILL
	// perintah ini dijalankan oleh utas yang mengerjakan `signal.Notify(...` pada baris ke 56
	// yaitu utas `s :=...` sebagai pengelola kongkuren rutin teratas
}

func deadlock_with_debuging() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	time.Sleep(5 * time.Second)

	c := make(chan bool)
	<-c
}
func checkCh() {

	ch := make(chan int)

	go func() {
		select {
		case <-ch:
			log.Printf("1.channel")
		default:
			log.Printf("1.default")
		}
		select {
		case <-ch:
			log.Printf("2.channel")
		}
		close(ch)
		select {
		case <-ch:
			log.Printf("3.channel")
		default:
			log.Printf("3.default")
		}
	}()
	time.Sleep(time.Second)
	ch <- 1
	time.Sleep(time.Second)
}
func main() {
	// PENGENALAN / SIMPLE DEADLOCK
	// UNCOMMENT CODE BELLOW TO SEE ALL DESCRIPTION BEHAVIOUR
	// simple_deadlock()

	// DEADLOCK DENGAN PENJELASAN MENDALAM DAN PENANGANANNYA
	// UNCOMMENT CODE BELLOW TO SEE ALL DESCRIPTION BEHAVIOUR
	// advance_deadlock()

	// saatnya melakukan pengamatan lebih dalam !!
	// deadlock_with_debuging()
	// https://golang.org/pkg/net/http/pprof/
	// wget -O trace.out http://localhost:6060/debug/pprof/trace?seconds=5
	// go tool trace trace.out

	// checkCh()

}
