package main

type toyota struct{}

type toyotaInfo struct {
	carInfo
}

func (t *toyota) makeCar() carSetters {

	return &toyotaInfo{
		carInfo: carInfo{
			carType: "sedan",
			brand:   "toyota86",
		},
	}

}
