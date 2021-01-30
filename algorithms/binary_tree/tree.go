// Package binary_tree
// @file: tree.go
// @date: 2021/1/19
package binary_tree

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 116. 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	var connectTwoNodes func(*Node, *Node)
	connectTwoNodes = func(node1, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}
		node1.Next = node2
		connectTwoNodes(node1.Left, node1.Right)
		connectTwoNodes(node1.Right, node2.Left)
		connectTwoNodes(node2.Left, node2.Right)
	}
	if root == nil {
		return nil
	}
	connectTwoNodes(root.Left, root.Right)
	return root
}

// 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 左右子树已经被拉成链表
	left, right := root.Left, root.Right
	// 将左子树接到右侧
	root.Left, root.Right = nil, left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	// 右子树接到左子树后
	p.Right = right
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	type resultType struct {
		singlePath int // 保存单边最大值
		maxPath    int // 保存最大值（单边或者两个单边+根的值）
	}

	var helper func(*TreeNode) resultType
	helper = func(root *TreeNode) resultType {
		if root == nil {
			return resultType{singlePath: 0, maxPath: math.MinInt64}
		}
		left := helper(root.Left)
		right := helper(root.Right)

		res := resultType{}
		if left.singlePath > right.singlePath {
			res.singlePath = max(left.singlePath+root.Val, 0)
		} else {
			res.singlePath = max(right.singlePath+root.Val, 0)
		}
		maxPath := max(left.maxPath, right.maxPath)
		res.maxPath = max(left.singlePath+right.singlePath+root.Val, maxPath)
		return res
	}

	return helper(root).maxPath
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}
	reverse(res)
	return res
}

func reverse(list [][]int) {
	length := len(list)
	for i := 0; i < length/2; i++ {
		list[i], list[length-1-i] = list[length-1-i], list[i]
	}
}

// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
	helper(root)
	if len(res) <= 1 {
		return true
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	p := root
	for {
		if p.Val >= val {
			if p.Left != nil {
				p = p.Left
			} else {
				p.Left = node
				break
			}
		} else {
			if p.Right != nil {
				p = p.Right
			} else {
				p.Right = node
				break
			}
		}
	}

	return root
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		if left == -1 || right == -1 || abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	if helper(root) == -1 {
		return false
	}
	return true
}
