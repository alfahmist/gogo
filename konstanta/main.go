package main

import "fmt"

func main() {
	const firstName string = "Fahmi"
	const lastName string = "Anwar"

	// firstName = "Fahmi Anwar" // This will cause an error because constants cannot be reassigned
	fmt.Println("Hello, my name is", firstName, lastName)

	// deklarasi multi konstanta
	const (
		square          = "kotak"
		isToday bool    = true
		pi      float64 = 3.14
		b
	)

	fmt.Println("Bentuk:", square)
	fmt.Println("today:", isToday)
	fmt.Println("pi:", pi)
	fmt.Println("b:", b)

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	fmt.Println("satu:", satu)
	fmt.Println("dua:", dua)
	fmt.Println("three:", three)
	fmt.Println("four:", four)
}
