package main

import "fmt"

type Car struct {
	BrandName string
	Model     string
	FuelType  string
}

func NewCar(brandName string, model string, fuel string) *Car {
	car1 := &Car{
		BrandName: brandName,
		Model:     model,
		FuelType:  fuel,
	}

	return car1
}

func (c *Car) Run(speed int) {
	fmt.Printf("%s is running with the speed of %d\n", c.Model, speed)
}

func (c *Car) Specification() {
	fmt.Printf("\nThe vehicle is: \n")
	fmt.Printf("Brand: %s\n", c.BrandName)
	fmt.Printf("Model: %s\n", c.Model)
	fmt.Printf("Fuel Type: %s\n", c.FuelType)
}

func main() {
	car := NewCar("Hundai", "Creta", "Petro")
	car.Specification()
	car.Run(20)

	car1 := NewCar("Honda", "Civic", "Petro")
	car1.Specification()
	car1.Run(20)

	car2 := NewCar("Honda", "Citi", "Petro")
	car2.Specification()
	car2.Run(20)

}
