package simplecalculator

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculator(t *testing.T) {
	parser := SimpleCalculator{}

	Convey("should parse int declaration ok", t, func() {
		code := "int a = b+3;"
		fmt.Println("解析变量声明语句： ", code)

		lexer := SimpleLexer{}
		tokens := lexer.tokenize(code)

		node, err := parser.IntDeclare(&tokens)

		So(err, ShouldBeNil)
		So(node, ShouldNotBeNil)

		DumpAST(node, "")
	})

	Convey("should evaluate script ok", t, func() {
		code := "2+3*5"
		fmt.Println("\n计算：", code, ",看上去一切正常")
		parser.Evaluate(code)

		code = "2+"
		fmt.Println("\n计算：", code, ",应该有语法错误。")
		parser.Evaluate(code)

		code = "2+3+4"
		fmt.Println("\n计算：", code, ", 结合性出现错误。")
		parser.Evaluate(code)
	})
}
