package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	ans := []int{}
	tmpans := make(map[int][]int)
	var getint func(root1 *TreeNode, dp int)
	getint = func(root1 *TreeNode, dp int) {
		if root1 != nil {
			tmpans[dp] = append(tmpans[dp], root1.Val)
			if root1.Left != nil && root1.Right != nil {
				dp++
				getint(root1.Left, dp)
				getint(root1.Right, dp)
			} else if root1.Left != nil {
				dp++
				getint(root1.Left, dp)
			} else if root1.Right != nil {
				dp++
				getint(root1.Right, dp)
			}
		}
	}
	getint(root, 0)
	for i := 0; i < len(tmpans); i++ {
		for _, v := range tmpans[i] {
			ans = append(ans, v)
		}
	}
	return ans
}
