package main

import "fmt"

func main() {
	// tipe data  numerik non-desimal
	var u8 uint8 = 8
	var u16 uint8 = 16
	var u32 uint32 = 32
	var u64 uint64 = 64
	var ut uint = 0
	var byt byte = 1
	var i8 int8 = -2
	var i16 int16 = -16
	var i32 int32 = -32
	var i64 int64 = -64
	var it int = -1
	var ru rune = 3

	fmt.Println(u8, u16, u32, u64, ut, byt, i8, i16, i32, i64, it, ru)

	// tipe data numerik desimal
	var f32 float32 = 3.14
	var f64 float64 = 3.14

	fmt.Println(f32, f64)
	fmt.Printf("%.3f %.10f \n", f32, f64)

	// tipe data boolean
	var isTrue bool = true
	var isFalse bool = false

	fmt.Printf("Boolean: %t %t \n", isTrue, isFalse)

	// tipe data string
	var name string = "Fahmi Anwar"
	fmt.Printf("%s name", name)

	//nil & zero value
}
