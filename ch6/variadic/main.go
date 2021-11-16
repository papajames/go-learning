package main

import "fmt"

func main() {
	severalInts(1)
	severalInts(1, 2, 3)

	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d")
	severalStrings()
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}
