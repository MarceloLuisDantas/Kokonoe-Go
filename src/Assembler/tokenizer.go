package assembler

type TokenType int

const (
	INSTRUCTION TokenType = iota
	REGISTER
	LABEL_DEF
	LABEL_REF
	SECTION
	TYPE
	STRING
	NUMBER
	VIRGULA
	NEW_LINE
	OPEN_PARENTH
	CLOSE_PARENTH
)

type Token struct {
	TokenType TokenType
	Value     string
	Line      int
	Column    int
}

func newToken(tt TokenType, value string, line int, column int) *Token {
	t := Token{tt, value, line, column}
	return &t
}

type Tokenizer struct {
	Data     string
	position int
	column   int
	line     int
}

func newTokenizer(data string) *Tokenizer {
	tokenizer := Tokenizer{data, 0, 0, 0}
	return &tokenizer
}

func (tokenizer *Tokenizer) advance() {
	tokenizer.position += 1
	tokenizer.column += 1
}

func (Tokenizer *Tokenizer) advanceX(value int) {
	Tokenizer.position += value
	Tokenizer.column += value
}

func (tokenizer *Tokenizer) nextLine() {
	tokenizer.position += 1
	tokenizer.column = 0
	tokenizer.line += 1
}

func (tokenizer *Tokenizer) getCurrentChar() byte {
	return tokenizer.Data[tokenizer.position]
}

func (tokenizer *Tokenizer) Tokenize() {
	for tokenizer.position != len(tokenizer.Data) {
		tokenizer.advance()
	}
}
