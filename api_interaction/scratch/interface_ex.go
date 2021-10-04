package main

import "fmt"

type Abser interface {
	OBS() float64
}

func main() {
	var a Abser

	f := MyFruit(-10.20)

	a = f

	fmt.Println(a.OBS())

}

type MyFruit float64

func (f MyFruit) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
