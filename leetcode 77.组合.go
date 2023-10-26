package main

import "fmt"

func main() {
	i, j := 0, 0
	fmt.Scan(&i, &j)
	fmt.Println(combine(i, j))
}
func combine(n int, k int) [][]int {
	list := []int{}
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	ans := [][]int{}
	var df func(max, index int)
	df = func(max, index int) {
		if index == k || k == n {
			tmp := []int{}
			for i := 0; i < k; i++ {
				tmp = append(tmp, list[i])
			}
			ans = append(ans, tmp)
		} else {
			for i := max; i <= n-k+index; i++ {
				list[i], list[index] = list[index], list[i]
				df(i+1, index+1)
				list[i], list[index] = list[index], list[i]
			}
		}
	}
	df(0, 0)
	return ans
}
