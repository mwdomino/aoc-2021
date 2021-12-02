package main

import "fmt"

type movement struct {
	direction string
	units     int
}

type position struct {
	x, y int
}

func main() {
	p := position{}

	for _, v := range inputs() {
		switch v.direction {
		case "forward":
			p.x += v.units
		case "up":
			p.y -= v.units
		case "down":
			p.y += v.units
		}
	}

	fmt.Println(p.x * p.y)
}
