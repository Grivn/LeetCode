package lc590

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int

	if root == nil {
		return res
	}

	for _, child := range root.Children {
		cRes := postorder(child)
		res = append(res, cRes...)
	}

	res = append(res, root.Val)

	return res
}
