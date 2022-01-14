package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	n := root
	if n == nil {
		return nil
	}
	for n.Left != nil {
		n = n.Left
	}
	return n
}
