package main

import "fmt"

func varDeclaration() {
	matchToken(T_VAR, "v")
	matchIdent()
	id := addGlobalSymbol(LastScannedIdent)
	genGlobalSymbol()

	if T.token == T_ASSIGN {
		matchToken(T_ASSIGN, "=")
		var left, right, tree *AstNode
		left = makeLeaf(A_ASSIGNVAL, -1, id)
		right = binExpr(0)
		tree = makeAstNode(A_ASSIGN, left, right, 0, -1)
		fmt.Fprintf(OutputFile, interpretAST(tree))
	}
}
