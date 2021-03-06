package simplecalculator

import (
	"errors"
)

type SimpleParser struct {
}

func (p *SimpleParser) Parse(code string) (node *SimpleASTNode, err error) {
	lexer := SimpleLexer{}
	tokens := lexer.tokenize(code)
	rootNode, err := p.program(&tokens)

	if err != nil {
		return nil, err
	}

	return rootNode, nil
}

// program 将 tokens 解析为 AST Node
func (p *SimpleParser) program(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	node = &SimpleASTNode{
		Type: ASTNodeTypeProgram,
		Text: "pwc",
	}

	for {
		if tokens.peek() == nil {
			break
		}

		child, err := p.intDeclare(tokens)

		if err != nil {
			return nil, err
		}

		if child == nil {
			child, err = p.expressionStatement(tokens)

			if err != nil {
				return nil, err
			}
		}

		if child == nil {
			child, err = p.assignmentStatement(tokens)

			if err != nil {
				return nil, err
			}
		}

		if child != nil {
			node.AddChild(child)
		} else {
			return nil, errors.New("unknown statement")
		}
	}

	child, err := p.additive(tokens)

	if err != nil {
		return nil, err
	}

	if child != nil {
		node.AddChild(child)
	}

	return node, nil
}

// additive 解析加法表达式
func (p *SimpleParser) additive(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	child1, err := p.multiplicative(tokens)

	if err != nil {
		return nil, err
	}

	node = child1

	// 使用尾递归
	if child1 != nil {
		for {
			token := tokens.peek()

			if token != nil && (token.getType() == TokenTypePlus || token.getType() == TokenTypeMinus) {
				token = tokens.read()

				child2, err := p.multiplicative(tokens)

				if err != nil {
					return nil, err
				}

				if child2 == nil {
					return nil, errors.New("invalid additive expression, expecting the right part")
				}

				node = &SimpleASTNode{
					Type: ASTNodeTypeAdditive,
					Text: token.getText(),
				}
				node.AddChild(child1)
				node.AddChild(child2)
				child1 = node
			} else {
				break
			}
		}
	}

	return node, nil
}

// multiplicative 解析乘法表达式
func (p *SimpleParser) multiplicative(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	child1, err := p.primary(tokens)

	if err != nil {
		return nil, err
	}

	node = child1
	token := tokens.peek()
	if child1 != nil && token != nil {

		// 匹配乘除
		if token.getType() == TokenTypeStar || token.getType() == TokenTypeSlash {
			token = tokens.read()

			// 递归匹配乘法表达式
			child2, err := p.multiplicative(tokens)

			if err != nil {
				return nil, err
			}

			if child2 != nil {
				node = &SimpleASTNode{
					Type: ASTNodeTypeMultiplicative,
					Text: token.getText(),
				}
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				return nil, errors.New("invalid multiplicative expression, expecting the right part ")
			}

		}
	}

	return node, nil
}

// primary 解析基础表达式
func (p *SimpleParser) primary(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	token := tokens.peek()

	if token != nil {
		if token.getType() == TokenTypeIntLiteral {
			token = tokens.read()
			node = &SimpleASTNode{
				Type: ASTNodeTypeIntLiteral,
				Text: token.getText(),
			}
		} else if token.getType() == TokenTypeIdentifier {
			token = tokens.read()
			node = &SimpleASTNode{
				Type: ASTNodeTypeIdentifier,
				Text: token.getText(),
			}
		} else if token.getType() == TokenTypeLeftParen {
			tokens.read()
			node, err = p.additive(tokens)

			if err != nil {
				return nil, err
			}

			if node != nil {
				token = tokens.peek()

				if token != nil && token.getType() == TokenTypeRightParen {
					tokens.read()
				} else {
					return nil, errors.New("expecting right parenthesis")
				}
			} else {
				return nil, errors.New("expecting an additive expression inside parenthesis")
			}
		}
	}

	return node, nil
}

func (p *SimpleParser) intDeclare(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	token := tokens.peek()

	if token != nil && token.Type == TokenTypeInt { // 匹配 Int
		token = tokens.read() // 消耗掉 int

		if tokens.peek().Type == TokenTypeIdentifier { // 匹配标识符
			token = tokens.read() // 消耗掉标识符

			node = &SimpleASTNode{
				Type: ASTNodeTypeIntDeclaration,
				Text: token.getText(),
			}

			token = tokens.peek()

			if token != nil && token.Type == TokenTypeAssignment {
				tokens.read() // 消耗掉等号

				child, err := p.additive(tokens) // 匹配一个表达式

				if err != nil {
					return nil, err
				}

				if child == nil {
					return nil, errors.New("invalid variable initialization, expecting an expression")
				} else {
					node.AddChild(child)
				}
			}
		} else {
			return nil, errors.New("variable name expected")
		}

		if node != nil {
			token = tokens.peek()
			if token != nil && token.getType() == TokenTypeSemiColon {
				tokens.read()
			} else {
				return nil, errors.New("invalid statement, expecting semicolon")
			}
		}
	}

	return node, nil
}

// 表达式语句, 表达式后面跟一个分号
func (p *SimpleParser) expressionStatement(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	pos := tokens.getPosition()

	node, err = p.additive(tokens)

	if err != nil {
		return nil, err
	}

	if node != nil {
		token := tokens.peek()

		if token != nil && token.getType() == TokenTypeSemiColon {
			tokens.read()
		} else {
			node = nil
			// 回溯
			tokens.setPosition(pos)
		}
	}

	return node, nil
}

// 赋值语句，如 age = 10 * 2;
func (p *SimpleParser) assignmentStatement(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	token := tokens.peek()

	if token != nil && token.getType() == TokenTypeIdentifier {
		token = tokens.read() // 读取标识符
		node = &SimpleASTNode{
			Type: ASTNodeTypeAssignmentStmt,
			Text: token.getText(),
		}

		// 预读，查看下个是不是等号
		token = tokens.peek()
		if token != nil && token.getType() == TokenTypeAssignment {
			token = tokens.read() // 取出等号
			child, err := p.additive(tokens)

			if err != nil {
				return nil, err
			}

			if child == nil {
				return nil, errors.New("invalid assignment statement, expecting an expression")
			}

			node.AddChild(child)  // 添加子节点
			token = tokens.peek() // 预读看后端是否是分号

			if token != nil && token.getType() == TokenTypeSemiColon {
				tokens.read()
			} else {
				return nil, errors.New("invalid statement, expecting semicolon")
			}
		} else {
			tokens.unread() // 回溯
			node = nil
		}
	}

	return node, nil
}
