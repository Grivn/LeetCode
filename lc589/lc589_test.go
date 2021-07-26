package lc589

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, root.Val)

	for _, child := range root.Children {
		cRes := preorder(child)
		res = append(res, cRes...)
	}

	return res
}

