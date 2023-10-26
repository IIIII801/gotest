package main

import "fmt"

func main() {
	//aa := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	//aa := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	aa := [][]int{{1}, {2}}
	ss := orangesRotting(aa)
	fmt.Println(ss)
}
func orangesRotting(grid [][]int) int {
	var l func(g [][]int)
	var copyy func(g [][]int) [][]int
	copyy = func(g [][]int) [][]int {
		tmp1 := make([][]int, len(grid))
		for i := 0; i < len(grid); i++ {
			tmp1[i] = make([]int, len(grid[0]))
			copy(tmp1[i], grid[i])
		}
		return tmp1
	}
	var cop [][]int
	cop = copyy(grid)
	var isok func(g [][]int) [2]int
	isok = func(g [][]int) [2]int {
		fla := [2]int{1, 0}
	tabl:
		for i, v := range g {
			for j, v2 := range v {
				if cop[i][j] != v2 {
					fla[0] = 0
					break tabl
				}
				if v2 == 1 {
					fla[1] = 1
				}
			}
		}
		return fla
	}
	ans := 0
	l = func(g [][]int) {
		for i := 0; i < len(g); i++ {
			for j := 0; j < len(g[0]); j++ {
				if cop[i][j] == 2 && i+1 < len(g) && cop[i+1][j] == 1 {
					g[i+1][j] = 2
				}
				if cop[i][j] == 2 && i-1 >= 0 && cop[i-1][j] == 1 {
					g[i-1][j] = 2
				}
				if cop[i][j] == 2 && j+1 < len(g[i]) && cop[i][j+1] == 1 {
					g[i][j+1] = 2
				}
				if cop[i][j] == 2 && j-1 >= 0 && cop[i][j-1] == 1 {
					g[i][j-1] = 2
				}
			}
		}
		ss := isok(g)
		if ss[0] == 0 {
			ans += 1
			cop = copyy(g)
			l(g)
		} else if ss[1] == 1 {
			ans = -1
		}
	}
	l(grid)
	return ans
}
