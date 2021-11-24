package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) Break() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Break() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Break()
	Steer(string)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Break()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func main() {
	TryVehicle(Car("Toyoda Yarvic"))
	TryVehicle(Truck("Fnord F180"))
}
