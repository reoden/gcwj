package main

import (
	"bufio"
	"os"
)

var Line int
var Putback int
var Column int
var FilePtr *os.File
var File *bufio.Reader
var TokenStr = []string{"+", "-", "*", "/", "intlit"}
