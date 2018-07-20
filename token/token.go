package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL       = "ILLEGAL"
	EOF           = "EOF"
	IDENTIFIER    = "IDENTIFIER"
	INTEGER       = "INTEGER"
	MOVE          = "MOV"
	ADD           = "ADD"
	MULTIPLY      = "MUL"
	JUMP          = "JMP"
	JUMP_NOT_ZERO = "JNZ"
	RETURN        = "RET"
	INTERRUPT     = "INT"
)

var keywords = map[string]TokenType{
	"mov": MOVE,
	"MOV": MOVE,

	"add": ADD,
	"ADD": ADD,

	"mul": MULTIPLY,
	"MUL": MULTIPLY,

	"jmp": JUMP,
	"JMP": JUMP,

	"jnz": JUMP_NOT_ZERO,
	"JNZ": JUMP_NOT_ZERO,

	"ret": RETURN,
	"RET": RETURN,

	"int": INTERRUPT,
	"INT": INTERRUPT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENTIFIER
}
