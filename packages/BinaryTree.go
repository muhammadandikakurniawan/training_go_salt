package packages

import "fmt"

type BinaryTree struct {
	root *BinaryTreeNode
}

type BinaryTreeNode struct {
	val   int64
	right *BinaryTreeNode
	left  *BinaryTreeNode
}

func (obj *BinaryTree) Insert(val int64) {
	if obj.root == nil {
		obj.root = &BinaryTreeNode{
			val: val,
		}
		return
	}

	obj.root.Insert(val)

}

func (obj *BinaryTree) Print() {

	if obj.root == nil {
		return
	}

	obj.root.left.Print()
	fmt.Println(obj.root.val)
	obj.root.right.Print()

}

func (obj *BinaryTreeNode) Print() {

	if obj == nil {
		return
	}

	fmt.Println(obj.left.val)
	fmt.Println(obj.val)
	fmt.Println(obj.right.val)

	obj.left.Print()
	obj.right.Print()
}

func (obj *BinaryTreeNode) Insert(val int64) {
	if obj == nil {
		return
	}

	if obj.val <= val {
		if obj.right == nil {
			obj.right = &BinaryTreeNode{
				val: val,
			}
		} else {
			obj.right.Insert(val)
		}
	} else {
		if obj.left == nil {
			obj.left = &BinaryTreeNode{
				val: val,
			}
		} else {
			obj.left.Insert(val)
		}
	}
}

// func NewBinaryTree(val interface{}) *BinaryTree {
// 	return &BinaryTree{
// 		Value: val,
// 	}
// }

// type BinaryTree struct {
// 	Value interface{}
// 	Left  *BinaryTree
// 	Right *BinaryTree
// }

// func (obj *BinaryTree) AddLeft(val interface{}) {

// 	if obj.Left == nil {
// 		obj.Left = NewBinaryTree(val)
// 		return
// 	} else {
// 		obj.Left.AddLeft(val)
// 	}
// }

// func (obj *BinaryTree) AddRight(val interface{}) {

// 	if obj.Right == nil {
// 		obj.Right = NewBinaryTree(val)
// 		return
// 	} else {
// 		obj.Right.AddRight(val)
// 	}
// }

// func (obj *BinaryTree) Print(val interface{}) {

// 	curr := obj

// 	for curr != nil {

// 	}

// }

// // func (obj *BinaryTree) ToArray() (result []interface{}) {

// // 	left := obj.Left

// // 	right := obj.Right

// // 	emitVal := func(bin *BinaryTree, arr *[]interface{}) {

// // 		if bin == nil {
// // 			return
// // 		}

// // 		arr = append(arr)
// // 	}

// // 	for left != nil {

// // 	}

// // 	for right != nil {

// // 		right = right.
// // 	}

// // 	return result
// // }

// func (obj *BinaryTree) LeftToArray() (result *[]interface{}) {

// 	return result
// }
