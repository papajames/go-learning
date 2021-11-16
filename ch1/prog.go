package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(int64(123)))
}
