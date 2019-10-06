package main

import (
	"fmt"
	"strconv"
	"strings"
)

//AddressGenerator - генерира списък от адреси по зададен шаблон
// синтаксис на шаблона
// R8-[02..04]-[A..D][1..2]
// всичко извън [] е статичен стринг
// в [] се задава (диапазон) range.
// от..до (генерират се всички стойности "от" "до"
// Има два типа диапазони:
//     1. Числови: при тях "от" и "до" трябва да са числа
//        широчината коригира с водещи нули (т.е. ако искаме
//        00, 01 .... пишем "0x..0y", дължината на стринга определя
//        дължината на генерирания израь.
//     2. Буквени - те трябва да са във вида "A..X" (т.е. само с по един символ)
func AddressGenerator(template string) []string {
	elems := strings.Split(template, "[")
	allelems := make([]string, 0)
	for _, e := range elems {
		tmp := strings.Split(e, "]")
		allelems = append(allelems, tmp...)
	}

	generatorFunc := func() []string {
		return []string{""}
	}

	for _, e := range allelems {
		rangeEl := strings.Split(e, "..")
		if len(rangeEl) == 1 {
			generatorFunc = staticGenerator(generatorFunc, e)
		} else {
			generatorFunc = rangeGenerator(generatorFunc, rangeEl[0], rangeEl[1])
		}
	}

	return generatorFunc()
}

func staticGenerator(prev func() []string, s string) func() []string {
	return func() []string {
		res := make([]string, 0)
		for _, prevEl := range prev() {
			res = append(res, fmt.Sprintf("%s%s", prevEl, s))
		}
		return res
	}
}

func rangeGenerator(prev func() []string, from string, to string) func() []string {
	if from[0] >= '0' && from[0] <= '9' {
		return func() []string {
			fromN, _ := strconv.Atoi(from)
			toN, _ := strconv.Atoi(to)
			res := make([]string, 0)

			width := len(to)

			for _, prevEl := range prev() {
				for i := fromN; i <= toN; i++ {
					res = append(res, fmt.Sprintf("%s%*.*d", prevEl, width, width, i))
				}
			}

			return res
		}
	}

	return func() []string {
		res := make([]string, 0)
		for _, prevEl := range prev() {
			for i := from[0]; i <= to[0]; i++ {
				res = append(res, fmt.Sprintf("%s%c", prevEl, i))
			}
		}
		return res
	}

}
