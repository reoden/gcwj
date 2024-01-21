package main

import (
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
	case A_ASSIGN:
		return genAssign(leftreg, rightreg)
	case A_ASSIGNVAL:
		return genAssignVal(n)
	case A_IDENT:
		return genIdent(n)
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		panic("Unknown AST operator")
	}
}
