// Extremely simple postfix parser
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var (
	reader    *bufio.Reader = bufio.NewReader(os.Stdin)
	lookahead int           = getToken()
)

func expr() {
	term()
	for {
		if lookahead == '+' {
			match('+')
			term()
			fmt.Print("+")
		} else if lookahead == '-' {
			match('-')
			term()
			fmt.Print("-")
		} else {
			return
		}
	}
}

func term() {
	if unicode.IsDigit(rune(lookahead)) {
		fmt.Print(string(lookahead))
		match(lookahead)
	} else {
		fmt.Println("unexpected character in term:", string(lookahead))
		os.Exit(1)
	}
}

func match(token int) {
	if lookahead == token {
		lookahead = getToken()
	} else {
		fmt.Println("unexpected character in match:", string(lookahead))
		os.Exit(1)
	}
}

func getToken() int {
	r, _, e := reader.ReadRune()
	if e != nil {
		fmt.Println("error:", e)
		os.Exit(1)
	}
	return int(r)
}

func main() {
	expr()
	fmt.Println()
}
