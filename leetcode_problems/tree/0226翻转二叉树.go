// Package tree contains some leetcode tree questions
// @file: 0226翻转二叉树.go
// @description: 翻转一棵二叉树。
// @author: SaltFish
// @date: 2020/09/10

/*
示例：
输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

package tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Left)
	left := invertTree(root.Right)
	root.Right = right
	root.Left = left
	return root
}
