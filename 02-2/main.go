package main

import "fmt"

const (
	forward string = "forward"
	up      string = "up"
	down    string = "down"
)

// movement struct stores the details of an intent to move
// the submarine
type movement struct {
	direction string
	units     int
}

// position struct stores the horizontal, vertical, and aim
// positions of our submarine
type position struct {
	x, y, aim int
}

// MoveForward executes a forward move n spaces
func (p *position) MoveForward(n int) {
	p.x += n
	p.y += p.aim * n

}

// MoveUp executes a move upwards n spaces
func (p *position) MoveUp(n int) {
	p.aim -= n

}

// MoveDown executes a move downwards n spaces
func (p *position) MoveDown(n int) {
	p.aim += n
}

func main() {
	p := position{}

	for _, v := range inputs() {
		switch v.direction {
		case forward:
			p.MoveForward(v.units)
		case up:
			p.MoveUp(v.units)
		case down:
			p.MoveDown(v.units)
		}
	}

	fmt.Println(p.x * p.y)
}
