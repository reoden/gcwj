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

	switch n.op {
	case A_ADD:
		return genAdd(leftreg, rightreg)
	case A_SUBTRACT:
		return genSub(leftreg, rightreg)
	case A_MULTIPLY:
		return genMul(leftreg, rightreg)
	case A_DIVIDE:
		return genDiv(leftreg, rightreg)
	case A_EQ:
		return genEq(leftreg, rightreg)
	case A_NEQ:
		return genNeq(leftreg, rightreg)
	case A_LT:
		return genLt(leftreg, rightreg)
	case A_GT:
		return genGt(leftreg, rightreg)
	case A_LE:
		return genLe(leftreg, rightreg)
	case A_GE:
		return genGe(leftreg, rightreg)
	case A_INTLIT:
		return genNumber(n)
	case A_ASSIGN:
		return genAssign(leftreg, rightreg)
	case A_ASSIGNVAL:
		return genAssignVal(n)
	case A_IDENT:
		return genIdent(n)
	case A_IF:
		return genIf(n)
	case A_PRINT:
		return genPrint(leftreg)
	case A_GLUETO:
		writeOutput(leftreg)
		writeOutput(rightreg)
		writeOutput("\n")
		return ""
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		panic("Unknown AST operator")
	}
}

func genIf(node *AstNode) string {
	ifhead := fmt.Sprintf("if %v {", interpretAST(node.left))
	trueBody := fmt.Sprintf(" %v }", interpretAST(node.mid))
	falseBody := ""
	if node.right != nil {
		falseBody = fmt.Sprintf("else {\n %v }", interpretAST(node.right))
	}

	return ifhead + trueBody + falseBody
}
