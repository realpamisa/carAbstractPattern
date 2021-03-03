package main

type carManufacturer interface {
	makeCar() carSetters
}

func carFactory(brand string) carManufacturer {
	if brand == "honda" {
		return &honda{}
	}
	if brand == "toyota" {
		return &toyota{}
	}
	if brand == "bmw" {
		return &bmw{}
	}
	if brand == "ford" {
		return &ford{}
	}

	return nil
}
