// Code generated from PlayScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePlayScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlayScriptVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitBlockItemList(ctx *BlockItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitBlockItem(ctx *BlockItemContext) interface{} {
	return v.VisitChildren(ctx)
}
