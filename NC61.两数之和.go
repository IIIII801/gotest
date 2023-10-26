package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	// write code here
	target += 22
	m1 := [99999999]int{} //记录numbers的原始下标+1
	for i, x := range numbers {
		y := target - (x + 11)
		if y >= 0 && m1[y] != 0 {
			if m1[y] == -100 {
				m1[y] = 0
			}
			fmt.Println([]int{m1[y] + 1, i + 1})
			return []int{m1[y] + 1, i + 1}
		}
		if i == 0 {
			m1[x+11] = -100
		} else {
			m1[x+11] = int(i)
		}
	}
	//m := make(map[int]int, len(numbers)) //记录numbers的原始下标+1
	//
	//for i, x := range numbers {
	//	y := target - x
	//	if j, ok := m[y]; ok {
	//		fmt.Println([]int{j + 1, i + 1})
	//		return []int{j + 1, i + 1}
	//	}
	//	m[x] = i
	//}
	//for a := 0; a < len(numbers); a++ {
	//	for b := a + 1; b < len(numbers); b++ {
	//		if numbers[a]+numbers[b] == target {
	//			ss := []int{a + 1, b + 1}
	//			return ss
	//			break
	//		}
	//	}
	//}
	return numbers
}
func main() {
	fmt.Println(twoSum([]int{-3, 2, 3}, 0))
}
