package main

import "fmt"

func main() {
	hondaCar := carFactory("honda")
	toyotaCar := carFactory("toyota")
	bmwCar := carFactory("bmw")
	fordCar := carFactory("ford")

	hondaCarInfo := hondaCar.makeCar()
	toyotaCarInfo := toyotaCar.makeCar()
	bmwCarInfo := bmwCar.makeCar()
	fordCarInfo := fordCar.makeCar()

	printCarInfo(hondaCarInfo)
	printCarInfo(toyotaCarInfo)
	printCarInfo(bmwCarInfo)
	printCarInfo(fordCarInfo)

}

func printCarInfo(c carSetters) {
	fmt.Println("Brand: ", c.getBrand())
	fmt.Println("Type: ", c.getType())
	fmt.Println()
}
