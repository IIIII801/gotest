package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		str := scan.Text()
		if len(str) < 9 {
			fmt.Println("NG")
			continue
		}
		mp2 := make(map[string]int, len(str))
		tum1 := [4]int{0, 0, 0, 0}
		flag := 0
		for i, v := range str {
			if i+3 <= len(str) {
				if _, ok := mp2[str[i:i+3]]; !ok {
					mp2[str[i:i+3]] = 1
				} else {
					flag = 1
					break
				}
			}
			if v >= 'A' && v <= 'Z' {
				tum1[0] = 1
			} else if v >= 'a' && v <= 'z' {
				tum1[1] = 1
			} else if v >= '0' && v <= '9' {
				tum1[2] = 1
			} else {
				tum1[3] = 1
			}
		}
		if (tum1[0]+tum1[1]+tum1[2]+tum1[3]) < 3 || flag == 1 {
			fmt.Println("NG")
			continue
		}
		fmt.Println("OK")
	}

}
