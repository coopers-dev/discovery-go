package _4

import "fmt"

func ExampleNewEvaluator() {
	eval := NewEvaluator(map[string]BinOp {
		"+" : func(a, b int) int {
			return a + b
		},
		"-" : func(a, b int) int {
			return a - b
		},
	}, PrecMap{
		"+": NewStrSet("+", "-"),
		"-": NewStrSet("+", "-"),
	})

	fmt.Println(eval("1 +      2"))

//	Output:
//	3
}