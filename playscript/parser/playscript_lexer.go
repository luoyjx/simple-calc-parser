// Code generated from PlayScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 192,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 84, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 121,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 127, 10, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 7, 24, 146, 10, 24, 12, 24, 14, 24, 149, 11,
	24, 3, 25, 3, 25, 5, 25, 153, 10, 25, 3, 26, 3, 26, 7, 26, 157, 10, 26,
	12, 26, 14, 26, 160, 11, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 30, 3, 30, 7, 30, 170, 10, 30, 12, 30, 14, 30, 173, 11, 30, 3, 30, 3,
	30, 3, 31, 6, 31, 178, 10, 31, 13, 31, 14, 31, 179, 3, 31, 3, 31, 3, 32,
	3, 32, 5, 32, 186, 10, 32, 3, 32, 5, 32, 189, 10, 32, 3, 32, 3, 32, 3,
	171, 2, 33, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 2, 53, 2, 55, 2, 57,
	2, 59, 27, 61, 28, 63, 29, 3, 2, 9, 4, 2, 45, 45, 47, 47, 4, 2, 44, 44,
	49, 49, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 3, 2, 50, 59, 3, 2, 51, 59, 4, 2, 11, 11, 34, 34, 2, 201, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 3, 65, 3,
	2, 2, 2, 5, 67, 3, 2, 2, 2, 7, 70, 3, 2, 2, 2, 9, 73, 3, 2, 2, 2, 11, 83,
	3, 2, 2, 2, 13, 85, 3, 2, 2, 2, 15, 87, 3, 2, 2, 2, 17, 89, 3, 2, 2, 2,
	19, 91, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 95, 3, 2, 2, 2, 25, 97, 3,
	2, 2, 2, 27, 99, 3, 2, 2, 2, 29, 101, 3, 2, 2, 2, 31, 103, 3, 2, 2, 2,
	33, 105, 3, 2, 2, 2, 35, 107, 3, 2, 2, 2, 37, 120, 3, 2, 2, 2, 39, 126,
	3, 2, 2, 2, 41, 128, 3, 2, 2, 2, 43, 133, 3, 2, 2, 2, 45, 139, 3, 2, 2,
	2, 47, 143, 3, 2, 2, 2, 49, 152, 3, 2, 2, 2, 51, 154, 3, 2, 2, 2, 53, 161,
	3, 2, 2, 2, 55, 163, 3, 2, 2, 2, 57, 165, 3, 2, 2, 2, 59, 167, 3, 2, 2,
	2, 61, 177, 3, 2, 2, 2, 63, 188, 3, 2, 2, 2, 65, 66, 7, 63, 2, 2, 66, 4,
	3, 2, 2, 2, 67, 68, 7, 126, 2, 2, 68, 69, 7, 126, 2, 2, 69, 6, 3, 2, 2,
	2, 70, 71, 7, 40, 2, 2, 71, 72, 7, 40, 2, 2, 72, 8, 3, 2, 2, 2, 73, 74,
	7, 35, 2, 2, 74, 10, 3, 2, 2, 2, 75, 76, 7, 63, 2, 2, 76, 84, 7, 63, 2,
	2, 77, 84, 7, 64, 2, 2, 78, 79, 7, 64, 2, 2, 79, 84, 7, 63, 2, 2, 80, 84,
	7, 62, 2, 2, 81, 82, 7, 62, 2, 2, 82, 84, 7, 63, 2, 2, 83, 75, 3, 2, 2,
	2, 83, 77, 3, 2, 2, 2, 83, 78, 3, 2, 2, 2, 83, 80, 3, 2, 2, 2, 83, 81,
	3, 2, 2, 2, 84, 12, 3, 2, 2, 2, 85, 86, 9, 2, 2, 2, 86, 14, 3, 2, 2, 2,
	87, 88, 9, 3, 2, 2, 88, 16, 3, 2, 2, 2, 89, 90, 7, 37, 2, 2, 90, 18, 3,
	2, 2, 2, 91, 92, 7, 61, 2, 2, 92, 20, 3, 2, 2, 2, 93, 94, 7, 48, 2, 2,
	94, 22, 3, 2, 2, 2, 95, 96, 7, 46, 2, 2, 96, 24, 3, 2, 2, 2, 97, 98, 7,
	93, 2, 2, 98, 26, 3, 2, 2, 2, 99, 100, 7, 95, 2, 2, 100, 28, 3, 2, 2, 2,
	101, 102, 7, 125, 2, 2, 102, 30, 3, 2, 2, 2, 103, 104, 7, 127, 2, 2, 104,
	32, 3, 2, 2, 2, 105, 106, 7, 42, 2, 2, 106, 34, 3, 2, 2, 2, 107, 108, 7,
	43, 2, 2, 108, 36, 3, 2, 2, 2, 109, 110, 7, 101, 2, 2, 110, 111, 7, 106,
	2, 2, 111, 112, 7, 99, 2, 2, 112, 121, 7, 116, 2, 2, 113, 114, 7, 107,
	2, 2, 114, 115, 7, 112, 2, 2, 115, 121, 7, 118, 2, 2, 116, 117, 7, 100,
	2, 2, 117, 118, 7, 113, 2, 2, 118, 119, 7, 113, 2, 2, 119, 121, 7, 110,
	2, 2, 120, 109, 3, 2, 2, 2, 120, 113, 3, 2, 2, 2, 120, 116, 3, 2, 2, 2,
	121, 38, 3, 2, 2, 2, 122, 123, 7, 107, 2, 2, 123, 127, 7, 104, 2, 2, 124,
	125, 7, 22916, 2, 2, 125, 127, 7, 26526, 2, 2, 126, 122, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 40, 3, 2, 2, 2, 128, 129, 7, 103, 2, 2, 129, 130,
	7, 110, 2, 2, 130, 131, 7, 117, 2, 2, 131, 132, 7, 103, 2, 2, 132, 42,
	3, 2, 2, 2, 133, 134, 7, 121, 2, 2, 134, 135, 7, 106, 2, 2, 135, 136, 7,
	107, 2, 2, 136, 137, 7, 110, 2, 2, 137, 138, 7, 103, 2, 2, 138, 44, 3,
	2, 2, 2, 139, 140, 7, 104, 2, 2, 140, 141, 7, 113, 2, 2, 141, 142, 7, 116,
	2, 2, 142, 46, 3, 2, 2, 2, 143, 147, 9, 4, 2, 2, 144, 146, 9, 5, 2, 2,
	145, 144, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 48, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 153, 5,
	51, 26, 2, 151, 153, 5, 59, 30, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2,
	2, 2, 153, 50, 3, 2, 2, 2, 154, 158, 5, 57, 29, 2, 155, 157, 5, 55, 28,
	2, 156, 155, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158,
	159, 3, 2, 2, 2, 159, 52, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 162, 9,
	4, 2, 2, 162, 54, 3, 2, 2, 2, 163, 164, 9, 6, 2, 2, 164, 56, 3, 2, 2, 2,
	165, 166, 9, 7, 2, 2, 166, 58, 3, 2, 2, 2, 167, 171, 7, 36, 2, 2, 168,
	170, 11, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171, 172,
	3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 171, 3, 2,
	2, 2, 174, 175, 7, 36, 2, 2, 175, 60, 3, 2, 2, 2, 176, 178, 9, 8, 2, 2,
	177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179,
	180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 8, 31, 2, 2, 182, 62,
	3, 2, 2, 2, 183, 185, 7, 15, 2, 2, 184, 186, 7, 12, 2, 2, 185, 184, 3,
	2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 189, 7, 12, 2,
	2, 188, 183, 3, 2, 2, 2, 188, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190,
	191, 8, 32, 2, 2, 191, 64, 3, 2, 2, 2, 14, 2, 83, 120, 126, 145, 147, 152,
	158, 171, 179, 185, 188, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'||'", "'&&'", "'!'", "", "", "", "'#'", "';'", "'.'", "','",
	"'['", "']'", "'{'", "'}'", "'('", "')'", "", "", "'else'", "'while'",
	"'for'",
}

var lexerSymbolicNames = []string{
	"", "Assignment", "Or", "And", "Not", "Relational", "Add", "Mul", "Sharp",
	"SemiColon", "Dot", "Comm", "LeftBracket", "RightBracket", "LeftBrace",
	"RightBrace", "LeftParen", "RightParen", "TypeSpecifier", "If", "Else",
	"While", "For", "Identifier", "Constant", "CharacterConstant", "Whitespace",
	"Newline",
}

var lexerRuleNames = []string{
	"Assignment", "Or", "And", "Not", "Relational", "Add", "Mul", "Sharp",
	"SemiColon", "Dot", "Comm", "LeftBracket", "RightBracket", "LeftBrace",
	"RightBrace", "LeftParen", "RightParen", "TypeSpecifier", "If", "Else",
	"While", "For", "Identifier", "Constant", "IntegerConstant", "Nondigit",
	"Digit", "NonzeroDigit", "CharacterConstant", "Whitespace", "Newline",
}

type PlayScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPlayScriptLexer(input antlr.CharStream) *PlayScriptLexer {

	l := new(PlayScriptLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PlayScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PlayScriptLexer tokens.
const (
	PlayScriptLexerAssignment        = 1
	PlayScriptLexerOr                = 2
	PlayScriptLexerAnd               = 3
	PlayScriptLexerNot               = 4
	PlayScriptLexerRelational        = 5
	PlayScriptLexerAdd               = 6
	PlayScriptLexerMul               = 7
	PlayScriptLexerSharp             = 8
	PlayScriptLexerSemiColon         = 9
	PlayScriptLexerDot               = 10
	PlayScriptLexerComm              = 11
	PlayScriptLexerLeftBracket       = 12
	PlayScriptLexerRightBracket      = 13
	PlayScriptLexerLeftBrace         = 14
	PlayScriptLexerRightBrace        = 15
	PlayScriptLexerLeftParen         = 16
	PlayScriptLexerRightParen        = 17
	PlayScriptLexerTypeSpecifier     = 18
	PlayScriptLexerIf                = 19
	PlayScriptLexerElse              = 20
	PlayScriptLexerWhile             = 21
	PlayScriptLexerFor               = 22
	PlayScriptLexerIdentifier        = 23
	PlayScriptLexerConstant          = 24
	PlayScriptLexerCharacterConstant = 25
	PlayScriptLexerWhitespace        = 26
	PlayScriptLexerNewline           = 27
)