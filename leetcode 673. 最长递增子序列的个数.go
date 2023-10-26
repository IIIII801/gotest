package main

func main() {

}
func findLengthOfLCIS2(nums []int) int {
	max, max2, tmp, ssun := 1, 1, 1, 0
	alist := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			max2++
			if ssun > 0 {
				alist = append(alist, ssun)
				ssun = 0
			}
		} else {
			ssun++
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			tmp++
		} else {
			tmp = 1
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
