package simplecalculator

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	parser := SimpleParser{}

	Convey("should parse int declaration expression ok ", t, func() {
		script := "int age = 45+2; age= 20; age+10*2;"
		fmt.Println("解析：", script)
		tree, err := parser.Parse(script)
		So(err, ShouldBeNil)
		DumpAST(tree, "")
	})

	Convey("should parse invalid additive script fail", t, func() {
		script := "2+3+;"
		fmt.Println("解析：", script)
		_, err := parser.Parse(script)
		fmt.Println(err.Error())
		So(err, ShouldNotBeNil)
	})

	Convey("should parse invalid multiplicative script fail", t, func() {
		script := "2+3*;"
		fmt.Println("解析：", script)
		_, err := parser.Parse(script)
		fmt.Println(err.Error())
		So(err, ShouldNotBeNil)
	})
}
