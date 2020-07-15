package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"

	"calc-parser/visitor/parser"
)

type Visitor struct {
	parser.BaseCalcVisitor
	stack []int
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (v *Visitor) push(i int) {
	v.stack = append(v.stack, i)
}

func (v *Visitor) pop() int {
	if len(v.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := v.stack[len(v.stack)-1]

	// Remove the last element from the stack.
	v.stack = v.stack[:len(v.stack)-1]

	return result
}

func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	node.Accept(v)
	return nil
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	fmt.Println("visis start")
	return v.visitRule(ctx.Expression())
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	i, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		panic(err.Error())
	}

	v.push(i)
	fmt.Println("visit number ", i)
	return nil
}

func (v *Visitor) VisitManulDiv(ctx *parser.ManulDivContext) interface{} {
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))

	right, left := v.pop(), v.pop()

	switch ctx.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		v.push(left * right)
	case parser.CalcParserDIV:
		v.push(left / right)
	default:
		panic("should not happen")
	}

	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))

	right, left := v.pop(), v.pop()

	switch ctx.GetOp().GetTokenType() {
	case parser.CalcLexerADD:
		v.push(left + right)
	case parser.CalcLexerSUB:
		v.push(left - right)
	default:
		panic("should not happen")
	}

	return nil
}

func (v *Visitor) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
	v.visitRule(ctx.Expression())
	fmt.Println("visit parenthesis ", v.stack)
	return nil
}

func calc(input string) int {
	is := antlr.NewInputStream(input)

	// 创建词法解析器
	lexer := parser.NewCalcLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建语法解析器
	p := parser.NewCalcParser(tokens)

	v := NewVisitor()

	p.Start().Accept(v)

	return v.pop()
}

func executor(in string) {
	fmt.Printf("Answer: %d\n", calc(in))
}

func completer(in prompt.Document) []prompt.Suggest {
	var ret []prompt.Suggest
	return ret
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("calc"),
	)
	p.Run()
}
