package main

import (
	"log"
	"os"
)

func opPrecedence(tokentype int) int {
	prec := OpPrec[tokentype]
	if prec == 0 {
		log.Fatalf("syntax / Precedence error on line %d column %d, token %d\n", Line, Column, tokentype)
		os.Exit(6)
	}

	return prec
}

func getAstType(tok int) int {
	switch tok {
	case T_PLUS:
		return (A_ADD)
	case T_MINUS:
		return (A_SUBTRACT)
	case T_STAR:
		return (A_MULTIPLY)
	case T_SLASH:
		return (A_DIVIDE)
	default:
		log.Fatalf("unknown token in getAstType() on line %d column %d\n", Line, Column)
		os.Exit(3)
	}
	return -1
}

func primary() *AstNode {
	var n *AstNode

	switch T.token {
	case T_INTLIT:
		n = makeLeaf(A_INTLIT, T.intvalue, -1)
		break
	case T_IDENT:
		id, err := findGlobalSymbol(LastScannedIdent)
		if err != nil {
			log.Fatalf("Unknown variable %s, %v\n", LastScannedIdent, err)
		}
		n = makeLeaf(A_IDENT, -1, id)
		break
	default:
		log.Fatalf("Syntax error on line %d column %d ", Line, Column)
		os.Exit(4)
	}
	scan(&T)
	return n
}

func binExpr(ptp int) *AstNode {
	var left, right *AstNode
	var tokentype int

	left = primary()
	tokentype = T.token
	if tokentype == T_NEWLINE || tokentype == T_EOF {
		return left
	}

	for opPrecedence(tokentype) > ptp {
		scan(&T)

		right = binExpr(OpPrec[tokentype])
		left = makeAstNode(getAstType(tokentype), left, right, 0, -1)
		tokentype = T.token
		if tokentype == T_NEWLINE || tokentype == T_EOF {
			return left
		}
	}

	return left
}
