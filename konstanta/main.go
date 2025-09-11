package main

import "fmt"

func main() {
	const PInum float64 = 22 / 7
	fmt.Print("Ini adalah PI Number : ", PInum, '\n')

	// multi const

	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 91  // manifest typing, langsung set tipe datanya apaan
		floatNum       = 1.9 // type inference, compiler yang cari tau sendiri tipe datanya apaan
	)

	const (
		a = "aku adalah string"
		b // Ini akan ngikut value dari const a, jadi kalau tidak di set kayak gini maka dia akna ngikut value dari const yang ada di atasnya
	)

	// nih saya gabung antara keduanya ya
	const (
		sunday = "weekend"
		minggu
		isMinggu bool = true
	)

	// const satu baris
	const woi, gw, ganteng string = "yoi", "sip", "serah lu" // bisa pakai manifest bisa juga pakai inference
}
