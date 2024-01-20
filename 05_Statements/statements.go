package main

func parseStatements() {
	var tree *AstNode
	for 1 == 1 {
		for isCurrentTokenNewLine() {
			matchNewLine()
		}
		matchToken(T_PRINT, "print")
		tree = binExpr(0)
		interpretAST(tree)
		genPrint()
		if T.token == T_EOF {
			return
		}
		//		matchNewLine()
	}
}
