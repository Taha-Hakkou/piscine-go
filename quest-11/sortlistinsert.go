package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}
	if l == nil {
		l = node
	} else if l.Data > data_ref {
		node.Next = l
		l = node
	} else {
		tmp := l
		for tmp.Next != nil && tmp.Next.Data < data_ref {
			tmp = tmp.Next
		}
		node.Next = tmp.Next
		tmp.Next = node
	}
	return l
}
