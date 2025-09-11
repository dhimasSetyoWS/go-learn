package main

import "fmt"

func main() {
	var a bool = true                  // Boolean
	var b int = 5                      // Integer
	var c float32 = 3.14               // Floating Number 32bit
	var d float64 = 3.1458912          // Floating Number 64bit
	var e string = "Aku adalah string" // String, string bisa pakai backtick seperti javascript `Aku adalah string` enaknya pakai backtick ini format stringnya ya sesuai dengan apa yang kita ketik di dalam backtick itu, jadi kalau kita buat baris baru di dalam backtick, maka format stringnya juga akan buat baris baru jg
	var f uint8 = 255                  // uint8 0 - 255
	/*
		uint8 = 0 <-> 255
		uint8, uint16, uint32, uint64. ini semua sama dimulai dari range dari 0 <-> maksimal nilai yang dapat di tampung
		bisa pakai uint saja nanti compiler akan menyesuaikan dengan value nya

		ada byte = uint8.
		int8 = -128 <-> 127 ini total dapat menampung 255 angka, mirip seperti uint tapi seperti di bagi 2 dari angka 0, sehingga di dapat jg nilai negatif nya

		sama seperti uint, ada int8, int16, int32, dan int64 range dari bagi 2 nya, dan ambil sisi yang paling negatif sebagai starting point

		bisa pakai int saja nanti compiler yang akan menyesuakna dengan valuenya


	*/
	// Nil value, atau zero value, di bawah ini semau adalah nil value, atau zero value
	var woi string = ""
	var cek bool = false
	var nondesimal = 0
	var desimal = 0.0
	//
	fmt.Println("Boolean : ", a)
	fmt.Println("Integer : ", b)
	fmt.Println("Float : ", c)
	fmt.Println("Float : ", d)
	fmt.Println("String : ", e)
	fmt.Println("Integer uint 8 bit : ", f)
}
