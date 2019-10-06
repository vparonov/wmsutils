package main

import (
	"fmt"
	"testing"
)

func TestAddressGen(t *testing.T) {
	template := "R8-[02..04]-[A..D][1..2]" //"S4-[1..31]-[A..E][1..5]"

	for _, g := range AddressGenerator(template) {
		fmt.Println(g)
	}
}
