package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for tmp := l.Head; tmp != nil; tmp = tmp.Next {
		if comp(tmp.Data, ref) {
			return &tmp.Data
		}
	}
	return nil
}
