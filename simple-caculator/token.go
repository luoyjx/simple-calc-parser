package simplecalculator

type Token interface {
	getType() TokenType
	getText() string
}

type SimpleToken struct {
	Type TokenType
	Text string
}

func (st SimpleToken) getType() TokenType {
	return st.Type
}

func (st SimpleToken) getText() string {
	return st.Text
}
