package main

import (
	"fmt"
	"testing"
)

func TestAddressGen(t *testing.T) {
	template := "S4-[1..31]-[A..E][1..5]" //"R8-[02..04]-[A..D][1..2]"

	for _, g := range AddressGenerator(template) {
		fmt.Println(g)
	}
}

type sol struct {
	s int
	b int
}

var memory map[sol]sol

const (
	small = 11.14
	big   = 16.38
)

var callCnt int64
var cacheHit int64

func TestGetTotes(t *testing.T) {
	goal := 101.0
	sol := getTotes(goal)

	fmt.Printf("cachehit %d, %d sol = %v, goal = %f, vol = %f\n", cacheHit, callCnt, sol, goal, calc(sol))
}

func getTotes(goal float64) sol {
	memory = make(map[sol]sol)
	sol := findSol(goal, sol{s: 0, b: 0}, 0)
	return sol
}

func findSol(goal float64, s sol, solVal float64) sol {
	callCnt++
	if solVal >= goal {
		return s
	}

	var sol1 sol
	var sol2 sol

	v1 := sol{s: s.s + 1, b: s.b}
	v2 := sol{s: s.s, b: s.b + 1}

	if cached, ok := memory[v1]; ok {
		sol1 = cached
		cacheHit++
	} else {
		sol1 = findSol(goal, sol{s: s.s + 1, b: s.b}, solVal+small)
		memory[v1] = sol1
	}

	if cached, ok := memory[v2]; ok {
		sol2 = cached
		cacheHit++
	} else {
		sol2 = findSol(goal, v2, solVal+big)
		memory[v2] = sol2
	}

	sol1Val := calc(sol1)
	sol2Val := calc(sol2)

	if sol1Val <= sol2Val {
		return sol1
	}

	return sol2
}

func calc(sol sol) float64 {
	return float64(sol.s)*small + float64(sol.b)*big
}
