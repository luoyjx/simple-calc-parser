package simplecaculator

type SimplASTNode struct {
	Type     ASTNodeType
	Text     string
	Children []SimplASTNode
}

func (n *SimplASTNode) AddChild(child SimplASTNode) {
	n.Children = append(n.Children, child)
}

type SimpleCaculator struct {
}

func (c *SimpleCaculator) IntDeclare() error {
	// var node SimplASTNode
	// token := tokens.peek()

	// if token != nil && token.Type == TokenTypeInt { // 匹配 Int
	// 	token = tokens.read() // 消耗掉 int

	// 	if tokens.peek().Type == TokenTypeIndentifier { // 匹配标识符
	// 		token = tokens.read() // 消耗掉标识符

	// 		node = SimplASTNode{
	// 			Type: ASTNodeTypeIntDeclaration,
	// 			Text: token.getText(),
	// 		}

	// 		token = tokens.peek()

	// 		if token != nil && token.Type == TokenTypeAssignment {
	// 			tokens.read() // 消耗掉等号

	// 			var child SimplASTNode = additive(tokens) // 匹配一个表达式

	// 			if child == nil {
	// 				return errors.New("invalid variable initialization, expecting an expression")
	// 			} else {
	// 				node.AddChild(child)
	// 			}
	// 		}
	// 	} else {
	// 		return errors.New("variable name expected")
	// 	}
	// }
	return nil
}
