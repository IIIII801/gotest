package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := ""
	fmt.Scan(&ss)
	fmt.Println(isValid(ss))
}
func isValid(s string) bool {
	// write code here
	s = strings.Trim(s, "\"")
	lis := []uint8{}
	flag := true
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			lis = append(lis, s[i])
		case '{':
			lis = append(lis, s[i])
		case '(':
			lis = append(lis, s[i])
		case ']':
			if len(lis) == 0 {
				flag = false
				break
			} else {
				if lis[len(lis)-1] == '[' {
					lis = lis[:len(lis)-1]
				} else {
					flag = false
				}
			}

		case '}':
			if len(lis) == 0 {
				flag = false
				break
			} else {
				if lis[len(lis)-1] == '{' {
					lis = lis[:len(lis)-1]
				} else {
					flag = false
				}
			}
		case ')':
			if len(lis) == 0 {
				flag = false
				break
			} else {
				if lis[len(lis)-1] == '(' {
					lis = lis[:len(lis)-1]
				} else {
					flag = false
				}
			}
		}
		if !flag {
			break
		}
	}
	if len(lis) > 0 {
		flag = false
	}
	return flag
}
