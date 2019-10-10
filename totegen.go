package main

//Solution - ...............
type Solution struct {
	smallTotes int
	bigTotes   int
}

// 	small = 11.14
// 	big   = 16.38

//GetTotes - ..........
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

func findSolution(solution Solution,
	targetVolume float64,
	smallToteVolume float64,
	bigToteVolume float64,
	memory map[Solution]Solution,
) Solution {

	if calcVolume(solution, smallToteVolume, bigToteVolume) >= targetVolume {
		return solution
	}

	var sol1 Solution
	var sol2 Solution

	v1 := Solution{smallTotes: solution.smallTotes + 1, bigTotes: solution.bigTotes}
	v2 := Solution{smallTotes: solution.smallTotes, bigTotes: solution.bigTotes + 1}

	if cached, ok := memory[v1]; ok {
		sol1 = cached
	} else {
		sol1 = findSolution(v1, targetVolume, smallToteVolume, bigToteVolume, memory)
		memory[v1] = sol1
	}

	if cached, ok := memory[v2]; ok {
		sol2 = cached
	} else {
		sol2 = findSolution(v2, targetVolume, smallToteVolume, bigToteVolume, memory)
		memory[v2] = sol2
	}

	sol1Val := calcVolume(sol1, smallToteVolume, bigToteVolume)
	sol2Val := calcVolume(sol2, smallToteVolume, bigToteVolume)

	if sol1Val <= sol2Val {
		return sol1
	}

	return sol2
}

func calcVolume(s Solution, smallToteVolume float64, bigToteVolume float64) float64 {
	return float64(s.smallTotes)*smallToteVolume + float64(s.bigTotes)*bigToteVolume
}
