package main

import "errors"

type Symtable struct {
	name string
}

var GlobalSymbols = []Symtable{}

func addGlobalSymbol(name string) int {
	GlobalSymbols = append(GlobalSymbols, Symtable{name: name})
	return len(GlobalSymbols) - 1
}

func findGlobalSymbol(name string) (int, error) {
	for i := 0; i < len(GlobalSymbols); i++ {
		if GlobalSymbols[i].name == name {
			return i, nil
		}
	}
	return -1, errors.New("Could not find global symbol")
}

func getNextGlobalPosition() int {
	return len(GlobalSymbols)
}

func findLastGlobalSymbol() int {
	return len(GlobalSymbols) - 1
}
