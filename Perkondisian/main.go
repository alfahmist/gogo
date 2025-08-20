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
	// Kurung Kurawal Pada Keyword case & default
	var pooi = 4

	switch pooi {
	case 1:
		fmt.Printf("1\n")
	case 2, 3, 4, 5, 6, 7:
		{

			fmt.Printf("2 3 5 6\n")
			fmt.Printf("2 3 5 6\n")
			fmt.Printf("2 3 5 6\n")
		}

	default:
		{

			fmt.Printf("default\n")
			fmt.Printf("default\n")
		}
	}

	switch {
	case (pooi < 1) && (pooi > 10):
		{

			fmt.Printf("2 3 5 6\n")
		}
	default:
		{

			fmt.Printf("default\n")
			fmt.Printf("default\n")
		}
	}

	// fallthrough

	var poiiiii = 6
	switch {
	case poiiiii == 8:
		fmt.Println("perfect")
	case (poiiiii < 8) && (poiiiii > 3):
		fmt.Println("awesome")
		fallthrough
	case poiiiii < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//A.13.8. Seleksi Kondisi Bersarang

	var iioopp = 0
	if iioopp > 7 {
		switch iioopp {
		case 10:
			fmt.Print("ADAWDAWD")
		default:
			fmt.Print("defaultt")
		}
	} else {
		if iioopp == 5 {
			fmt.Println("not bad")
		} else if iioopp == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if iioopp == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
