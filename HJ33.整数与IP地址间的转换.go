package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		if num, err := strconv.Atoi(str); err == nil {
			fmt.Println(numToIp(num))
		} else {
			fmt.Println(ipToNum(str))
		}
	}
}
func ipToNum(s string) int {
	strs := strings.Split(s, ".")
	sum := 0
	for _, v := range strs {
		if num, err := strconv.Atoi(v); err == nil {
			sum += sum*256 + num
		}
	}
	return sum
}
func numToIp(v int) string {
	strs := [4]string{}
	for i := 3; i >= 0; i-- {
		strs[i] = strconv.Itoa(v % int(math.Pow(2, 8)))
		v /= int(math.Pow(2, 8))
	}
	return strings.Join(strs[0:], ".")
}
