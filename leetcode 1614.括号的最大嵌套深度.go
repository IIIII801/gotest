package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	fmt.Println(maxDepth(s))

}
func maxDepth(s string) int {
	max := 0
	list := []rune{}
	for _, v := range s {
		if v == '(' {
			list = append(list, v)
			if len(list) > max {
				max = len(list)
			}
		}
		if v == ')' {
			list = list[:len(list)-1]
		}
	}
	return max
}
