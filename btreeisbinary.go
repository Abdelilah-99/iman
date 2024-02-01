package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	lt := BTreeIsBinary(root.Left)
	rt := BTreeIsBinary(root.Right)
	if lt == true && rt == true {
		return true
	}
	return false
}
