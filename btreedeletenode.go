package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		var temp *TreeNode
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			temp = root
			root = root.Right
			temp = nil
		} else if root.Right == nil {
			temp = root
			root = root.Left
			temp = nil
		} else {
			temp = BTreeMin(root.Right)
			root.Data = temp.Data
			root.Right = BTreeDeleteNode(node.Right, node)
		}
	}
	return root
}
