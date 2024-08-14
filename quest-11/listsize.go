package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	return 1 + ListSize(&List{Head: l.Head.Next})
}
