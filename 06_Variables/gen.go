package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var VariableCounter = 0

func openOutputFile(path string) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	OutputFilePtr = file
	OutputFile = bufio.NewWriter(file)
}

func genMainFuncStart() {
	fmt.Fprintf(OutputFile, "package main \n\n import \"fmt\" \n\nfunc main(){\n")
}

func genPrint(value string) {
	fmt.Fprintf(OutputFile, "fmt.Printf(\"%%v\\n\", %s)\n", value)
}

func genMainFuncEnd() {
	fmt.Fprintf(OutputFile, "}")
}

func getLastGenVariable() string {
	return "v" + strconv.Itoa(VariableCounter)
}

func generateVariable() string {
	VariableCounter++
	return getLastGenVariable()
}

func genMathExpression(left, op, right string) string {
	return fmt.Sprintf("(%s %s %s)", left, op, right)
}

func genAdd(left, right string) string {
	return genMathExpression(left, "+", right)
}

func genSub(left, right string) string {
	return genMathExpression(left, "-", right)
}

func genMul(left, right string) string {
	return genMathExpression(left, "*", right)
}

func genDiv(left, right string) string {
	return genMathExpression(left, "/", right)
}

func genNumber(node *AstNode) string {
	return fmt.Sprintf("%s", strconv.Itoa(node.v.intval))
}

func genGlobalSymbol() {
	fmt.Fprintf(OutputFile, "var %s int\n", GlobalSymbols[findLastGlobalSymbol()].name)
}

func genAssignVal(n *AstNode) string {
	return fmt.Sprintf("%s", GlobalSymbols[n.v.id].name)
}

func genIdent(n *AstNode) string {
	return fmt.Sprintf("%s", GlobalSymbols[n.v.id].name)
}

func genAssign(left, right string) string {
	return fmt.Sprintf("%s = %s\n", left, right)
}
