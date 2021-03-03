package main

type honda struct{}

type hondaInfo struct {
	carInfo
}

func (h *honda) makeCar() carSetters {
	return &hondaInfo{
		carInfo: carInfo{
			brand:   "city",
			carType: "sedan",
		},
	}
}
