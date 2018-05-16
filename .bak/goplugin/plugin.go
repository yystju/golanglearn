package main

import (
	"fmt"
	"math"
)

func SayFromPlugin(message string) string {
	return fmt.Sprintf("DATA : %s.", message)
}

func CheckPrime(n int) bool {
	var r int = int(math.Sqrt(float64(n))) + 1

	for i := 2; i < r; i++ {
		var s int = int(n/i) * i

		if s == n {
			return false
		}
	}

	return true
}
