package main

import (
	"fmt"
	"log"
	"os"
)

func compundStatements() *AstNode {
	var left, tree *AstNode
	left = nil
	for isCurrentTokenNewLine() {
		matchNewLine()
	}

	matchLBrace()
	for 1 == 1 {
		for isCurrentTokenNewLine() {
			matchNewLine()
		}

		switch T.token {
		case T_PRINT:
			tree = printStatement()
			break
		case T_VAR:
			varDeclaration()
			tree = nil
			break
		case T_IDENT:
			tree = assignmentStatement()
			break
		case T_IF:
			tree = ifStatement()
			break
		case T_RBRACE:
			matchRBrace()
			return left
		case T_WHILE:
			tree = whileStatement()
			break
		default:
			log.Fatalf("Syntax error, token line %d column %d\n", Line, Column)
			os.Exit(6)
		}

		OutputFile.Flush()
		if tree != nil {
			if left == nil {
				left = tree
			} else {
				left = makeAstNode(A_GLUETO, left, nil, tree, -1, -1)
			}
		}
	}

	return nil
}

func assignmentStatement() *AstNode {
	var left, right, tree *AstNode
	matchIdent()
	id, err := findGlobalSymbol(LastScannedIdent)
	if err != nil {
		log.Fatalf("Undeclared variable %s", LastScannedIdent)
		os.Exit(7)
	}

	left = makeLeaf(A_ASSIGNVAL, -1, id)
	matchToken(T_ASSIGN, "=")
	right = binExpr(0)
	tree = makeAstNode(A_ASSIGN, left, nil, right, 0, -1)
	fmt.Fprint(OutputFile, interpretAST(tree))
	OutputFile.Flush()

	return tree
}

func printStatement() *AstNode {
	var n *AstNode
	matchToken(T_PRINT, "print")
	n = binExpr(0)
	return mkastunary(A_PRINT, n, -1, -1)
}

func ifStatement() *AstNode {
	var condAST, trueAST, falseAST *AstNode
	matchToken(T_IF, "if")
	matchLPAREN()

	condAST = binExpr(0)

	if condAST.op < A_EQ || condAST.op > A_GE {
		log.Fatalf("Comparison is not returning a boolean on line %d\n", Line)
	}

	matchRPAREN()
	trueAST = compundStatements()

	if T.token == T_ELSE {
		scan(&T)
		falseAST = compundStatements()
	}

	return makeAstNode(A_IF, condAST, trueAST, falseAST, -1, -1)
}

func whileStatement() *AstNode {
	var condAST, bodyAst *AstNode
	matchToken(T_WHILE, "while")
	matchLPAREN()

	condAST = binExpr(0)
	if condAST.op < A_EQ || condAST.op > A_GE {
		log.Fatalf("Condition is not returning a boolean on line %d\n", Line)
	}

	matchRPAREN()
	bodyAst = compundStatements()

	return makeAstNode(A_WHILE, condAST, bodyAst, nil, -1, -1)
}
