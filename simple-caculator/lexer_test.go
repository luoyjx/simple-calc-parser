package simplecaculator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLexer(t *testing.T) {
	lexer := SimpleLexer{}

	Convey("should tokenize int declaration ok", t, func() {
		reader := lexer.tokenize("int age = 45;")

		token := reader.read()
		So(token.Type, ShouldEqual, TokenTypeInt)
		So(token.Text, ShouldEqual, "int")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeIdentifier)
		So(token.Text, ShouldEqual, "age")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeAssignment)
		So(token.Text, ShouldEqual, "=")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeIntLiteral)
		So(token.Text, ShouldEqual, "45")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeSemiColon)
		So(token.Text, ShouldEqual, ";")
	})

	Convey("should tokenize comparison ok", t, func() {
		reader := lexer.tokenize("age >= 45")

		token := reader.read()
		So(token.Type, ShouldEqual, TokenTypeIdentifier)
		So(token.Text, ShouldEqual, "age")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeGE)
		So(token.Text, ShouldEqual, ">=")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeIntLiteral)
		So(token.Text, ShouldEqual, "45")
	})

	Convey("should tokenize calculate expression ok", t, func() {
		reader := lexer.tokenize("2+3*5")

		token := reader.read()
		So(token.Type, ShouldEqual, TokenTypeIntLiteral)
		So(token.Text, ShouldEqual, "2")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypePlus)
		So(token.Text, ShouldEqual, "+")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeIntLiteral)
		So(token.Text, ShouldEqual, "3")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeStar)
		So(token.Text, ShouldEqual, "*")

		token = reader.read()
		So(token.Type, ShouldEqual, TokenTypeIntLiteral)
		So(token.Text, ShouldEqual, "5")
	})
}
