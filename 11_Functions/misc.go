package main

func matchLBrace() {
	matchToken(T_LBRACE, "{")
}

func matchRBrace() {
	matchToken(T_RBRACE, "}")
}

func matchLPAREN() {
	matchToken(T_LPAREN, "(")
}

func matchRPAREN() {
	matchToken(T_RPAREN, ")")
}

func matchNewLine() {
	matchToken(T_NEWLINE, "\\n")
}

func matchIdent() {
	matchToken(T_IDENT, "identifier")
}

func matchSemi() {
	matchToken(T_SEMI, ";")
}
