package main

import (
	"LeetCode/solutions"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	//var res = solutions.FindMedianSortedArrays([]int{1, 2}, []int{3, 4, 5})
	//var res = solutions.FindMedianSortedArrays([]int{0, 1, 4}, []int{1, 2, 3, 4})

	//var res = solutions.FindMedianSortedArrays([]int{1, 3}, []int{2})
	//var res = solutions.FindMedianSortedArrays([]int{0, 1, 4}, []int{1, 2, 3, 4})
	//var res = solutions.FindMedianSortedArrays([]int{0}, []int{1})
	//var res = solutions.FindMedianSortedArrays([]int{0}, []int{0})
	var res = solutions.FindMedianSortedArrays([]int{1, 3}, []int{2, 7})
	//var res = solutions.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	// 0 1 1 2 3 4
	// 0 1 1 =2= 3 4 4
	fmt.Println(res)
	//var ln = solutions.LengthOfLongestSubstring(" ")
	//list1 := &solutions.ListNode{Val: 1}
	//list1.Next = &solutions.ListNode{Val: 4}
	//list1.Next.Next = &solutions.ListNode{Val: 5}
	//
	//list2 := &solutions.ListNode{Val: 2}
	//list2.Next = &solutions.ListNode{Val: 7}
	//
	//list3 := solutions.AddTwoNumbers(list1, list2)
	//
	//for list3 != nil {
	//	fmt.Printf("%d\n", list3.Val)
	//	list3 = list3.Next
	//}
}
