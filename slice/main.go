package main

import "fmt"

func main() {
	/*
		Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan
		dari manipulasi sebuah array ataupun slice lainnya. Karena slice merupakan data
		reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada
		slice lain yang memiliki alamat memori yang sama.

	*/
	// Membuat slice sama seprti membuat array, bedanya tidak perlu mendefinisikan panjang elemen array

	var fruits = []string{"Apple", "Orange", "Watermelon", "Grape"} // Jadi alih alih menggunakan [...], gunakan [] untuk membuat slice, jika membuat menggunakan [...] maka kita membuat arrray

	fmt.Println(fruits[0])

	// Kalau cuman beda saat pembuatan alokasi panjang array kenapa harus di bedakan? Sebenarnya tidak bisa dibedakan karena mereka sebuah kesatuan Array kumpulan nilai atau elemen sedangkan slide adalah referensi tiap elemen tersebut, jadi kalau elemen dalam array berubah maka nilai slice pun berubah, nah jadi slice ini bisa di bentuk dari array yang sudah di definisikan, dengan menggunakan 2 index untuk mengambil elemen-nya

	// Contoh slice dengan menggunakan 2 index untuk mengambil elemennya
	var newFruits = fruits[0:2] // Jadi akan mengambil dari index 0, hingga index sebelum 2. mirip range() di python, yang mana sifatnya eksklusif, index 2 itu tidak akan termasuk, jadi elemen nya di ambil kemudian di simpan dalam variabel baru

	fmt.Println(newFruits) // Apple Orange,

	// Jadi logic nya seperti ini, ketika kita buat variabel baru a = array[0], maka nilai yang di dapat itu adalah hasil COPY dari si array[0], beda dengan mengakses menggunakan 2 index tadi, nilai yang di dapat adalah reference elemen atau SLICE, jadi reference tuh kek alamat memori gitu lah

	// Berikut adalah macam macam operasi slice
	fmt.Println(fruits[0:2]) // [Apple Orange]
	fmt.Println(fruits[0:4]) // [Apple Orange Watermelon Grape]
	fmt.Println(fruits[0:0]) // [] Kosong karena tidak ada elemen sebelum index 0
	fmt.Println(fruits[4:4]) // [] Kosong karena tidak ada elemen index ke-4
	fmt.Println(fruits[2:])  // [Watermelon Grape], Semua elemen dimulai dari index ke-2
	fmt.Println(fruits[:3])  // [Apple Orange Watermelon], Semua elemen sebelum index ke-3

	/*
			Slice merupakan tipe data reference atau referensi. Artinya jika ada slice baru
		yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki
		alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang
		terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang
		memiliki referensi yang sama.

	*/
	print("===============\n")
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	fmt.Println(fruits)   // [apple orange watermelon grape]
	fmt.Println(aFruits)  // [apple orange watermelon]
	fmt.Println(bFruits)  // [orange watermelon grape]
	fmt.Println(aaFruits) // [orange]
	fmt.Println(baFruits) // [orange]
	// Buah "orange" diubah menjadi "pinnaple", maka semua juga ikut berubah, karena yang di ubah value yang ada di dalam reference nya atau di dalam memori nya
	baFruits[0] = "pinnaple"
	print("===============\n")
	fmt.Println(fruits)   // [apple orange watermelon grape]
	fmt.Println(aFruits)  // [apple orange watermelon]
	fmt.Println(bFruits)  // [orange watermelon grape]
	fmt.Println(aaFruits) // [orange]
	fmt.Println(baFruits) // [orange]
	print("===============\n")
	print("===============\n")

	fmt.Println(len(aFruits)) // Menghitung panjang slice atau array
	fmt.Println(cap(aFruits)) // Menghitung panjang slice berdasarkan array asli nya atau slice aslinya, jadi misal sebelum di slice panjang nya adalah 4, dan setelah di slice panjangnya 3 maka jika gunakan cap akan mengembalikan nilai 4 karena aslinya adalah 4

	fmt.Println(cap(bFruits)) // tapi ini tetap cap nya 3, kenapa? karena index start di mulai dari angka 1, yakni lebih dari 0, nah ketika di slice maka nilai index 1, akan menjadi index 0 di slice nya sehingga tetap akan terhitung 3
	// contoh visual
	/*
		afruits = [0,1,2, __] <-- ini dia akan tetap starting index nya di 0, sehingga ya tetap akan hitung cap nya dari index 0 itu, sehingga index 3 itu akan tetap masuk juga karena kita tidak potong dari startin point, beda seperti yang di bawah
		bfruits = __ [1,2,3] <-- disini motong nya pada saat starting point, jadi nya yang kepotong itu tidak akan masuk hitungan ketika menggunakan cap
	*/

	// Append, menambahkan item baru ke dalam array atau slice, kalau kita append item baru ke dalam slice, jika masih ada ruang dalam slice nya, dalama artian len nya < cap nya, maka dia akan mengambil referensi yang sama, jika tidak maka dia akan buat referensi yang baru. jadi jika kita tambah lewat slice yang len nya < cap nya, maka array asli nya akan berubah, karena sharing referensi yang sama
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits)  // [apple orange watermelon grape]
	fmt.Println(cFruits) // [apple orange watermelon grape, papaya]

	// Contoh mengenai len < cap
	/*
		fruits = [Apple, Banana, Orange]
		afruits = fruits[0:2] <-- berarti ini [Apple, Banana] dong
		nah jika kita append ke afruits, maka yang terjadi adalah kita akan tambah ke referensi yang ada di index[2] yakni si orange

		cfruits = append(afruits, "Papaya") <-- nah ini akan ganti value dari referensi di index[2]

		visualisasi
		afruits = [Apple, Banana, __] <-- tanda __ adalah referensi yang telah ada tapi tidak di ambil oleh afruits, ketika kita tambahkan item baru ke dalam afruits, maka value akan masuk ke referensi itu, makanya value orange di index[2] akan berubah, jadi akan terdeteksi sebagai perubahan nilai pada referensi yang lama


		cfruits = append(afruits, "Papaya")
		print(fruits) <-- [Apple, Banana, Orange]
	*/

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	// Jadi copy ini dia akan mencopy isi dari src ke dst, jika panjang dst adalah 3 dan isi dari src adalah 4, maka hanya 3 item yang dari src yang di copy. Begitupula jika panjang dan isi dst adalah 5 dan isi dari src cuman 2, maka 2 item dst akan di ganti menjadi item src dan sisanya masih ada, dan return value dari copy ini adalah jumlah [int] yang tercopy dari src ke dst

	// Slice mneggunakan 3 index
	/*
			3 index adalah teknik slicing untuk pengaksesan elemen yang sekaligus
		menentukan kapasitasnya. Cara penggunaannya yaitu dengan menyisipkan
		angka kapasitas di belakang, seperti fruits[0:1:1] . Angka kapasitas yang
		diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.

		Jadi kita bisa langsung set cap nya berapa
	*/

	bfruits := fruits[0:2:2] // Ini berarti mulai dari index 0 hingga index 1, kemudian set cap nya menjadi 2. Jadi ini benar benar akan motong sesuai dengan yang kita mau, tidak ada lagi reference yang bisa di ubah dengan append, karena reference nya udah ga keikut karena kita udah cap slice nya

	fmt.Println(len(bfruits))
	fmt.Println(cap(bfruits))

}
