package main

import "fmt"

func main() {

	// go tidak support ternary operator
	var point int8 = 7

	if point == 10 {
		fmt.Println("Point anda sempurna")
	} else if cekgenap := point % 2; cekgenap == 0 { // cekgenap adalah variabel temporary, jadi hanya bisa di akses di dalam block else dan kebawah itu saja
		fmt.Println("Point yang anda punya adalah point genap :", point, cekgenap)
	} else {
		fmt.Println("Point anda ganjil :", point, cekgenap)
	}

	// switch case
	x := "b"
	switch x {
	case "a":
		fmt.Println("A Coy : Case ini adalah :", x)
		fallthrough // ini akan meneruskan case ke case selanjutnya, jadi di go tidak ada break; untuk switch case nya tapi kita bisa lanjutkan cek case dengan fallthrough
	case "b": // bisa menggunakan kurung kurawal seperti ini jika ada banyak statement di dalamnya, bisa digunakan di case ataupun default keyword
		{
			fmt.Println("B Coy : Case ini adalah :", x)
			fmt.Println("B Coy : Case ini adalah :", x)
			fmt.Println("B Coy : Case ini adalah :", x)
		}
	case "c", "d", "e", "f": // multi case
		fmt.Println("C Coy : Case ini adalah :", x)
	// case x == "z": // ini bertindak seperti if, uniknya golang bisa melakukan hal ini, ini bisa dilakukan jika switch tidak punya expression di atas, jadi cmn switch {case x == "a"} <-- ini bisa
	// 	fmt.Println("C Coy : Case ini adalah :", x)
	default:
		fmt.Println("Ga ada case ini mah")
	}

	nilai := 4

	switch {
	case nilai < 10:
		fmt.Println("Nilai kurang dari 10")
		fallthrough // lanjur ke case selanjutnya
	case nilai > 3:
		fmt.Println("Ngeri komiu, lebih dari 3")
	}

}
