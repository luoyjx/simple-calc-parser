package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"calc-parser/parser"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty")
	}

	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

// 乘除法结束时
func (l *calcListener) ExitManulDiv(c *parser.ManulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexcepted op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexcepted op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err)
	}

	l.push(i)
}

func calc(input string) int {
	// 初始化输入
	is := antlr.NewInputStream("1 + 2 * 3")

	// 创建词法分析器
	lexer := parser.NewCalcLexer(is)

	// 构建 Token Stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	var listener calcListener
	// 创建语法解析器
	p := parser.NewCalcParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()
}

func main() {
	result := calc("1 + 2 * 3")

	fmt.Println("caculate result : ", result)
}
