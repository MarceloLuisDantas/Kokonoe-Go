package assembler

import (
	"fmt"
)

const (
	ADD     string = "add"
	ADDI    string = "addi"
	ADDU    string = "addu"
	ADDUI   string = "addui"
	SUB     string = "sub"
	SUBI    string = "subi"
	SUBU    string = "subu"
	SUBUI   string = "subui"
	MULT    string = "mult"
	MULTI   string = "multi"
	DIV     string = "div"
	DIVI    string = "divi"
	OR      string = "or"
	ORI     string = "ori"
	AND     string = "and"
	ANDI    string = "andi"
	SLL     string = "sll"
	SRL     string = "srl"
	SLT     string = "slt"
	SLTI    string = "slti"
	LI      string = "li"
	LA      string = "la"
	MOVE    string = "move"
	J       string = "j"
	JR      string = "jr"
	JAL     string = "jal"
	BEQ     string = "beq"
	BNE     string = "bne"
	BGT     string = "bgt"
	BGE     string = "bge"
	BLT     string = "blt"
	BLE     string = "ble"
	LW      string = "lw"
	LB      string = "lb"
	SW      string = "sw"
	SB      string = "sb"
	LV      string = "lv"
	SV      string = "sv"
	LRW     string = "lrw"
	LRB     string = "lrb"
	INC     string = "inc"
	DEC     string = "dec"
	SYSCALL string = "syscall"
	RETURN  string = "return"
	RAND    string = "rand"
)

var RegistersInstructions = map[string]bool{
	ADD: true, ADDU: true, SUB: true, SUBU: true,
	MULT: true, DIV: true, OR: true, AND: true, SLT: true,
}

var ImmediateInstructions = map[string]bool{
	ADDI: true, ADDUI: true, SUBI: true, SUBUI: true, MULTI: true, DIVI: true,
	ORI: true, ANDI: true, SLTI: true, SLL: true, SRL: true,
}

var BranchInstructions = map[string]bool{
	BEQ: true, BNE: true, BGT: true, BGE: true, BLT: true, BLE: true,
}

var MemoriInstructions = map[string]bool{
	LW: true, LB: true, SW: true, SB: true, LV: true, SV: true, LRW: true, LRB: true,
}

var JumpInstructions = map[string]bool{
	J: true, JR: true, JAL: true,
}

type Type string

const (
	STR    Type = "string"
	INT8   Type = "int8"
	INT16  Type = "int16"
	UINT8  Type = "uint8"
	UINT16 Type = "uint16"
)

type Section string

const (
	TEXT Section = ".text"
	DATA Section = ".data"
)

type Instruction []string

type Parser struct {
	Tokens       []Token
	Instructions []Instruction
	Position     int
	Len          int
	Labels       map[string]int
	Gp           int
}

func newParser(tokens []Token) *Parser {
	p := Parser{tokens, []Instruction{}, 0, len(tokens), make(map[string]int), 0}
	return &p
}

func (parser *Parser) getCurrentToken() Token {
	return parser.Tokens[parser.Position]
}

func (parser *Parser) getNextXToken(x int) Token {
	return parser.Tokens[parser.Position+x]
}

func invalidRegister(token Token) error {
	fmt.Printf("Registrador esperado em - linha: %d coluna: %d, mas foi encontrado %s.\n",
		token.Line, token.Column, token.TokenType)
	return fmt.Errorf("invalid reg")
}

func invalidImmediate(token Token) error {
	fmt.Printf("Valor imediato esperado em - linha: %d coluna: %d, mas foi encontrado %s.\n",
		token.Line, token.Column, token.TokenType)
	return fmt.Errorf("invalid imme")
}

func invalidLabel(token Token) error {
	fmt.Printf("Label esperado em - linha %d coluna: %d, mas foi encontrado %s.\n",
		token.Line, token.Column, token.TokenType)
	return fmt.Errorf("invalid label")
}

func invalidRegOrLabel(token Token) error {
	fmt.Printf("Label ou Registrador esperando em linha: %d coluna: %d, mas foi encontrado %s.\n",
		token.Line, token.Column, token.TokenType)
	return fmt.Errorf("invalid register or label")
}

func toMuchArgs(token Token, args int) error {
	fmt.Printf("Instrução \"%s\" recebe apena %d argumentos.\n", token.Value, args)
	return fmt.Errorf("to much args")
}

func virgulaMissing(line, column int) error {
	fmt.Printf("Virgula esperada em linha: %d coluna: %d.\n", line, column)
	return fmt.Errorf("missing separator")
}

func openParenMissing(line, column int) error {
	fmt.Printf("Abertura de parenteses esperada ao acessar memoria linha: %d coluna: %d.\n", line, column)
	return fmt.Errorf("missing separator")
}

func closeParenMissing(line, column int) error {
	fmt.Printf("Fechamento de parenteses esperada ao acessar memoria linha: %d coluna: %d.\n", line, column)
	return fmt.Errorf("missing separator")
}

// add, addu, subu, sub, mult, div, or, and, slt
// add $r1, $r2, $r2 \n
func (parser *Parser) parseRegisterInstruciton() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	arg2 := parser.getNextXToken(3)
	if arg2.TokenType != REGISTER {
		return invalidRegister(arg2)
	}

	virgula2 := parser.getNextXToken(4)
	if virgula2.TokenType != VIRGULA {
		return virgulaMissing(virgula2.Line, virgula2.Column)
	}

	arg3 := parser.getNextXToken(5)
	if arg3.TokenType != REGISTER {
		return invalidRegister(arg3)
	}

	nl := parser.getNextXToken(6)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 3)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, arg2.Value)
	ins = append(ins, arg3.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 7
	return nil
}

// addi, addui, subi, subui, multi, divi, ori, andi, slti, sll, srl
// addi $r1, $r2, value \n
func (parser *Parser) parseImediateInstruciton() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	arg2 := parser.getNextXToken(3)
	if arg2.TokenType != REGISTER {
		return invalidRegister(arg2)
	}

	virgula2 := parser.getNextXToken(4)
	if virgula2.TokenType != VIRGULA {
		return virgulaMissing(virgula2.Line, virgula2.Column)
	}

	arg3 := parser.getNextXToken(5)
	if arg3.TokenType != NUMBER {
		return invalidImmediate(arg3)
	}

	nl := parser.getNextXToken(6)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 3)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, arg2.Value)
	ins = append(ins, arg3.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 7
	return nil
}

// j, jr, jal
func (parser *Parser) parseJump() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg := parser.getNextXToken(1)
	if parser.getCurrentToken().Value == "j" || parser.getCurrentToken().Value == "jal" {
		if arg.TokenType != LABEL_REF {
			return invalidLabel(arg)
		}
	} else if parser.getCurrentToken().Value == "jr" {
		if arg.TokenType != REGISTER {
			return invalidRegister(arg)
		}
	}

	nl := parser.getNextXToken(2)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	ins = append(ins, arg.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 3
	return nil
}

// beq, bne, bgt, bge, blt, ble
func (parser *Parser) parseBranch() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	arg2 := parser.getNextXToken(3)
	if arg2.TokenType != REGISTER {
		return invalidRegister(arg2)
	}

	virgula2 := parser.getNextXToken(4)
	if virgula2.TokenType != VIRGULA {
		return virgulaMissing(virgula2.Line, virgula2.Column)
	}

	arg3 := parser.getNextXToken(5)
	if arg3.TokenType != LABEL_REF {
		return invalidLabel(arg3)
	}

	nl := parser.getNextXToken(6)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 3)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, arg2.Value)
	ins = append(ins, arg3.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 7
	return nil
}

// lw, lb, sw, sb, lv, sv, lrw, lrb
// lw $r1, 0(*label)
// lw $r1, 0($r2)
// lw $r1, $r2($r3)
func (parser *Parser) parseMemorie() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	offset := parser.getNextXToken(3)
	if offset.TokenType != NUMBER && offset.TokenType != REGISTER {
		return invalidRegOrLabel(offset)
	}

	open_paren := parser.getNextXToken(4)
	if open_paren.TokenType != OPEN_PAREN {
		return openParenMissing(open_paren.Line, open_paren.Column)
	}

	arg2 := parser.getNextXToken(5)
	if arg2.TokenType != REGISTER {
		return invalidRegister(arg2)
	}

	close_paren := parser.getNextXToken(6)
	if close_paren.TokenType != CLOSE_PAREN {
		return closeParenMissing(close_paren.Line, close_paren.Column)
	}

	nl := parser.getNextXToken(7)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 3)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, offset.Value)
	ins = append(ins, arg2.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 8
	return nil
}

// inc, dec
func (parser *Parser) parseIncDec() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg := parser.getNextXToken(1)
	if arg.TokenType != REGISTER {
		return invalidRegister(arg)
	}

	nl := parser.getNextXToken(2)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	ins = append(ins, arg.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 3
	return nil
}

// Syscall, Return
func (parser *Parser) parseSyscallReturn() error {
	ins := Instruction{parser.getCurrentToken().Value}

	nl := parser.getNextXToken(1)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 2
	return nil
}

func (parser *Parser) parseMove() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	arg2 := parser.getNextXToken(3)
	if arg2.TokenType != REGISTER {
		return invalidRegister(arg2)
	}

	nl := parser.getNextXToken(4)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, arg2.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 5
	return nil
}

func (parser *Parser) parseLi() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	arg2 := parser.getNextXToken(3)
	if arg2.TokenType != NUMBER {
		return invalidImmediate(arg2)
	}

	nl := parser.getNextXToken(4)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, arg2.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 5
	return nil
}

func (parser *Parser) parseLa() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	virgula1 := parser.getNextXToken(2)
	if virgula1.TokenType != VIRGULA {
		return virgulaMissing(virgula1.Line, virgula1.Column)
	}

	arg2 := parser.getNextXToken(3)
	if arg2.TokenType != LABEL_REF {
		return invalidLabel(arg2)
	}

	nl := parser.getNextXToken(4)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	ins = append(ins, arg1.Value)
	ins = append(ins, arg2.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 5
	return nil
}

func (parser *Parser) parseRand() error {
	ins := Instruction{parser.getCurrentToken().Value}

	arg1 := parser.getNextXToken(1)
	if arg1.TokenType != REGISTER {
		return invalidRegister(arg1)
	}

	nl := parser.getNextXToken(2)
	if nl.TokenType != NEW_LINE {
		return toMuchArgs(parser.getCurrentToken(), 1)
	}

	ins = append(ins, arg1.Value)
	parser.Instructions = append(parser.Instructions, ins)
	parser.Len += 1
	parser.Position += 3
	return nil
}

func (parser *Parser) parseLabelDef() error {
	parser.Labels[parser.getCurrentToken().Value] = parser.Len - 1
	parser.Position += 1
	return nil
}

func (parser *Parser) Parse() []Instruction {
	currentToken := parser.getCurrentToken()
	if (currentToken.TokenType != SECTION) && (currentToken.Value != ".text") {
		println("Seção de text deve ser devinida no começo do arquivo")
		return nil
	}

	validos := map[TokenType]bool{
		SECTION:      true,
		INSTRUCTION:  true,
		LABEL_DEF:    true,
		NEW_LINE:     true,
		STR_LABEL:    true,
		INT8_LABEL:   true,
		INT16_LABEL:  true,
		UINT8_LABEL:  true,
		UINT16_LABEL: true,
	}

	parser.Position += 2
	for parser.Position < len(parser.Tokens) {
		currentToken = parser.getCurrentToken()
		// fmt.Println(currentToken)

		if !validos[currentToken.TokenType] {
			fmt.Printf("Valor inesperado na linha: %d coluna: %d, %s.\n",
				currentToken.Line, currentToken.Column, currentToken.TokenType)
			return nil
		}

		var err error = nil
		if RegistersInstructions[currentToken.Value] {
			err = parser.parseRegisterInstruciton()
		} else if ImmediateInstructions[currentToken.Value] {
			err = parser.parseImediateInstruciton()
		} else if JumpInstructions[currentToken.Value] {
			err = parser.parseJump()
		} else if BranchInstructions[currentToken.Value] {
			err = parser.parseBranch()
		} else if MemoriInstructions[currentToken.Value] {
			err = parser.parseMemorie()
		} else if currentToken.Value == "inc" || currentToken.Value == "dec" {
			err = parser.parseIncDec()
		} else if currentToken.Value == "return" || currentToken.Value == "syscall" {
			err = parser.parseSyscallReturn()
		} else if currentToken.Value == "move" {
			err = parser.parseMove()
		} else if currentToken.Value == "li" {
			err = parser.parseLi()
		} else if currentToken.Value == "la" {
			err = parser.parseLa()
		} else if currentToken.Value == "rand" {
			err = parser.parseRand()
		} else if currentToken.TokenType == LABEL_DEF {
			err = parser.parseLabelDef()
		} else if currentToken.TokenType == NEW_LINE {
			parser.Position += 1
		} else if currentToken.TokenType == SECTION {
			parser.Instructions = append(parser.Instructions, Instruction{".data"})
			parser.Position += 1
		}

		if err != nil {
			return nil
		}
	}

	return parser.Instructions
}
