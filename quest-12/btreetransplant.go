package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if node == nil {
		return root
	} else if node == root {
		return rplc
	}
	rplc.Parent = node.Parent
	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	return root
}
