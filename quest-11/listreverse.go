package piscine

func ListReverse(l *List) {
	if l.Head == l.Tail {
		return
	}
	tmpTail := l.Tail
	tmpHead := l.Head
	for tmpTail != tmpHead {
		for tmpHead.Next != tmpTail {
			tmpHead = tmpHead.Next
		}
		tmpTail.Next = tmpHead
		tmpHead.Next = nil
		tmpHead = l.Head
		tmpTail = tmpTail.Next
	}
	l.Head, l.Tail = l.Tail, l.Head
}
