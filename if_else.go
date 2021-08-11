package main

import "fmt"

func main() {

	x := 1
	x = 2

	if x == 0 {
		fmt.Println("Nilai x adalah ganjil, x  = ", x)
	} else if x == 2 {
		fmt.Println("Nilai x adalah genap, x = ", x)
	} else {
		fmt.Println("Nilai x tidak boleh kosong")
	}

	fmt.Println()

	var nilai uint = 8

	if nilai <= 6 {
		fmt.Println("Grade C")
	} else if nilai <= 8 {
		fmt.Println("Grade B")
	} else if nilai <= 10 {
		fmt.Println("Grade A")
	} else {
		fmt.Println("Nilai Kosong")
	}

	/*
		Variable Temporary
		Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan
		tanda ( := ) . Penggunaan keyword var disitu tidak diperbolehkan karena akan menyebabkan error
	*/
	var point = 10000.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	/*
		Switch

	*/
	var grade = "A"

	switch grade {
	case "A":
		fmt.Println("Perfect")
	case "B":
		fmt.Println("Very Good")
	case "C":
		fmt.Println("Good")
	case "D":
		fmt.Println("Not Bad")
	case "E":
		fmt.Println("Bad")
	case "F":
		fmt.Println("Very Bad")
	default:
		fmt.Println("Value Grade is null")
	}

	/*
		Switch Case Multiple Conditional
	*/
	var values = 6

	switch values {
	case 1, 2, 3:
		fmt.Println("Sangat Buruk")
	case 4, 5:
		fmt.Println("Buruk")
	case 6, 7:
		fmt.Println("Baik")
	case 8, 9:
		fmt.Println("Sangat Baik")
	case 10:
		fmt.Println("Perfect")
	default:
		fmt.Printf("Nilai %d\n", values)
	}

	/*
		Kurung Kurawal ( {} ) dapat digunakan pada statement case dan default,
		jika didalamnya terdapat banyak kodingan direkomendasikan untuk memakai {}
		agar kode lebih rapih
	*/

	var jenisKelamin string = "PL"
	switch jenisKelamin {
	case "P":
		fmt.Println("Jenis Kelaminnya Perempuan")
	case "L":
		fmt.Println("Jenis Kelamin Laki-Laki")
	default:
		{
			fmt.Println("Jenis kelaminnya tidak  jelas")
			fmt.Println("Entah Lelaki atau Perempuan")
		}
	}

	/*
		switch case yg menggunakan conditional
		Keyword fallthrough digunakan untuk memaksa proses pengecekkan diteruskan ke case selanjutnya.
	*/
	switch {
	case values == 8:
		{
			fmt.Printf("Nilai %d", values)
			fmt.Println()
		}
	case point >= 10000.0:
		fmt.Printf("%.2f\n", point)
		fallthrough // keyword ini digunakan untuk memaksa pengecekan kondisi dilanjutkan
	default:
		fmt.Println("Default")
	}

	/*
		Kondisi bersarang
	*/
	var nilais = 6
	if nilais > 7 {
		switch nilais {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if nilais == 5 {
			fmt.Println("not bad")
		} else if nilais == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if nilais == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
