package packages

func NewStack(val interface{}) *Stack {
	return &Stack{
		Val: val,
	}
}

type Stack struct {
	Val  interface{}
	next *Stack
}

func (obj *Stack) Push(val interface{}) {

	currNode := obj

	for currNode != nil {
		if currNode.next == nil {
			currNode.next = NewStack(val)
			break
		}
		currNode = currNode.next
	}

}

func (obj *Stack) Pop() (result Stack) {

	currNode := obj
	for currNode.next != nil {
		if currNode.next.next == nil {
			result = *currNode.next

			currNode.next = currNode.next.next
			break
		}
		currNode = currNode.next
	}

	return result
}
