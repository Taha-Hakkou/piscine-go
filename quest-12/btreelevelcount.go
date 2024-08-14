package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftTreeLevelCount := BTreeLevelCount(root.Left)
	rightTreeLevelCount := BTreeLevelCount(root.Right)
	if leftTreeLevelCount < rightTreeLevelCount {
		return 1 + rightTreeLevelCount
	}
	return 1 + leftTreeLevelCount
}
