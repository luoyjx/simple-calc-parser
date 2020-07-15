package simplecaculator

import (
	"fmt"
	"unicode"
)

type SimpleToken struct {
	Type TokenType
	Text string
}

type SimpleLexer struct {
	Tokens    []SimpleToken // 解析完成的 Token
	Token     SimpleToken   // 当前正在解析的 Token
	TokenText []rune        // 未解析完的 Token 的临时文本
}

func (s *SimpleLexer) tokenize(script string) SimpleTokenReader {
	s.Tokens = make([]SimpleToken, 0)
	s.Token = SimpleToken{}

	chars := []rune(script)

	state := DfaStateInitial
	for _, ch := range chars {
		switch state {
		case DfaStateInitial:
			state = s.initToken(ch)
		case DfaStateID:
			if isAlpha(ch) || isDigit(ch) {
				s.TokenText = append(s.TokenText, ch) // 保持状态
			} else {
				state = s.initToken(ch) // token 退出，保存 Token
			}
		case DfaStateGT:
			if ch == '=' {
				s.Token.Type = TokenTypeGE
				state = DfaStateGE
				s.TokenText = append(s.TokenText, ch)
			} else {
				state = s.initToken(ch)
			}
		case DfaStateGE, DfaStateAssignment, DfaStatePlus, DfaStateMinus, DfaStateStar,
			DfaStateSlash, DfaStateSemiColon, DfaStateLeftParen, DfaStateRightParen:
			state = s.initToken(ch)
		case DfaStateIntLiteral:
			if isDigit(ch) {
				s.TokenText = append(s.TokenText, ch) // 继续保持数字字面量
			} else {
				state = s.initToken(ch)
			}
		case DfaStateID_Int1:
			if ch == 'n' {
				state = DfaStateID_Int2
				s.TokenText = append(s.TokenText, ch)
			} else if isDigit(ch) || isAlpha(ch) {
				state = DfaStateID
				s.TokenText = append(s.TokenText, ch)
			} else {
				state = s.initToken(ch)
			}
		case DfaStateID_Int2:
			if ch == 't' {
				state = DfaStateID_Int3
				s.TokenText = append(s.TokenText, ch)
			} else if isDigit(ch) || isDigit(ch) {
				state = DfaStateID
				s.TokenText = append(s.TokenText, ch)
			} else {
				state = s.initToken(ch)
			}
		case DfaStateID_Int3:
			if isBlank(ch) {
				s.Token.Type = TokenTypeInt
				state = s.initToken(ch)
			} else {
				state = DfaStateID
				s.TokenText = append(s.TokenText, ch)
			}
		}
	}

	if len(s.TokenText) > 0 {
		s.initToken(rune(-1))
	}

	return SimpleTokenReader{
		Tokens:   s.Tokens,
		Position: 0,
	}
}

func (s *SimpleLexer) initToken(ch rune) DfaState {
	if len(s.TokenText) > 0 {
		s.Token.Text = string(s.TokenText)
		s.Tokens = append(s.Tokens, s.Token)

		s.TokenText = make([]rune, 0)
		s.Token = SimpleToken{}
	}

	newState := DfaStateInitial
	if isAlpha(ch) {
		if ch == 'i' {
			newState = DfaStateID_Int1
			s.TokenText = append(s.TokenText, ch)
		} else {
			newState = DfaStateID // 进入 ID 状态
			s.TokenText = append(s.TokenText, ch)
		}
	} else if isDigit(ch) { // 第一个字符是数字
		newState = DfaStateIntLiteral
		s.Token.Type = TokenTypeIntLiteral
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '>' { // 第一个字符是 >
		newState = DfaStateGT
		s.Token.Type = TokenTypeGT
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '+' {
		newState = DfaStatePlus
		s.Token.Type = TokenTypePlus
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '-' {
		newState = DfaStateMinus
		s.Token.Type = TokenTypeMinus
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '*' {
		newState = DfaStateStar
		s.Token.Type = TokenTypeStar
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '/' {
		newState = DfaStateSlash
		s.Token.Type = TokenTypeSlash
	} else if ch == ';' {
		newState = DfaStateSemiColon
		s.Token.Type = TokenTypeSemiColon
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '(' {
		newState = DfaStateLeftParen
		s.Token.Type = TokenTypeLeftParen
		s.TokenText = append(s.TokenText, ch)
	} else if ch == ')' {
		newState = DfaStateRightParen
		s.Token.Type = TokenTypeRightParen
		s.TokenText = append(s.TokenText, ch)
	} else if ch == '=' {
		newState = DfaStateAssignment
		s.Token.Type = TokenTypeAssignment
		s.TokenText = append(s.TokenText, ch)
	} else {
		newState = DfaStateInitial
	}

	return newState
}

type SimpleTokenReader struct {
	Tokens   []SimpleToken
	Position int
}

func (r *SimpleTokenReader) read() *SimpleToken {
	if len(r.Tokens) > r.Position {
		token := r.Tokens[r.Position]
		r.Position++
		return &token
	}

	return nil
}

func isAlpha(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func Dump(reader SimpleTokenReader) {
	for {
		token := reader.read()

		if token == nil {
			break
		}

		fmt.Println(token.Text, "\t\t", token.Type)
	}
	fmt.Println("===========> Reading Tokens done\n ")
}
