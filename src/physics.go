package engine

import (
	"math/rand"
)

var X float64 = 0
var Y float64 = 0

//var right bool = true

func TestCode() {

	X = rand.Float64() * 300
	Y = rand.Float64() * 300

	/*
		if right && X <= 100 {
			X++
			Y += 1.25

			if X > 100 {
				right = !right
			}
		} else if !right && X >= 0 {
			X--
			Y -= 1.2

			if X < 0 {
				right = !right
			}
		}
	*/
}
