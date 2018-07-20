package lexer

import "github.com/marahin/yaaibig/token"

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int // current reading position
	character    byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	default:
		if isLetter(lexer.character) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		} else if isDigit(lexer.character) {
			tok.Type = token.INTEGER
			tok.Literal = lexer.readNumber()

			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.character)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lexer.readChar()

	return tok
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.character) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isLetter(character byte) bool {
	is_between_a_to_z := 'a' <= character && character <= 'z'
	is_between_A_to_Z := 'A' <= character && character <= 'Z'
	is_special_case := false // change if special cases appear

	return is_between_a_to_z || is_between_A_to_Z || is_special_case
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.character) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) skipWhitespace() {
	switch lexer.character {
	case
		' ',
		'\t',
		'\n',
		'\r':
		lexer.readChar()
	}
}
