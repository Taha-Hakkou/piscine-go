package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	} else if pos == 0 {
		return l
	}
	return ListAt(l.Next, pos-1)
}
