package main

import (
	"fmt"
	"sort"
)

type tmp struct {
	key   int
	value int
}
type tmps []tmp

func (t tmps) Len() int {
	return len(t)
}

func (t tmps) Less(i, j int) bool {
	return t[i].key < t[j].key
}

func (t tmps) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	nub := 0
	fmt.Scan(&nub)
	m := make(map[int]int, nub)
	for i := 0; i < nub; i++ {
		index, value := 0, 0
		fmt.Scan(&index, &value)
		m[index] += value
	}
	var ts tmps
	for i, v := range m {
		ts = append(ts, tmp{key: i, value: v})
	}
	sort.Sort(ts)
	for _, v := range ts {
		fmt.Println(v.key, v.value)
	}
}
