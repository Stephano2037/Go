package main

import "fmt"

//primitives
func main() {

	// c := 1 == 2
	// a := 10 //1010
	// b := 3  //0011

	// fmt.Printf("%v,%T\n", c, c)
	// fmt.Println(a & b) //0010
	// fmt.Println(a | b)
	// fmt.Println(a ^ b)
	// fmt.Println(a &^ b)

	// var n float64 = 3.14
	// n = 2.1e14
	// fmt.Printf("%v,%T\n", n, n)

	//complex , 복소수
	// var j complex64 = 1 + 2i
	// fmt.Printf("%v,%T\n", j, j)
	// fmt.Printf("%v,%T\n", real(j), real(j))
	// fmt.Printf("%v,%T\n", imag(j), imag(j))

	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	b := []byte(s)
	fmt.Printf("%v,%T\n", string(b), b)
	//can't do this
	//fmt.Printf("%v, %T\n", s, s)
}
