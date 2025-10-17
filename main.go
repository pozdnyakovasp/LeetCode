package main

import (
	"LeetCode/solutions"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	list1 := &solutions.ListNode{Val: 1}
	list1.Next = &solutions.ListNode{Val: 4}
	list1.Next.Next = &solutions.ListNode{Val: 5}

	list2 := &solutions.ListNode{Val: 2}
	list2.Next = &solutions.ListNode{Val: 7}

	list3 := solutions.AddTwoNumbers(list1, list2)

	for list3 != nil {
		fmt.Printf("%d\n", list3.Val)
		list3 = list3.Next
	}
}
