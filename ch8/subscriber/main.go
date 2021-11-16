package main

import (
	"ch8/magazine"
	"fmt"
)

func printInfo(s *magazine.Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly rate:", s.Rate)
	fmt.Println("Active?", s.Active)
}

func defaultSubscriber(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true

	return &s
}

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 4.99
}

func main() {
	s1 := defaultSubscriber("Aman Singh")
	applyDiscount(s1)
	printInfo(s1)

	s2 := defaultSubscriber("Beth Ryan")
	printInfo(s2)
}
