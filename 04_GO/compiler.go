package main

import (
	"fmt"
	"log"
)

var AstOps = []string{"+", "-", "*", "/"}

func interpretAST(n *AstNode) string {
	var leftreg, rightreg string
	if n.left != nil {
		leftreg = interpretAST(n.left)
	}
	if n.right != nil {
		rightreg = interpretAST(n.right)
	}

	if n.op == A_INTLIT {
		fmt.Printf("int %d\n", n.intval)
	} else {
		fmt.Printf("%s %s %s\n", leftreg, AstOps[n.op], rightreg)
	}
	switch n.op {
	case A_ADD:
		return genAdd(leftreg, rightreg)
	case A_SUBTRACT:
		return genSub(leftreg, rightreg)
	case A_MULTIPLY:
		return genMul(leftreg, rightreg)
	case A_DIVIDE:
		return genDiv(leftreg, rightreg)
	case A_INTLIT:
		return genNumber(n)
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		panic("Unknown AST operator")
	}
}
