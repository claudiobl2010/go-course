package main

type Calculator struct{}

func (*Calculator) Add(x, y int) int {
	return x + y
}

func (*Calculator) Multiply(x, y int) int {
	return x * y
}
