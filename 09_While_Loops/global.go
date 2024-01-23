package main

import (
	"bufio"
	"os"
)

var Line int
var Column int
var Putback rune
var InputFilePtr *os.File
var OutputFilePtr *os.File
var OutputFile *bufio.Writer
var InputFile *bufio.Reader
var TokenStr = []string{"+", "-", "*", "/", "intlit"}
var T Token
var OpPrec = []int{
	0, 10, 10, // T_EOF, T_PLUS, T_MINUS
	20, 20, // T_STAR, T_SLASH
	30, 30, // T_EQ, T_NEQ
	40, 40, 40, 40, //T_LT, T_GT, T_LE, T_GE
}
var LastScannedIdent string

const (
	T_EOF = iota
	// calc
	T_PLUS
	T_MINUS
	T_STAR
	T_SLASH
	T_EQ
	T_NEQ
	T_LT
	T_GT
	T_LE
	T_GE

	//no precedence
	T_INTLIT
	T_NEWLINE
	T_PRINT
	T_ASSIGN
	T_LBRACE
	T_RBRACE
	T_LPAREN
	T_RPAREN

	// keyword
	T_VAR
	T_IDENT
	T_IF
	T_ELSE
	T_WHILE
)

const (
	A_EOF_PLACEHOLDED = iota
	A_ADD
	A_SUBTRACT
	A_MULTIPLY
	A_DIVIDE
	A_EQ
	A_NEQ
	A_LT
	A_GT
	A_LE
	A_GE

	A_INTLIT
	A_ASSIGNVAL
	A_ASSIGN
	A_IDENT
	A_PRINT
	A_GLUETO
	A_IF
	A_WHILE
)
