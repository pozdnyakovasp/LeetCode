package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	isOverTen := false
	stepper1 := l1
	stepper2 := l2

	for stepper1 != nil || stepper2 != nil {
		currenSum := 0
		if stepper1 != nil {
			currenSum = stepper1.Val
			stepper1 = stepper1.Next
		}
		if stepper2 != nil {
			currenSum = currenSum + stepper2.Val
			stepper2 = stepper2.Next
		}
		if isOverTen {
			currenSum = currenSum + 1
		}

		currentVal := currenSum % 10
		isOverTen = currenSum/10 != 0
		res.Val = currentVal

		if stepper1 != nil || stepper2 != nil {
			res.Next = &ListNode{Next: nil}
			res = res.Next
		}
	}
	if isOverTen {
		res.Next = &ListNode{Val: 1}
	}
	return head
}
