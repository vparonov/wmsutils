package main

import (
	"fmt"
	"math"
	"testing"
)

func Test1(t *testing.T) {
	targetVolume := 108.025000
	//9.500000
	//14.750000
	sol1 := GetTotes(targetVolume, 9.5, 15.00)

	fmt.Printf("(1) solution = %v, target = %f, solution volume = %f\n", sol1,
		targetVolume,
		calcVolume(sol1, 9.5, 15.00))

	sol2 := GetTotes(targetVolume, 11.14, 16.38)
	fmt.Printf("(2) solution = %v, target = %f, solution volume = %f\n", sol2, targetVolume,
		calcVolume(sol1, 11.14, 16.38))
}

func Test2(t *testing.T) {
	targetVolume := 17.0

	small := 11.14
	big := 16.38

	cs := int(math.Ceil(targetVolume / small))
	cb := int(math.Ceil(targetVolume / big))
	minV := 10000000.0
	solS := 0
	solV := 0
	for s := 0; s <= cs; s++ {
		for b := 0; b <= cb; b++ {
			v := float64(s)*small + float64(b)*big
			if v >= targetVolume {
				if minV >= v {
					minV = v
					solS = s
					solV = b
				}
				break
			}
		}
	}

	fmt.Printf("targetVolume = %f solution = %f, s = %d b =%d\n", targetVolume, minV, solS, solV)

}
