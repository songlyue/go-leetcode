package main

import "fmt"

/*
*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
中左右
*/
func main() {
	ints := test()
	fmt.Println("%v", ints)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//// 同一迭代法
//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	res := make([]int, 0)
//	stack := list.New()
//	// 把根节点加入栈中
//	stack.PushBack(root)
//	for stack.Len() > 0 {
//		// 每次都把中节点弹出
//		e := stack.Back()
//		stack.Remove(e)
//		if e.Value == nil {
//			e = stack.Back()
//			stack.Remove(e)
//			node := e.Value.(*TreeNode)
//			res = append(res, node.Val)
//			continue
//		}
//		// 最晚放入的元素
//		node := e.Value.(*TreeNode)
//		if node.Right != nil {
//			stack.PushBack(node.Right)
//		}
//		if node.Left != nil {
//			stack.PushBack(node.Left)
//		}
//		// nil的前面 第一个放入res中的元素
//		stack.PushBack(node)
//		stack.PushBack(nil) // 压栈
//
//	}
//	return res
//}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}
