package packages

import "fmt"

// func NewStack(val interface{}) *Stack {
// 	return &Stack{
// 		Val: val,
// 	}
// }

func NewStack() *Stack {
	return &Stack{}
}

type Stack []interface {
}

func (obj *Stack) IsEmpty() bool {
	return len(*obj) == 0
}

func (obj *Stack) Push(val interface{}) {
	*obj = append(*obj, val)
}

func (obj *Stack) Pop() interface{} {
	li := len(*obj) - 1

	res := (*obj)[li]

	*obj = (*obj)[:li]

	return res
}

func (obj *Stack) Print() {

	li := len(*obj) - 1

	for i := 0; i < len(*obj); i++ {
		v := (*obj)[li]
		fmt.Println(v)
		li--
	}

}

// type Stack struct {
// 	Val  interface{}
// 	next *Stack
// }

// func (obj *Stack) Push(val interface{}) {

// 	currNode := obj

// 	for currNode != nil {
// 		if currNode.next == nil {
// 			currNode.next = NewStack(val)
// 			break
// 		}
// 		currNode = currNode.next
// 	}

// }

// func (obj *Stack) Print() {

// 	currNode := obj

// 	for currNode != nil {
// 		fmt.Println(currNode.Val)
// 		currNode = currNode.next
// 	}

// }

// func (obj *Stack) Pop() (result Stack) {

// 	currNode := obj
// 	for currNode.next != nil {
// 		if currNode.next.next == nil {
// 			result = *currNode.next

// 			currNode.next = currNode.next.next
// 			break
// 		}
// 		currNode = currNode.next
// 	}

// 	return result
// }
