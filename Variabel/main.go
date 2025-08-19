package main

import "fmt"

func main() {
	// deklarasi variabel dengan tipe data
	var firstname string = "Fahmi"
	var lastname string = "Anwar"

	fmt.Println("Hello, my name is", firstname, lastname)
	fmt.Printf("Halo %s %s"+"!", firstname, lastname)

	// deklarasi variabel tanpa tipe data

	name := "John"
	fmt.Printf("Halo %s \n", name)

	name = "Elton"
	fmt.Printf("Halo %s \n", name)

	// Deklarasi multi variable
	var first, second, third string = "satu", "dua", "tiga"
	four, five, sing := "empat", "lima", "enam"
	fmt.Println(first, second, third, four, five, sing)

	// Variable Underscore

	_ = "Belajar Golang"
	name, _ = "Jono", "Wiku"

	// Deklarasi variablle menggunakan keyword new

	newName := new(string)
	fmt.Println("New Name:", newName)
	fmt.Println("New Name:", *newName)

	// Deklarasi variable menggunakan keyword make

	//channel
	//slice
	//map
}
