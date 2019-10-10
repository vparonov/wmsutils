package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	targetVolume := 15.44
	sol := GetTotes(targetVolume, 11.14, 16.38)

	fmt.Printf("solution = %v, target = %f, solution volume = %f\n", sol, targetVolume, calcVolume(sol, 11.14, 16.38))
}
