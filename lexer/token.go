package lexer

// A Token identifies the type of a token value. It either contains
// the Unicode value of a single character token, or one of the constants
// NUM, ID, TRUE or FALSE.
type Token struct {
	tag int
}

func NewToken(t int) Token {
	return Token{t}
}

const (
	NUM   int = 0x20000 + iota // Number
	ID                         // Identifier
	TRUE                       // Keyword true
	FALSE                      // Keyword false
)

// A Number token represents an integer.
type Num struct {
	Token
	value int
}

func NewNum(v int) Num {
	return Num{Token: NewToken(NUM), value: v}
}

// A Word token represents an identifier or keyword.
type Word struct {
	Token
	lexeme string
}

func NewWord(t int, s string) Word {
	return Word{Token: NewToken(t), lexeme: s}
}
