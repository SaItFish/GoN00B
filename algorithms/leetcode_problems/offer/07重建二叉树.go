package offer

type TreeNode07 struct {
	Val   int
	Left  *TreeNode07
	Right *TreeNode07
}

func buildTree07(preorder []int, inorder []int) *TreeNode07 {
	var recur func(preRoot int, inLeft int, inRight int) *TreeNode07
	var ino = make(map[int]int)

	// 建立中序遍历的索引，前提是这个树无重复值
	for i, x := range inorder {
		ino[x] = i
	}

	recur = func(preRoot int, inLeft int, inRight int) *TreeNode07 {
		if inLeft > inRight {
			return nil // 中序遍历为空，直接返回
		}
		root := &TreeNode07{Val: preorder[preRoot]}          // 生成当前树的节点
		i := ino[preorder[preRoot]]                          // 找到该节点在中序遍历中的位置
		root.Left = recur(preRoot+1, inLeft, i-1)            // 左子树根节点 = 根节点索引 + 1，[inLeft:i-1]
		root.Right = recur(i-inLeft+preRoot+1, i+1, inRight) // 右子树根节点 = 左子树长度 + 根节点索引 + 1，[i+1:inRight]
		return root
	}

	return recur(0, 0, len(preorder)-1)
}
