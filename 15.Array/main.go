package main

import "fmt"

func main() {

	// A.15.1. Pengisian Elemen Array yang Melebihi Alokasi Awal

	var names [4]string

	names[0] = "a"
	names[1] = "b"
	names[2] = "c"
	names[3] = "d"

	fmt.Println(names[0], names[1], names[2], names[3])

	// A.15.2. Inisialisasi Nilai Awal Array

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah elemen \t\t", len(fruits))
	fmt.Println("Isi element \t", fruits)

}
