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

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		return
	}
	for l.Head.Next != nil && l.Head.Next.Data == data_ref {
		l.Head.Next = l.Head.Next.Next
	}
	ListRemoveIf(&List{Head: l.Head.Next}, data_ref)
}
