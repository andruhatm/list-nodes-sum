package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}

	var lknList1, lknList2 int

	lknList1 = parseLinkedListAsValue(l1)
	lknList2 = parseLinkedListAsValue(l2)

	//create node returning add result in reverse
	sum := lknList1 + lknList2

	fmt.Printf("lknList1: %d lknList2: %d with sum: %d countDigits(sum): %d \n", lknList1, lknList2, sum, countDigits(sum))

	itVal := sum
	for i := 1; i != countDigits(sum)+1; i++ {
		val := itVal % 10

		//TODO refactor
		if i == 1 {

		}

		res.addNode(val)
		fmt.Printf("before division itVal: %d\n", itVal)
		itVal = itVal / 10
		fmt.Printf("cur itVal: %d\n", itVal)
	}

	return res

}

func parseLinkedListAsValue(ll *ListNode) int {
	var (
		res int
	)

	iterator := 0
	curNode := ll

	fmt.Println("----------------")
	for {

		//TODO referse passed vals
		fmt.Printf("cur value %d \n", curNode.Val)
		val := curNode.Val * int(math.Pow(10, float64(iterator)))
		fmt.Printf("adding: %d \n", val)
		res += val
		iterator = iterator + 1

		if curNode.Next == nil {
			break
		}

		curNode = curNode.Next
	}
	fmt.Println("----------------")

	return res
}

func countDigits(num int) int {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	return count
}

func main() {

	tLL1 := &ListNode{Val: 1}
	tLL1.addRandNode()
	tLL1.addRandNode()
	tLL1.addRandNode()

	tLL2 := &ListNode{Val: 7}
	tLL2.addRandNode()
	tLL2.addRandNode()
	tLL2.addRandNode()

	fmt.Printf("tll1: %s tll2: %s \n", tLL1.toString(), tLL2.toString())

	fmt.Printf("res list: %s", addTwoNumbers(tLL1, tLL2).toString())

}

func (ll *ListNode) addRandNode() {
	curNode := ll

	for {
		if curNode.Next == nil {
			curNode.Next = &ListNode{Val: rand.Intn(10), Next: nil}
			break
		} else {
			curNode = curNode.Next
		}
	}
}

func (ll *ListNode) addNode(val int) {
	fmt.Printf("res: adding val - %d \n", val)
	curNode := ll
	for {
		if curNode.Next == nil {
			curNode.Next = &ListNode{Val: val, Next: nil}
			break
		} else {
			curNode = curNode.Next
		}
	}
}

func (ll *ListNode) toString() string {
	var res string
	curNode := ll

	for {

		res += strconv.Itoa(curNode.Val)

		if curNode.Next == nil {
			break
		}

		curNode = curNode.Next
	}

	return res
}
