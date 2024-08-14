package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	var tmpTail *NodeI
	for tmpTail != l.Next {
		if l.Data > l.Next.Data {
			tmp := l.Next
			l.Next = tmp.Next
			tmp.Next = l
			l = tmp
		}
		tmpHead := l
		for tmpHead.Next.Next != tmpTail {
			tmp := tmpHead.Next
			if tmp.Data > tmp.Next.Data {
				tmpHead.Next = tmp.Next
				tmp.Next = tmpHead.Next.Next
				tmpHead.Next.Next = tmp
			}
			tmpHead = tmpHead.Next
		}
		tmpTail = tmpHead.Next
	}
	return l
}
