package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

//type IIntervals []*Interval

//	func (s intervals) Len() int {
//		return len(s)
//	}
//
//	func (s intervals) Less(i, j int) bool {
//		return s[i].Start < s[j].Start
//	}
//
//	func (s intervals) Swap(i, j int) {
//		s[i].Start, s[j].Start = s[j].Start, s[i].Start
//		s[i].End, s[j].End = s[j].End, s[i].End
//	}
func main() {
	str := ""
	fmt.Scan(&str)
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	nstr := strings.Split(str, ",")
	var nums []*Interval
	for i := 0; i < len(nstr)-1; i += 2 {
		var tmp Interval
		tmp.Start, _ = strconv.Atoi(nstr[i])
		tmp.End, _ = strconv.Atoi(nstr[i+1])
		nums = append(nums, &tmp)
	}
	fmt.Println(nums)
}
func merge(intervals []*Interval) []*Interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var sum []*Interval
	for i := 0; i < len(intervals); i++ {
		if len(sum) == 0 {
			var tmp1 Interval
			tmp1.End = intervals[i].End
			tmp1.Start = intervals[i].Start
			sum = append(sum, &tmp1)
		} else if sum[len(sum)-1].End >= intervals[i].Start {
			if sum[len(sum)-1].End < intervals[i].End {
				sum[len(sum)-1].End = intervals[i].End
			}
		} else {
			var tmp1 Interval
			tmp1.End = intervals[i].End
			tmp1.Start = intervals[i].Start
			sum = append(sum, &tmp1)
		}
	}
	return sum
}
