package main

import "fmt"

func main() {
	// Seperti operator artimatika pada umumnya, + - * / %
	var value = (((2*4)%3)*4 - 2) / 3
	fmt.Println(value)

	// Operator perbandingan
	var isEqual = (value == 2)
	fmt.Printf("nilai : %d, apakah sama ? (%t)\n", value, isEqual)
	// untuk format string nilai boolean menggunakan %t

	// Operator Logika
	// && || !
	var left = true
	var right = false

	var leftAndRight = left && right
	fmt.Printf("Kiri dan kanan = %t\n", leftAndRight)
	var leftOrRight = left || right
	fmt.Printf("Kiri atau kanan = %t\n", leftOrRight)
	var notLeft = !left
	fmt.Printf("Bukan Kiri/Reverse Kiri = %t\n", notLeft)
	var notRight = !right
	fmt.Printf("Bukan kanan/Reverse Kanan = %t\n", notRight)

}
