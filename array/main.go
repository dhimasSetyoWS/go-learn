package main

import "fmt"

func main() {

	var names [3]string // Panjang array 3, dan isinya adalah string
	names[0] = "Dhimas"
	names[1] = "Setyo"
	names[2] = "Wahyu"
	// names[3] = "Santoso" // Ini akan error karena melebihi jumlah alokas
	for index, item := range names {
		fmt.Printf("Index : [%d] - Isi : %s\n", index, item)
	}
	fmt.Println("===============")
	println()
	// Pengisian array juga bisa dari awal, caranya dengan menggunakan kurung kurawal, tapi gunakan =, jadi kayak set value untuk variabel tersebut

	var foods = [3]string{
		"Pizza",
		"Doughnut",
		"Pasta", // !Wajib ada juga koma di akhir kalau pakai cara vertikal!
	}

	// yang di atas ekuivalensi dengan horizontal seperti ini var foods = [3]string{"Pizza","Doughnut","Pasta"}
	fmt.Println("Jumlah makanan yang ada adalah :", len(foods))
	for index, item := range foods {
		fmt.Printf("Index : [%d] - Isi : %s\n", index, item)
	}

	fmt.Println("===============")
	println()

	// Bisa juga kita tidak set alokasi array nya berapa, jadi panjang array itu kita belum tahu atau kita belum set
	var numbers = [...]int{1, 2, 3}
	var angkadah = [10]string{"a", "b", "c"} // Kalau saya set 10 panjang array, jika itu int, maka isi dari sisa array itu berisi 0, begitupun jika boolean, maka isi nya adalah false, jika string maka string kosong, jadi akan di isi dengan nil value, atau zero value

	fmt.Println(angkadah)

	fmt.Println("Ini data array\t:", numbers)
	fmt.Println("Jumlah elemen\t:", len(numbers)) // Otomatis kapasitas nya akan ngikut

	// Array Multidimensi
	print("=============\n")
	dimensi1 := [2][3]int{[3]int{3, 2, 3}, [3]int{2, 1, 1}}
	// fmt.Println(dimensi1[0][1])
	// fmt.Println(dimensi1[1][2]) // Ya mirip mirip manggil matriks lah jadi array[baris][kolom] <-- ini simple nya

	// Pakai increment index
	for i := 0; i < len(dimensi1); i++ {
		for j := 0; j < len(dimensi1[i]); j++ {
			fmt.Print(dimensi1[i][j])
		}
		fmt.Println()
	}
	// Pakai range
	for index, item := range dimensi1 {
		fmt.Printf("Index : [%d] - Isi : %d\n", index, item)
	}

	// Jika kita tidak mau index nya gunakan _ sebagai variabel, jadi si underscore ini kek keranjang sampah lah, di tampung nilai nya terus di buang, jika sebaliknya ya di variabel pertama ganti index, variabel ke-2 pakai _
	for _, item := range dimensi1 {
		fmt.Printf("Isi : %d\n", item)
	}

	// Deklarasi ataupun alokasi kapasitas array menggunakan make()

	var gwarray = make([]string, 2) // Parameter pertama tipe datanya, param ke-2 panjang nya
	gwarray[0] = "Gigi"
	gwarray[1] = "Kuning"

	fmt.Println(gwarray)

}
