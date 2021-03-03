package main

type bmw struct{}

type bmwInfo struct {
	carInfo
}

func (b *bmw) makeCar() carSetters {
	return &bmwInfo{
		carInfo: carInfo{
			brand:   "g150",
			carType: "sedan",
		},
	}
}
