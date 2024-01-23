package main

import "fmt"

type AstNode struct {
	left  *AstNode
	mid   *AstNode
	right *AstNode
	op    int
	v     AstNodeValue
}

type AstNodeValue struct {
	intval int
	id     int
}

func makeAstNode(op int, left, mid, right *AstNode, intval, id int) *AstNode {
	n := AstNode{
		left:  left,
		mid:   mid,
		right: right,
		op:    op,
		v: AstNodeValue{
			intval: intval,
			id:     id,
		},
	}
	return &n
}

func makeLeaf(op int, intval, id int) *AstNode {
	return makeAstNode(op, nil, nil, nil, intval, id)
}

func mkastunary(op int, left *AstNode, intval, id int) *AstNode {
	return makeAstNode(op, left, nil, nil, intval, id)
}

func printAstDepth(n AstNode) {
	fmt.Printf("Height %d\n", getDepth(&n))

}
func getDepth(n *AstNode) int {
	if n == nil {
		return 0
	}

	depth := 1
	depthLeft := getDepth(n.left)
	depthRight := getDepth(n.right)

	if depthLeft > depthRight {
		depth += depthLeft
	} else {
		depth += depthRight
	}

	return depth
}
