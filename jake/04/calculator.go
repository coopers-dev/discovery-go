package _4

import (
	"strconv"
	"strings"
)

type BinOp func(int, int) int

type StrSet map[string]struct{}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}

	for _, str := range strs {
		m[str] = struct{}{}
	}

	return m
}

type PrecMap map[string]StrSet

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	ops := []string{"("}
	var nums []int

	pop := func() int {
		last := nums[len(nums) - 1]
		nums = nums[:len(nums) - 1]
		return last
	}

	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops) - 1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				return
			}
			ops = ops[:len(ops) - 1]
			if op == "(" {
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}

	for _, token := range strings.Fields(expr) {
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce(")")

	return nums[0]
}