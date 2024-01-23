package main

import (
	"log"
	"os"
)

func singleStatements() *AstNode {
	var tree *AstNode = nil
	switch T.token {
	case T_PRINT:
		tree = printStatement()
	case T_FUNC:
		tree = functionDeclaration()
	case T_VAR:
		tree = varDeclaration()
	case T_IDENT:
		tree = identifierStatement()
	case T_IF:
		tree = ifStatement()
	case T_WHILE:
		tree = whileStatement()
	case T_FOR:
		tree = forStatement()
	default:
		log.Fatalf("Syntax error, token line %d column %d\n", Line, Column)
	}
	return tree
}

func compundStatements() *AstNode {
	var left, tree *AstNode
	left = nil
	for isCurrentTokenNewLine() {
		matchNewLine()
	}

	matchLBrace()
	for {
		for isCurrentTokenNewLine() {
			matchNewLine()
		}

		tree = singleStatements()

		if tree != nil {
			if left == nil {
				left = tree
			} else {
				left = makeAstNode(A_GLUETO, left, nil, tree, -1, -1)
			}
		}

		for isCurrentTokenNewLine() {
			matchNewLine()
		}

		if T.token == T_RBRACE {
			matchRBrace()
			return left
		}
	}
}

func assignmentStatement(symbolId int) *AstNode {
	var left, right, tree *AstNode
	left = makeLeaf(A_ASSIGNVAL, -1, symbolId)
	matchToken(T_ASSIGN, "=")
	right = binExpr(0)
	tree = makeAstNode(A_ASSIGN, left, nil, right, 0, -1)

	return tree
}

func identifierStatement() *AstNode {
	matchIdent()
	id, err := findGlobalSymbol(LastScannedIdent)
	if err != nil {
		log.Fatalf("Undeclared variable %s", LastScannedIdent)
		os.Exit(7)
	}
	switch GlobalSymbols[id].symType {
	case TYPE_FUNC:
		return functionCallStatement(id)
	case TYPE_INT:
		return assignmentStatement(id)
	default:
		log.Fatalf("Unknown type for identifier %v\n", GlobalSymbols[id].name)
		os.Exit(8)
	}

	return nil
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

func forStatement() *AstNode {
	var condAST, bodyAst, preopAst, postopAst, tree *AstNode
	matchToken(T_FOR, "for")
	matchLPAREN()

	preopAst = singleStatements()
	matchSemi()

	condAST = binExpr(0)
	if condAST.op < A_EQ || condAST.op > A_GE {
		log.Fatalf("condition is not returing a boolean on line %d\n", Line)
	}
	matchSemi()

	postopAst = singleStatements()
	matchRPAREN()

	bodyAst = compundStatements()

	tree = makeAstNode(A_GLUETO, bodyAst, nil, postopAst, -1, -1)
	tree = makeAstNode(A_WHILE, condAST, tree, nil, -1, -1)
	tree = makeAstNode(A_GLUETO, preopAst, nil, tree, -1, -1)
	return tree
}

func functionCallStatement(symbolId int) *AstNode {
	matchLPAREN()
	matchRPAREN()
	var tree *AstNode
	makeLeaf(A_FUNCALL, -1, symbolId)
	return mkastunary(A_FUNCALL, tree, -1, symbolId)
}
