package get_intersection_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func _getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	hash := make(map[*ListNode]bool)
	for current := headA; current != nil; current = current.Next {
		hash[current] = true
	}
	for current := headB; current != nil; current = current.Next {
		if hash[current] {
			return current
		}
	}
	return nil
}

func __getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for currentA, currentB := headA, headB; currentA != nil || currentB != nil; {
		if currentA == currentB {
			return currentA
		}
		if currentA != nil {
			currentA = currentA.Next
		} else {
			currentA = headB
		}
		if currentB != nil {
			currentB = currentB.Next
		} else {
			currentB = headA
		}
	}
	return nil
}
