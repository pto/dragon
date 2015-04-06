package lexer

var (
	Line  int             = 1                        // Current line number
	peek  rune            = ' '                      // Next source char or ' '
	words map[string]Word = make(map[string]Word, 2) // Symbol table
)

func reserve(t Word) {
	words[t.lexeme] = t
}

func init() {
	reserve(NewWord(TRUE, "true"))
	reserve(NewWord(FALSE, "false"))
}

func scan() Token {
}
