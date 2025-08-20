package main

import "fmt"

func main() {

	// basic for loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("angka %d \n", i)
	}

	// for dengan argumen hanya kondisi

	var d = 1
	for d <= 5 {
		fmt.Printf("Angka %d\n", d)
		d++
	}

	// for tanpa argumen
	var a = 1
	for {
		fmt.Printf("Angka %d\n", a)

		a++
		if a == 5 {
			break
		}
	}

	// for range

	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", string(v))
	}
	fmt.Println("=ARRAY=")

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var yb = [...]int{1, 2, 3, 4, 5} // array
	for _, x := range yb {
		fmt.Printf("Value=%d\n", x)
	}
	fmt.Println("SLICE")
	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	var lbd = map[byte]int{'a': 1, 'b': 2, 'c': 3}
	for l, b := range lbd {
		fmt.Println("ley=", string(l), "balue", b)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}
}
