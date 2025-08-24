package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)

	var newFruits = fruits[0:2]
	fmt.Println(newFruits)

	var fruitsa = fruits[:2]
	fmt.Println(fruitsa)

	fruitsa[0] = "LEMON"
	fmt.Println(fruits)

	var appFruits = append(fruits, "papaya")
	fmt.Println(appFruits)

}
