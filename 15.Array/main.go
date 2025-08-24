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

	// cara horizontal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// cara vertikal
	// fruits  = [4]string{
	//     "apple",
	//     "grape",
	//     "banana",
	//     "melon",
	// }
	fmt.Println("Jumlah elemen \t\t", len(fruits))
	fmt.Println("Isi element \t", fruits)

	// tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Array multi dimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// A.15.6. Perulangan Elemen Array Menggunakan Keyword for

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range

	for i, fruits := range fruits {
		fmt.Printf("Element %d : %s\n", i, fruits)
	}

	// A.15.8. Penggunaan Variabel Underscore _ Dalam for - range

	for _, fruits := range fruits {
		fmt.Printf("fruit: %s\n", fruits)
	}

	// A.15.9. Alokasi Elemen Array Menggunakan Keyword make

	var fruitss = make([]string, 2)
	fruitss[0] = "apple"
	fruitss[1] = "manggo"

	fmt.Println(fruitss) // [apple manggo]
}
