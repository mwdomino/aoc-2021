package main

import (
	"fmt"
)

func main() {
	in := inputs()
	last := 0
	tally := 0

	for i, v := range in {
		// skip first measurement
		if i == 0 {
			last = v
		}

		// if measurement is > last measurement, increment
		if v > last {
			tally++
		}

		last = v
	}

	fmt.Println(tally)
}
