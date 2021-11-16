package main

import (
	"fmt"
	"reflect"
)

func main() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)
	fmt.Println(reflect.TypeOf(&amount))
}
