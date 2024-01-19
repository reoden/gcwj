package main

import (
	"fmt"
	"log"
	"os"
)

var AstOps = []string{"+", "-", "*", "/"}

func interpretAST(n *AstNode) int {
	var leftval, rightval int

	if n.left != nil {
		leftval = interpretAST(n.left)
	}
	if n.right != nil {
		rightval = interpretAST(n.right)
	}

	if n.op == A_INTLIT {
		fmt.Printf("int %d\n", n.intval)
	} else {
		fmt.Printf("%d %s %d\n", leftval, AstOps[n.op], rightval)
	}

	switch n.op {
	case A_ADD:
		return (leftval + rightval)
	case A_SUBTRACT:
		return (leftval - rightval)
	case A_MULTIPLY:
		return (leftval * rightval)
	case A_DIVIDE:
		return (leftval / rightval)
	case A_INTLIT:
		return n.intval
	default:
		log.Fatalf("Unknown AST operator %v\n", n.op)
		os.Exit(5)
	}
	return -1
}
