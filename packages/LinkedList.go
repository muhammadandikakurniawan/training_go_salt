package packages

import "fmt"

func NewLinkedList(value interface{}) *LinkedList {
	res := &LinkedList{
		Value: value,
	}
	return res
}

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
	Prev  *LinkedList
}

func (obj *LinkedList) InsertBegin(value interface{}) {

	curr := *obj //bikin memory baru

	*obj = *NewLinkedList(value)

	obj.Next = &curr
}

func (obj *LinkedList) InsertEnd(value interface{}) {
	obj.Tail().Next = NewLinkedList(value)
}

func (obj *LinkedList) DeleteTail() {
	curr := obj

	for curr.Next.Next != nil {
		curr = curr.Next
	}

	curr.Next = nil
}

func (obj *LinkedList) DeletePeak() {
	obj.Delete(1)
}

func (obj *LinkedList) Delete(i int) {
	curr := obj

	for i > 0 {
		if i == 1 && curr != nil {
			*curr = *curr.Next
			break
		}
		curr = curr.Next
		i--
	}
}

func (obj *LinkedList) Tail() *LinkedList {

	curr := obj

	for curr.Next != nil {
		curr = curr.Next
	}

	return curr
}

func (obj *LinkedList) Display() {

	curr := obj

	for curr != nil {
		fmt.Println(curr)
		curr = curr.Next
	}
}

func (obj *LinkedList) Search(value interface{}) *LinkedList {

	curr := obj

	for curr != nil {
		if curr.Value == value {
			return curr
		}
		curr = curr.Next
	}

	return nil
}
