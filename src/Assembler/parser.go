package assembler

import "fmt"

type Instruction string

const (
	ADD     Instruction = "add"
	ADDI    Instruction = "addi"
	ADDU    Instruction = "addu"
	ADDUI   Instruction = "addui"
	SUB     Instruction = "sub"
	SUBI    Instruction = "subi"
	SUBU    Instruction = "subu"
	SUBUI   Instruction = "subui"
	MULT    Instruction = "mult"
	MULTI   Instruction = "multi"
	DIV     Instruction = "div"
	DIVI    Instruction = "divi"
	OR      Instruction = "or"
	ORI     Instruction = "ori"
	AND     Instruction = "and"
	ANDI    Instruction = "andi"
	SLL     Instruction = "sll"
	SRL     Instruction = "srl"
	SLT     Instruction = "slt"
	SLTI    Instruction = "slti"
	LI      Instruction = "li"
	LA      Instruction = "la"
	MOVE    Instruction = "move"
	J       Instruction = "j"
	JR      Instruction = "jr"
	JAL     Instruction = "jal"
	BEQ     Instruction = "beq"
	BNE     Instruction = "bne"
	BGT     Instruction = "bgt"
	BGE     Instruction = "bge"
	BLT     Instruction = "blt"
	BLE     Instruction = "ble"
	LW      Instruction = "lw"
	LB      Instruction = "lb"
	SW      Instruction = "sw"
	SB      Instruction = "sb"
	LV      Instruction = "lv"
	SV      Instruction = "sv"
	LRW     Instruction = "lrw"
	LRB     Instruction = "lrb"
	INC     Instruction = "inc"
	DEC     Instruction = "dec"
	SYSCALL Instruction = "syscall"
	RETURN  Instruction = "return"
	RAND    Instruction = "rand"
)

var RegistersInstructions = []Instruction{
	ADD, ADDU, SUB, SUBU, MULT, DIV, OR, AND, SLT,
}

var ImmediateInstructions = []Instruction{
	ADDI, ADDUI, SUBI, SUBUI, MULTI, DIVI, ORI, ANDI, SLTI, SLL, SRL,
}

var BranchInstructions = []Instruction{
	BEQ, BNE, BGT, BGE, BLT, BLE,
}

var MemoriInstructions = []Instruction{
	LW, LB, SW, SB, LV, SV, LRW, LRB,
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

type Parser struct {
	Tokens       []Token
	Instructions []string
	Position     int
	Len          int
}

func newParser(tokens []Token) *Parser {
	p := Parser{tokens, []string{}, 0, 0}
	return &p
}

// add, addu, subu, sub, mult, div, or, and, slt
func (parser *Parser) parseRegisterInstruciton() error {
	return nil
}

// addi, addui, subi, subui, multi, divi, ori, andi, slti, sll, srl
func (parser *Parser) parseImediateInstruciton() error {
	return nil
}

// j, jr, jal
func (parser *Parser) parseJump() error {
	return nil
}

// beq, bne, bgt, bge, blt, ble
func (parser *Parser) parseBranch() error {
	return nil
}

// lw, lb, sw, sb, lv, sv, lrw, lrb
func (parser *Parser) parseMemorie() error {
	return nil
}

// inc, dec
func (parser *Parser) parseIncDec() error {
	return nil
}

// Syscall, Return
func (parser *Parser) parseSyscallReturn() error {
	return nil
}

func (parser *Parser) parseMove() error {
	return nil
}

func (parser *Parser) parseLi() error {
	return nil
}

func (parser *Parser) parseLa() error {
	return nil
}

func (parser *Parser) parseRand() error {
	return nil
}

func (parser *Parser) parseLabelRef() error {
	return nil
}

func (parser *Parser) parseLabelDef() error {
	return nil
}

func (parser *Parser) parseRegister() error {
	return nil
}

func (parser *Parser) Parse() []string {
	fmt.Println("Parser")

	return parser.Instructions
}
