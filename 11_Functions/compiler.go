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
		return leftreg + rightreg + "\n"
	case A_WHILE:
		return genWhile(n)
	case A_FUNC:
		return genFunction(n)
	case A_FUNCALL:
		return genFunctionCall(n)
	default:
		log.Fatalf("Unknown AST operator %d\n", n.op)
		panic("Unknown AST operator")
	}
}

func genFunction(node *AstNode) string {
	fnhead := fmt.Sprintf("%v = func(){\n", GlobalSymbols[node.v.id].name)
	fnBody := fmt.Sprintf("%v }\n", interpretAST(node.left))
	return fnhead + fnBody
}

func genFunctionCall(node *AstNode) string {
	fnCall := fmt.Sprintf("%v()\n", GlobalSymbols[node.v.id].name)
	return fnCall
}

func genWhile(node *AstNode) string {
	whilehead := fmt.Sprintf("for %v {\n", interpretAST(node.left))
	trueBody := fmt.Sprintf(" %v }\n", interpretAST(node.mid))
	return whilehead + trueBody
}

func genIf(node *AstNode) string {
	ifhead := fmt.Sprintf("if %v {\n", interpretAST(node.left))
	trueBody := fmt.Sprintf(" %v }\n", interpretAST(node.mid))
	falseBody := ""
	if node.right != nil {
		falseBody = fmt.Sprintf("else {\n %v }\n", interpretAST(node.right))
	}

	return ifhead + trueBody + falseBody
}
