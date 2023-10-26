package main

func main() {

}
func findLengthOfLCIS(nums []int) int {
	max, tmp := 1, 1
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
