package main

import "math"

//Solution - contains combination of small and big totes
type Solution struct {
	smallTotes int
	bigTotes   int
}

//GetTotes - returns the optimal combination of small and big totes for given targetVolume
func GetTotes(targetVolume float64, smallToteVolume float64, bigToteVolume float64) Solution {
	memory := make(map[Solution]Solution)
	return findSolution(
		Solution{smallTotes: 0, bigTotes: 0},
		targetVolume,
		smallToteVolume,
		bigToteVolume,
		memory,
	)
}

//GetTotes2 - returns the optimal combination of small and big totes for given targetVolume
// variant 2
func GetTotes2(targetVolume float64, smallToteVolume float64, bigToteVolume float64) Solution {
	cs := int(math.Ceil(targetVolume / smallToteVolume))
	cb := int(math.Ceil(targetVolume / bigToteVolume))
	sol := Solution{}
	minV := targetVolume + bigToteVolume

	for s := 0; s <= cs; s++ {
		for b := 0; b <= cb; b++ {
			v := float64(s)*smallToteVolume + float64(b)*bigToteVolume
			if v >= targetVolume {
				if minV >= v {
					minV = v
					sol.smallTotes = s
					sol.bigTotes = b
				}
				break
			}
		}
	}

	return sol
}

//
// finds the optimal solution
//
func findSolution(solution Solution,
	targetVolume float64,
	smallToteVolume float64,
	bigToteVolume float64,
	memory map[Solution]Solution,
) Solution {

	// if the volume is enough -  that is our solution
	if calcVolume(solution, smallToteVolume, bigToteVolume) >= targetVolume {
		return solution
	}

	var sol1 Solution
	var sol2 Solution

	// variant 1 => add one small tote to the previous solution
	v1 := Solution{smallTotes: solution.smallTotes + 1, bigTotes: solution.bigTotes}
	// variant 2 => add one big tote to the previous solution
	v2 := Solution{smallTotes: solution.smallTotes, bigTotes: solution.bigTotes + 1}

	// check if we have already calculated the solution for variant 1
	if cachedSolution, ok := memory[v1]; ok {
		sol1 = cachedSolution
	} else {
		sol1 = findSolution(v1, targetVolume, smallToteVolume, bigToteVolume, memory)
		memory[v1] = sol1
	}

	// the same for variant 2
	if cachedSolution, ok := memory[v2]; ok {
		sol2 = cachedSolution
	} else {
		sol2 = findSolution(v2, targetVolume, smallToteVolume, bigToteVolume, memory)
		memory[v2] = sol2
	}

	sol1Val := calcVolume(sol1, smallToteVolume, bigToteVolume)
	sol2Val := calcVolume(sol2, smallToteVolume, bigToteVolume)

	// the optimal solution is the one with minimum volume
	if sol1Val <= sol2Val {
		return sol1
	}

	return sol2
}

func calcVolume(s Solution, smallToteVolume float64, bigToteVolume float64) float64 {
	return float64(s.smallTotes)*smallToteVolume + float64(s.bigTotes)*bigToteVolume
}
