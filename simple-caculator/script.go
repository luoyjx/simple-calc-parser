package simplecalculator

import (
	"errors"
	"fmt"
	"strconv"
)

type SimpleScript struct {
	variables map[string]int
	verbose   bool
}

func (s *SimpleScript) Evaluate(node *SimpleASTNode, indent string) (result int, err error) {
	fmt.Println(indent, "Calculating: ", node.GetTypeStr())

	switch node.GetType() {
	case ASTNodeTypeProgram:
		for _, child := range node.Children {
			result, err = s.Evaluate(child, indent+"\t")

			if err != nil {
				return 0, err
			}
		}

	case ASTNodeTypeAdditive:
		child1 := node.Children[0]
		value1, err := s.Evaluate(child1, indent+"\t")

		if err != nil {
			return 0, err
		}

		child2 := node.Children[1]
		value2, err := s.Evaluate(child2, indent+"\t")

		if err != nil {
			return 0, err
		}

		if node.GetText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}

	case ASTNodeTypeMultiplicative:
		child1 := node.Children[0]
		value1, err := s.Evaluate(child1, indent+"\t")

		if err != nil {
			return 0, err
		}

		child2 := node.Children[1]
		value2, err := s.Evaluate(child2, indent+"\t")

		if err != nil {
			return 0, err
		}

		if node.GetText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}

	case ASTNodeTypeIntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	case ASTNodeTypeIdentifier:
		varName := node.GetText()

		if val, exists := s.variables[varName]; exists {

			// 判断变量是否被赋值过
			if val != 0 {
				result = val
			} else {
				return 0, errors.New(fmt.Sprintf("variable %s has not been set any value", varName))
			}
		} else {
			// 未被定义过
			return 0, errors.New(fmt.Sprintf("unknown variable %s", varName))
		}
	case ASTNodeTypeAssignmentStmt:
		varName := node.GetText()

		if _, exists := s.variables[varName]; !exists {
			return 0, errors.New(fmt.Sprintf("unknown variable %s", varName))
		}

		fallthrough
	case ASTNodeTypeIntDeclaration:
		varName := node.GetText()

		var varVal int

		if len(node.Children) > 0 {
			child := node.Children[0]
			result, err = s.Evaluate(child, indent+"\t")

			if err != nil {
				return 0, err
			}

			varVal = result
		}

		s.variables[varName] = varVal
	}

	fmt.Println(indent, "Result: ", result)

	return result, nil
}
