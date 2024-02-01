package piscine

func BTreeLevelCount(root *TreeNode) int {
	for root != nil {
		l := BTreeLevelCount(root.Left)
		r := BTreeLevelCount(root.Right)
		if l < r {
			return r + 1
		} else {
			return l + 1
		}
	}
	return 0
}
