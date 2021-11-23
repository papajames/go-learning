package main

import (
	"ch10/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}

}
