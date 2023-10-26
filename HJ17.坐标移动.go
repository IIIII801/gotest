package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := ""
	fmt.Scanln(&str)
	strs := strings.Split(str, ";")
	x, y, step := 0, 0, 0
	for _, v := range strs {
		if len(v) < 2 {
			continue
		}
		if v, ok := strconv.Atoi(v[1:]); ok != nil {
			continue
		} else {
			step = v
		}
		switch v[0] {
		case 'W':
			y += step
		case 'A':
			x -= step
		case 'S':
			y -= step
		case 'D':
			x += step
		}
	}
	fmt.Printf("%d,%d", x, y)
}
