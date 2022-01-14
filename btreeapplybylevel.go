package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 1; i <= BTreeLevelCount(root); i++ {
		Lev(root, i, f)
	}
}

func Lev(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else if i > 1 {
		Lev(root.Left, i-1, f)
		Lev(root.Right, i-1, f)

	}
}
