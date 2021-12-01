package main

import "fmt"

type window [3]int

func main() {
	in := inputs()
	length := len(in)
	last := 0
	tally := 0

	for i := 1; i < length; i++ {
		var w window
		if i+2 < length {
			w[0] = in[i]
			w[1] = in[i+1]
			w[2] = in[i+2]
		} else {
			break
		}

		if w.sum() > last {
			tally++
		}

		last = w.sum()
	}

	fmt.Println(tally)
}

func (w *window) sum() int {
	sum := 0
	for _, v := range w {
		sum += v
	}

	return sum
}
