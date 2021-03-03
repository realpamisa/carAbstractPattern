package main

type ford struct{}

type fordInfo struct {
	carInfo
}

func (f *ford) makeCar() carSetters {
	return &fordInfo{
		carInfo: carInfo{
			brand:   "ranger",
			carType: "pickup",
		},
	}
}
