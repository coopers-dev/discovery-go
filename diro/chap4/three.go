package chap4


import (
	"regexp"
	"strings"
	"strconv"
)

type BinOp func(int, int) int
type StrSet map[string]struct{}
type PrecMap map[string]StrSet

func EvalReplaceAll(in string) string {
	rx := regexp.MustCompile(`{[^}]+}`)
	return rx.ReplaceAllStringFunc(in, func(expr string) string {
		calexpr := strings.Trim(expr, "{ }")
		return strconv.Itoa(ExampleNewEvaluator(calexpr))
	})
}

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

func ExampleNewEvaluator(expr string) int {

	eval := NewEvaluator(map[string]BinOp {
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}

			if b < 0 {
				return 0
			}

			r := 1
			for i :=0; i<b; i++ {
				r *= a
			}
			return r
		},
		"*": func(a, b int) int { return a * b},
		"/": func(a, b int) int { return a / b},
		"mod": func(a, b int) int { return a % b},
		"+": func(a, b int) int { return a + b},
		"-": func(a, b int) int { return a - b},

	}, PrecMap {
		"**": NewStrSet(),
		"*": NewStrSet("**", "*", "/", "mod"),
		"/": NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+": NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-": NewStrSet("**", "*", "/", "mod", "+", "-"),
	})



	return eval(expr)
}

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	ops := []string{"("}
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]

			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				return
			}

			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}

			b, a := pop(), pop()

			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}

		}
	}

	for _, token := range strings.Split(expr, " ") {
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