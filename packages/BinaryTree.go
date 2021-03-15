package packages

func NewBinaryTree(val interface{}) *BinaryTree {
	return &BinaryTree{
		Value: val,
	}
}

type BinaryTree struct {
	Value interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

func (obj *BinaryTree) AddLeft(val interface{}) {

	if obj.Left == nil {
		obj.Left = NewBinaryTree(val)
		return
	} else {
		obj.Left.AddLeft(val)
	}
}

func (obj *BinaryTree) AddRight(val interface{}) {

	if obj.Right == nil {
		obj.Right = NewBinaryTree(val)
		return
	} else {
		obj.Right.AddRight(val)
	}
}

func (obj *BinaryTree) Print(val interface{}) {

}

// func (obj *BinaryTree) ToArray() (result []interface{}) {

// 	left := obj.Left

// 	right := obj.Right

// 	emitVal := func(bin *BinaryTree, arr *[]interface{}) {

// 		if bin == nil {
// 			return
// 		}

// 		arr = append(arr)
// 	}

// 	for left != nil {

// 	}

// 	for right != nil {

// 		right = right.
// 	}

// 	return result
// }

func (obj *BinaryTree) LeftToArray() (result *[]interface{}) {

	return result
}
