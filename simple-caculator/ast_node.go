package simplecalculator

// ASTNodeType alisa
type ASTNodeType = int

// AST 节点类型
const (
	ASTNodeTypeIntDeclaration ASTNodeType = iota // 整型变量声明
	ASTNodeTypeProgram
	ASTNodeTypeAdditive       // 加法表达式
	ASTNodeTypeMultiplicative // 乘法表达式
	ASTNodeTypeIntLiteral     // 整型字面量
	ASTNodeTypeIdentifier     // 标识符

)

var ASTNodeTypeStr = []string{
	"IntDeclaration",
	"Program",
	"Additive",
	"Multiplicative",
	"IntLiteral",
	"Identifier",
}

type ASTNode interface {
	AddChild() ASTNode
	GetType() ASTNodeType
	GetText() string
}

type SimpleASTNode struct {
	Type     ASTNodeType
	Text     string
	Children []*SimpleASTNode
}

// AddChild adds a child SimpleASTNode
func (n *SimpleASTNode) AddChild(child *SimpleASTNode) {
	n.Children = append(n.Children, child)
}

func (n *SimpleASTNode) GetType() ASTNodeType {
	return n.Type
}

func (n *SimpleASTNode) GetTypeStr() string {
	return ASTNodeTypeStr[int(n.Type)]
}

func (n *SimpleASTNode) GetText() string {
	return n.Text
}