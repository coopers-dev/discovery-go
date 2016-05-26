package maps

import "fmt"

func ExampleMultiSet() {
	m := Multiset{}
	fmt.Println(m)
	fmt.Println(m.Count("test"))

	m.Insert("test")
	fmt.Println(m.String())

	fmt.Println(m.Count("test"))

	m.Erase("test")
	fmt.Println(m.String())

	// Output:
	// { }
	// 0
	// { test }
	// 1
	// { }
}
