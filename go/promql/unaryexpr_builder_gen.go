// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*UnaryExprBuilder)(nil)

// Represents a PromQL expression.
type UnaryExprBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewUnaryExprBuilder() *UnaryExprBuilder {
	resource := NewExpr()
	builder := &UnaryExprBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	if builder.internal.UnaryExpr == nil {
		builder.internal.UnaryExpr = NewUnaryExpr()
	}
	builder.internal.UnaryExpr.Type = "unaryExpr"

	return builder
}

// Negation unary operator.
func Neg(expr cog.Builder[Expr]) *UnaryExprBuilder {
	builder := NewUnaryExprBuilder()
	builder.Op(UnaryOpMinus)
	builder.Expr(expr)

	return builder
}

// Identity unary operator.
func Id(expr cog.Builder[Expr]) *UnaryExprBuilder {
	builder := NewUnaryExprBuilder()
	builder.Op(UnaryOpPlus)
	builder.Expr(expr)

	return builder
}

func (builder *UnaryExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	return *builder.internal, nil
}

func (builder UnaryExprBuilder) String() string {
	return builder.internal.String()
}

func (builder *UnaryExprBuilder) Op(op UnaryOp) *UnaryExprBuilder {
	if builder.internal.UnaryExpr == nil {
		builder.internal.UnaryExpr = NewUnaryExpr()
	}
	builder.internal.UnaryExpr.Op = op

	return builder
}

func (builder *UnaryExprBuilder) Expr(expr cog.Builder[Expr]) *UnaryExprBuilder {
	if builder.internal.UnaryExpr == nil {
		builder.internal.UnaryExpr = NewUnaryExpr()
	}
	exprResource, err := expr.Build()
	if err != nil {
		builder.errors["UnaryExpr.expr"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.UnaryExpr.Expr = exprResource

	return builder
}
