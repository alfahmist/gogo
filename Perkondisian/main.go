package main

import "fmt"

func main() {

	// if, else if, else
	var number int = 100

	if number > 100 {
		fmt.Println("Number is greater than 100")
	} else if number < 100 {
		fmt.Println("Number is less than 100")
	} else if number == 100 {
		fmt.Println("Number is equal to 100")
	}

	// temporary variable
	var point = 10000.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.2f%s perfect!\n", percent, "%s")
	} else if percent >= 70 {
		fmt.Printf("%.2f%s Good!\n", percent, "%s")
	} else {
		fmt.Printf("%.2f%s Not Bad!\n", percent, "%s")
	}

	// seleksi swich case

	var poinnto = 9

	switch poinnto {
	case 8:
		fmt.Printf("8\n")
	case 7:
		fmt.Printf("7\n")
	case 6:
		fmt.Printf("6\n")
	default:
		fmt.Printf("default\n")
	}

	// A.13.4. Pemanfaatan case Untuk Banyak Kondisi
}
