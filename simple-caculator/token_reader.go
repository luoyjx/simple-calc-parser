package simplecalculator

type TokenReader interface {
	read() Token
	peek() Token
	unread()
	getPosition() int
	setPosition(pos int)
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

func (r *SimpleTokenReader) peek() *SimpleToken {
	if len(r.Tokens) > r.Position {
		token := r.Tokens[r.Position]
		return &token
	}

	return nil
}

func (r *SimpleTokenReader) unread() {
	if r.Position > 0 {
		r.Position--
	}
}

func (r *SimpleTokenReader) getPosition() int {
	return r.Position
}

func (r *SimpleTokenReader) setPosition(pos int) {
	if r.Position >= 0 && r.Position < len(r.Tokens) {
		r.Position = pos
	}
}