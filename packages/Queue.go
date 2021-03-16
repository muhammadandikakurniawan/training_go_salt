package packages

import "fmt"

func NewQueue() *Queue {
	return &Queue{}
}

type Queue []interface {
}

func (obj *Queue) IsEmpty() bool {
	return len(*obj) == 0
}

func (obj *Queue) Push(val interface{}) {

	*obj = append(*obj, val)

}

func (obj *Queue) Dequeue() interface{} {

	if obj == nil {
		return nil
	}
	res := (*obj)[0]

	*obj = (*obj)[1:]

	return res
}

func (obj *Queue) Print() {
	for i := 0; i < len(*obj); i++ {
		fmt.Printf("%v <- ", (*obj)[i])
	}
}
