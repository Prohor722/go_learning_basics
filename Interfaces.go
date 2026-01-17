package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	W, H float64
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

