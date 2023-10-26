package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type strss []byte

func (t strss) Len() int {
	return len(t)
}
func (t strss) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t strss) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {

	s := []string{"98", "988", "89999", "9121"}
	sort.Slice(s, func(i, j int) bool {
		for j1 := 0; true; j1++ {
			if s[i][j1:j1+1] != "" && s[j][j1:j1+1] != "" {
				return s[i][j1] > s[j][j1]
			} else if s[i][j1:j1+1] == "" {
				return true
			} else if s[j][j1:j1+1] == "" {
				return false
			}
		}
		return false
	})
	//vv := strings.Join(s, "")
	//aa := [][]int{}
	////bb := []int{}
	//ss := []int{}
	//ss = append(ss, 9)
	//aa = append(aa, ss)
	//dst2 := make([]int, len(ss))
	//aa[0][0] = 8
	//copy(dst2, ss)
	//dst2[0] = 999
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	cc := scan.Bytes()
	aa := scan.Text()
	//scan.Scan()
	//ss := scan.Text()
	fmt.Println(aa, cc)

}
func ll(a int) {
	println(a)
	if a < 3 {
		a++
		ll(a)
		ll(a)
	}
}
