package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head, l.Tail = node, node
	} else {
		l.Tail.Next = node
		l.Tail = l.Tail.Next
	}
}
