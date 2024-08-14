package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		leftTreeRM := root.Left
		for leftTreeRM.Right != nil {
			leftTreeRM = leftTreeRM.Right
		}
		if leftTreeRM.Data >= root.Data {
			return false
		}
	}
	if root.Right != nil {
		rightTreeLM := root.Right
		for rightTreeLM.Left != nil {
			rightTreeLM = rightTreeLM.Left
		}
		if rightTreeLM.Data <= root.Data {
			return false
		}
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
