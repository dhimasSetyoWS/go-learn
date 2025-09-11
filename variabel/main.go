package main

import "fmt"

func main() {
	var firstName string = "John" // normal assign value
	var lastName string           // warning karena tidak di assign value tapi ini juga bisa
	lastName = "Doe"

	var fullName string = firstName + lastName // jika tidak digunakan maka akan error juga, solusi nya pakai _variabel

	_ = "Gw fullname" // nah ini bisa nih, tapi underscore variabel ga akan bisa ditampilkan soalnya dia kek tempat sampah aja, lu masukin ke situ ya ujung ujungnya bakal lenyap, nah untuk underscore value ini biasanya digunakan untuk menampung return value dari function yang tidak dipakai gitu

	name := new(string) // nah sekarang ada pointer nih kek di c++, jadi pointer ini bakal buat name ini as a memory gitu, jadi lu panggil si name ini, yang muncul tuh hanya alamat memory nya aja, kalau lu mau panggil value nya ya di dereference dlu caranya pakai (*)
	// contoh *name
	*name = "woilah" //assign value ke pointer
	fmt.Printf("Hello World and Welcome : %s %s %s \n", firstName, lastName, fullName)
	fmt.Println(name)
	fmt.Println(*name)
}
