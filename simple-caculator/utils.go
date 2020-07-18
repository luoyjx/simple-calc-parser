package simplecalculator

import (
	"fmt"
	"unicode"
)

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

func DumpAST(node *SimpleASTNode, indent string) {
	fmt.Println(fmt.Sprintf("%s %s %s", indent, node.GetTypeStr(), node.GetText()))
	for _, child := range node.Children {
		DumpAST(child, fmt.Sprintf("%s\t", indent))
	}
}
