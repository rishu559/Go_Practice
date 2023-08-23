// Assignment 1

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type square struct {
// 	sideLength float64
// }

// type triagle struct {
// 	height float64
// 	base   float64
// }

// type shape interface{
// 	getArea() float64
// }

// func main() {
// 	sq := square{13.4}
// 	tr := triagle{11.2,2.4}
// 	printArea(sq)
// 	printArea(tr)
// }

// func printArea(s shape){
// 	fmt.Println(s.getArea())
// }

// func (t triagle) getArea() float64 {
// 	return t.base * t.height * 0.5
// }
// func (s square) getArea() float64 {
// 	return math.Pow(s.sideLength, 2)
// }

// Assignment 2

package main

import (
	"fmt"
	"io"
	"os"
)

type myReader struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing Parameter")
	}

	file, err := os.Open(os.Args[1])

	if err!=nil {
		fmt.Println("Cant read File:",os.Args[1])
		panic(err)
	}
	
	io.Copy(os.Stdout, file)
}

