package main

import (
	"fmt"
	"log"
	"os"
)

func parseStatements() {
	for 1 == 1 {
		for isCurrentTokenNewLine() {
			matchNewLine()
		}

		switch T.token {
		case T_PRINT:
			printStatement()
			break
		case T_VAR:
			varDeclaration()
			break
		case T_IDENT:
			assignmentStatement()
			break
		case T_EOF:
			return
		default:
			log.Fatalf("Syntax error, token line %d column %d\n", Line, Column)
			os.Exit(6)
		}

		if T.token == T_EOF {
			return
		}
		OutputFile.Flush()
	}
}

func assignmentStatement() {
	var left, right, tree *AstNode
	matchIdent()
	id, err := findGlobalSymbol(LastScannedIdent)
	if err != nil {
		log.Fatalf("Undeclared variable %s", LastScannedIdent)
		os.Exit(7)
	}

	left = makeLeaf(A_ASSIGNVAL, -1, id)
	matchToken(T_EQ, "=")
	right = binExpr(0)
	tree = makeAstNode(A_ASSIGN, left, right, 0, -1)
	fmt.Fprint(OutputFile, interpretAST(tree))
	OutputFile.Flush()
}

func printStatement() {
	var n *AstNode
	matchToken(T_PRINT, "print")
	n = binExpr(0)
	genPrint(interpretAST(n))
}
