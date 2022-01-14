package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	n := root
	if n == nil {
		return nil
	}
	for n.Right != nil {
		n = n.Right
	}
	return n
}
