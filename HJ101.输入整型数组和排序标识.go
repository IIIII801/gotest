package main

import (
	"fmt"
	"sort"
)

func main() {
	numconut := 0
	fmt.Scan(&numconut)
	nums := []int{}
	for i := 0; i < numconut; i++ {
		num := 0
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	flag := 0
	fmt.Scan(&flag)
	sort.Ints(nums)
	if flag == 0 {
		for _, v := range nums {
			fmt.Print(v, " ")
		}
	} else {
		for i := numconut - 1; i >= 0; i-- {
			fmt.Print(nums[i], " ")
		}
	}

}
