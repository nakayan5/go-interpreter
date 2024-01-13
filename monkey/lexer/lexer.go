// 字句解析機（レキサー）
package lexer

import "go-interpreter/monkey/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字を指し示す)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

// 入力を受け取り、それをトークンに分割する。
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 現在の文字を読み込み、その文字をchに設定します。
// さらに、positionとreadPositionとchを更新します。
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// readXChar()との違いは、現在の文字を読み進めずに、次の文字を返す点です。
// l.positionとl.readPositionは更新されません。
// peek (覗き見)
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		// 終端に達した場合は0を返す。
		return 0
	} else {
		// まだ終端に達していない場合は、次の文字を返す。
		return l.input[l.readPosition]
	}
}

// 現在の文字を検査し、その文字に応じてトークンを返します。
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	// 特定の記号の場合は、その記号に対応するトークンを返します。
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	// end of file
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	// それ以外の場合は、識別子か数字かを判定します。
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// 現在の文字が英字である限り、文字を読み進めます。
func (l *Lexer) readIdentifier() string {
	// positionを固定して、下のreadChar()でl.positionが進む。
	// その結果、l.input[position:l.position]で識別子を取得できる。
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 特定の文字かどうかを判定します。
// [a-zA-Z_]のいずれかであればtrueを返します。
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Token構造体を生成します。
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 空白文字をスキップ。
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 現在の文字が数字である限り、文字を読み進めます。
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 文字が数字かどうかを判定します。
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
