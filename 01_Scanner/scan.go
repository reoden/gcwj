package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func next() (rune, error) {
	var c rune

	if Putback > 0 {
		c = rune(Putback)
		Putback = 0
		return c, nil
	}

	c, _, err := File.ReadRune()
	Column++

	if err != nil {
		if err == io.EOF {
			return 0, err
		} else {
			log.Fatal(err)
			panic(err)
		}
	}

	if c == '\n' {
		Line++
	}

	return c, nil
}

func putback(c rune) {
	Putback = int(c)
}

func skip() (rune, error) {
	c, err := next()

	if err != nil {
		return 0, err
	}

	for c == ' ' || c == '\t' || c == '\n' || c == '\f' || c == '\r' {
		c, err = next()
		if err != nil {
			return 0, err
		}
	}

	return c, nil
}

func scanint(c rune) int {
	var err error
	var val = 0

	for unicode.IsDigit(c) {
		val = val*10 + int(c-'0')
		c, err = next()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Could not read integer. Line %d Col %d Char: %q", Line, Column, c)
			} else {
				break
			}
		}
	}
	putback(c)
	return val
}

func scan(t *Token) bool {
	c, err := skip()
	if err != nil {
		return false
	}

	switch c {
	case '+':
		t.token = T_PLUS
		break
	case '-':
		t.token = T_MINUS
		break
	case '*':
		t.token = T_STAR
		break
	case '/':
		t.token = T_SLASH
		break
	default:
		if unicode.IsDigit(c) {
			t.intvalue = scanint(c)
			t.token = T_INTLIT
		} else {
			fmt.Printf("Unrecognised character %c on line %d\n", c, Line)
			os.Exit(2)
		}
	}

	return true
}

func scanFile() {
	var t Token
	for scan(&t) {
		fmt.Printf("Toekn %q", TokenStr[t.token])
		if t.token == T_INTLIT {
			fmt.Printf(", value %d", t.intvalue)
		}
		fmt.Println()
	}
}

func openfile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	FilePtr = file
	File = bufio.NewReader(file)
}
