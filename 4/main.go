package main

import (
	"fmt"
	"math"
)

func main() {
	first()
}

func first() {
	ok := 0

	for i := 372304; i <= 847060; i++ {
		dd := false
		dc := 1

		di := true
		pd := 0

		for j := 6; j >= 1; j-- {
			nd := digit(i, j)

			if nd == pd {
				dc++
			} else {
				if dc == 2 {
					dd = true
				}
				dc = 1
			}

			if nd < pd {
				di = false
			}

			pd = nd
		}
		if dc == 2 {
			dd = true
		}

		if dd && di {
			ok++
		}
	}

	fmt.Println(ok)
}

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}
