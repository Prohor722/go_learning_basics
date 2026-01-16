package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	W, H float64
}
