package piscine

func ListClear(l *List) {
	l.Head, l.Tail = nil, nil
}
