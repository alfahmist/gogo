package main

import "fmt"

func main() {

	// This is the main entry point for the operator.
	// You can initialize your operator here, set up controllers, etc.

	// Operator Perbandingan

	var value = (((2+6)%3)*4 - 2) / 3
	var equalValue = 4
	var isEqual = (value == equalValue)

	fmt.Printf("Value: %d, Is Equal to %d: %t\n", value, equalValue, isEqual)

	//operator logika

	var left = false
	var right = true

	var leftAndRight = left && right
	var leftOrRight = left || right

	fmt.Printf("Left: %t,\nRight %t,\nLeft AND Right: %t,\nLeft OR Right: %t\n", left, right, leftAndRight, leftOrRight)

}
