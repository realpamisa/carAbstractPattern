package main

type carSetters interface {
	setBrand(brand string)
	setType(carType string)
	getBrand() string
	getType() string
}

type carInfo struct {
	brand   string
	carType string
}

func (c carInfo) setBrand(brand string) {
	c.brand = brand
}

func (c carInfo) setType(carType string) {
	c.carType = carType
}

func (c carInfo) getBrand() string {
	return c.brand
}

func (c carInfo) getType() string {
	return c.carType
}
