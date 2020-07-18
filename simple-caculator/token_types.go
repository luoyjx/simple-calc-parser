package simplecaculator

type DfaState = int
type TokenType = int

const (
	DfaStateInitial DfaState = iota
	DfaStateID
	DfaStateID_Int1
	DfaStateID_Int2
	DfaStateID_Int3
	DfaStateIntLiteral
	DfaStateGT
	DfaStateGE
	DfaStatePlus
	DfaStateMinus
	DfaStateStar
	DfaStateSlash
	DfaStateSemiColon
	DfaStateLeftParen
	DfaStateRightParen
	DfaStateAssignment
)

const (
	TokenTypeIdentifier = iota
	TokenTypeIntLiteral
	TokenTypeGT
	TokenTypeGE
	TokenTypePlus
	TokenTypeMinus
	TokenTypeStar
	TokenTypeSlash
	TokenTypeSemiColon
	TokenTypeLeftParen
	TokenTypeRightParen
	TokenTypeAssignment
	TokenTypeInt
	TokenTypeMultiplicative
)
