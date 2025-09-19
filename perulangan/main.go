package main

import "fmt"

func main() {
	/*
		ada 3 perulangan di go seperti bahasa pemrograman lain:
		1. for
		2. foreach
		3. while
	*/

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	a := 0
	// Penggunaan for menggunakan kondisi saja
	for a < 5 {
		fmt.Println(a)
		a++
	}
	fmt.Println("==========")
	// Penggunaan for tanpa argumen
	for {
		fmt.Println(a)
		a++
		if a == 10 {
			break // hentikan perulangan
		}
	}
	fmt.Println("==========")
	// Penggunaan for dan range, digunakan untuk looping data array dll
	var xs = "ABC"
	// string jika digunakan range, maka yang akan di return adalah unicode nya, format menjadi %c agar menjadi karakter
	for index, item := range xs {
		fmt.Printf("Index : %d - Isinya : %c\n", index, item)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		// Jika isi dari kvs ada 3, maka perulangan akan jalan 3x
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i, "\n") // 01234
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue // Ini akan melanjutkan perulangan, jadi ketika bilangan ganjil, dia langsung lanjut ke perulangan selanjutnya
		}
		if i > 8 {
			// Jika i lebih dari 8, maka hentikan perulangan
			break
		}
		fmt.Println("Angka", i)
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	// disiapkan sebuah label bernama outerLoop untuk for di bawah nya
outerLoop: // Ini adalah label
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				// Ini akan memberhentikan loop terluar, jadi jika i == 3, maka dia akan hentikan looping yang ada di luar, jadi break ataupun continue tidak hanya berfungsi pada blok loop yang dia tempati saja, melainkan bisa ke loop yang ada di luar nya, dengan menggunakan label
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
