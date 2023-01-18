package main

/*
*给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 */
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left //交换
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return root
}
