package maps

import (
	"strings"
)

type Multiset map[string]int

func (m Multiset) Insert(val string) {
	m[val]++
}

func (m Multiset) Erase(val string) {
	delete(m, val)
}

func (m Multiset) Count(val string) int {
	return m[val]
}

func (m Multiset) String() string {
	v := make([]string, len(m))

	for value, _ := range m {
		v = append(v, value)
	}

	return "{" + strings.Join(v, " ") + " }"
}