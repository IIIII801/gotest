package main

import "fmt"

func main() {
	str := "123"
	fmt.Println(permutation1(str))
}

func permutation1(S string) []string {
	ans := make(map[string]int)
	byte := []byte(S)
	ll := len(S)
	var dfs func(i int)
	dfs = func(i int) {
		if i == ll-1 {
			ans[string(byte)] = 1
		} else {
			for i1 := i; i1 < ll; i1++ {
				byte[i1], byte[i] = byte[i], byte[i1]
				dfs(i + 1)
				byte[i1], byte[i] = byte[i], byte[i1]
			}
		}
	}
	dfs(0)
	reans := []string{}
	for v, _ := range ans {
		reans = append(reans, v)
	}
	return reans
}
