package main

import (
	"fmt"
	"math"
)

//Sqrt finds squareroot of a given number
func Sqrt(x float64) float64 {
	var z float64 = x / 2
	lastDifference := 0.0
	previousDoubleDifference := 0.0
	for {
		previousZ := z
		z -= (z*z - x) / (2 * z)
		if difference := z - previousZ; difference == lastDifference {
			return z
		} else if currentDoubleDifference := math.Abs(difference - lastDifference); currentDoubleDifference == previousDoubleDifference {
			if math.Min(difference, lastDifference) == difference {
				return z
			} else {
				return previousZ
			}
		} else {
			lastDifference = difference
			previousDoubleDifference = currentDoubleDifference
		}
	}
}

func main() {
	fmt.Println(Sqrt(971681347080147))
	fmt.Println("Math: ", math.Sqrt(971681347080147))
}
