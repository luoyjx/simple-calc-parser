package simplecaculator

import (
	"errors"
	"fmt"
	"strconv"
)

type SimpleCalculator struct {
}

func (c *SimpleCalculator) Evaluate(code string) {
	tree, err := c.parse(code)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	DumpAST(tree, "")
	c.evaluate(tree, "")
}

func (c *SimpleCalculator) parse(code string) (node *SimpleASTNode, err error) {
	lexer := SimpleLexer{}
	tokens := lexer.tokenize(code)
	rootNode, err := c.program(&tokens)

	if err != nil {
		return nil, err
	}

	return rootNode, nil
}

// program 将 tokens 解析为 AST Node
func (c *SimpleCalculator) program(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	node = &SimpleASTNode{
		Type: ASTNodeTypeProgram,
		Text: "Calculator",
	}

	child, err := c.additive(tokens)

	if err != nil {
		return nil, err
	}

	if child != nil {
		node.AddChild(child)
	}

	return node, nil
}

func (c *SimpleCalculator) evaluate(node *SimpleASTNode, indent string) int {
	var result int
	fmt.Println(indent, "Calculating: ", node.GetTypeStr())

	switch node.GetType() {
	case ASTNodeTypeProgram:
		for _, child := range node.Children {
			result = c.evaluate(child, indent + "\t")
		}

	case ASTNodeTypeAdditive:
		child1 := node.Children[0]
		value1 := c.evaluate(child1, indent + "\t")

		child2 := node.Children[1]
		value2 := c.evaluate(child2, indent + "\t")

		if node.GetText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}

	case ASTNodeTypeMultiplicative:
		child1 := node.Children[0]
		value1 := c.evaluate(child1, indent + "\t")

		child2 := node.Children[1]
		value2 := c.evaluate(child2, indent + "\t")

		if node.GetText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}

	case ASTNodeTypeIntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	}

	fmt.Println(indent, "Result: ", result)

	return result
}

// additive 解析加法表达式
func (c *SimpleCalculator) additive(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	child1, err := c.multiplicative(tokens)

	if err != nil {
		return nil, err
	}

	node = child1

	// 使用尾递归
	if child1 != nil {
		for {
			token := tokens.peek()

			if token != nil && (token.getType() == TokenTypePlus  || token.getType() == TokenTypeMinus) {
				token = tokens.read()

				child2, err := c.multiplicative(tokens)

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

	// 使用左递归的方式自顶向下解析的方法，此方法解析加法存在左结合性不对的问题
	//if child1 != nil && token != nil {
	//	if token.getType() == TokenTypePlus || token.getType() == TokenTypeMinus {
	//		// 消耗 token
	//		token = tokens.read()
	//
	//		// 递归下降
	//		child2, err := c.additive(tokens)
	//
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		if child2 != nil {
	//			node = &SimpleASTNode{
	//				Type: ASTNodeTypeAdditive,
	//				Text: token.getText(),
	//			}
	//			node.AddChild(child1)
	//			node.AddChild(child2)
	//		} else {
	//			return nil, errors.New("invalid additive expression, expecting the right part")
	//		}
	//	}
	//}
	//
	//return node, nil
}

// multiplicative 解析乘法表达式
func (c *SimpleCalculator) multiplicative(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
	child1, err := c.primary(tokens)

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
			child2, err := c.multiplicative(tokens)

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
func (c *SimpleCalculator) primary(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
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
			node, err = c.additive(tokens)

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

func (c *SimpleCalculator) IntDeclare(tokens *SimpleTokenReader) (node *SimpleASTNode, err error) {
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

				child, err := c.additive(tokens) // 匹配一个表达式

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
