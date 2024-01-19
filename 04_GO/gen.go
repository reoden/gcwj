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
	fmt.Fprintf(OutputFile, "package main \n\nfunc main(){\n")
}

func genMainFuncEnd() {
	fmt.Fprintf(OutputFile, "println(%s)\n}", getLastGenVariable())
}

func getLastGenVariable() string {
	return "v" + strconv.Itoa(VariableCounter)
}

func generateVariable() string {
	VariableCounter++
	return getLastGenVariable()
}

func generateMathExpression(operator string, left, right string) string {
	variableName := generateVariable()
	fmt.Fprintf(OutputFile, "%s := (%v %s %v)\n", variableName, left, operator, right)
	return variableName
}

func genAdd(left, right string) string {
	return generateMathExpression("+", left, right)
}

func genSub(left, right string) string {
	return generateMathExpression("-", left, right)
}

func genMul(left, right string) string {
	return generateMathExpression("*", left, right)
}

func genDiv(left, right string) string {
	return generateMathExpression("/", left, right)
}

func genNumber(node *AstNode) string {
	varialbleName := generateVariable()
	fmt.Fprintf(OutputFile, "%s := %s\n", varialbleName, strconv.Itoa(node.intval))
	return varialbleName
}
