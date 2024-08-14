package piscine

func ListPushFront(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: l.Head}
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
}
