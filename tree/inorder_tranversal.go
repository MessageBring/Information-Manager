package tree

func InorderTranversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, InorderTranversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InorderTranversal(root.Right)...)
	return
}
