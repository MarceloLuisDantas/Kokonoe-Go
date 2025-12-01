package assembler

var INSTRUCTIONS_SET = map[string]bool{
	"add":     true,
	"addi":    true,
	"sub":     true,
	"subi":    true,
	"mult":    true,
	"multi":   true,
	"div":     true,
	"divi":    true,
	"move":    true,
	"or":      true,
	"ori":     true,
	"and":     true,
	"andi":    true,
	"sll":     true,
	"srl":     true,
	"slt":     true,
	"slti":    true,
	"li":      true,
	"syscall": true,
	"j":       true,
	"jr":      true,
	"jal":     true,
	"beq":     true,
	"bne":     true,
	"bgt":     true,
	"bge":     true,
	"blt":     true,
	"ble":     true,
	"return":  true,
	"lw":      true,
	"lb":      true,
	"sw":      true,
	"sb":      true,
	"lv":      true,
	"sv":      true,
	"lrw":     true,
	"lrb":     true,
	"inc":     true,
	"dec":     true,
	".text":   true,
	".data":   true,
	"la":      true,
	"rand":    true,
}

type TokenType int

const (
	IDENTIFIER TokenType = iota
	INSTRUCTION
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
	len      int
	tokens   []Token
}

func newTokenizer(data string) *Tokenizer {
	tokenizer := Tokenizer{data, 0, 0, 0, 0, []Token{}}
	return &tokenizer
}

func (tokenizer *Tokenizer) addToken(token_type TokenType, value string) {
	token := newToken(token_type, value, tokenizer.column, tokenizer.line)
	tokenizer.tokens = append(tokenizer.tokens, *token)
	tokenizer.len += 1
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

func (tokenizer *Tokenizer) handleSpace() {
	tokenizer.advance()
}

func (tokenizer *Tokenizer) handleNewLine() {
	tokenizer.addToken(NEW_LINE, "\n")
	tokenizer.nextLine()
}

func isValidCharacterToIdentifier(s byte) bool {
	matched := (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || s == '_' || (s >= '0' && s <= '9')
	return matched
}

func (tokenizer *Tokenizer) handleIdentifier() {
	start := tokenizer.position

	for isValidCharacterToIdentifier(tokenizer.getCurrentChar()) {
		tokenizer.advance()
	}

	identifier := tokenizer.Data[start:tokenizer.position]
	tokenizer.addToken(IDENTIFIER, identifier)
}

func (tokenizer *Tokenizer) Tokenize() {
	for tokenizer.position != len(tokenizer.Data) {
		current := tokenizer.getCurrentChar()
		if current == ' ' {
			tokenizer.handleSpace()
		} else if current == '\n' {
			tokenizer.handleNewLine()
		} else if (current >= 'a' && current <= 'z') || (current >= 'A' && current <= 'Z') || current == '_' {
			tokenizer.handleIdentifier()
		}
	}
}
