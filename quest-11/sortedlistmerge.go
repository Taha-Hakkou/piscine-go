package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	var result *NodeI
	var resultTail *NodeI
	if n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			result, resultTail = n1, n1
			n1 = n1.Next
		} else {
			result, resultTail = n2, n2
			n2 = n2.Next
		}
	}
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			resultTail.Next = n1
			n1 = n1.Next
			resultTail.Next.Next = nil
		} else {
			resultTail.Next = n2
			n2 = n2.Next
			resultTail.Next.Next = nil
		}
		resultTail = resultTail.Next
	}
	for n1 != nil {
		resultTail.Next = n1
		resultTail = resultTail.Next
		n1 = n1.Next
	}
	for n2 != nil {
		resultTail.Next = n2
		resultTail = resultTail.Next
		n2 = n2.Next
	}
	return result
}
