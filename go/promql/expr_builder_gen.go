// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*ExprBuilder)(nil)

// Represents a PromQL expression.
type ExprBuilder struct {
	internal *Expr
	errors   cog.BuildErrors
}

func NewExprBuilder() *ExprBuilder {
	resource := NewExpr()
	builder := &ExprBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	if len(builder.errors) > 0 {
		return Expr{}, cog.MakeBuildErrors("promql.expr", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprBuilder) RecordError(path string, err error) *ExprBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ExprBuilder) NumberLiteralExpr(numberLiteralExpr NumberLiteralExpr) *ExprBuilder {
	builder.internal.NumberLiteralExpr = &numberLiteralExpr

	return builder
}

func (builder *ExprBuilder) StringLiteralExpr(stringLiteralExpr StringLiteralExpr) *ExprBuilder {
	builder.internal.StringLiteralExpr = &stringLiteralExpr

	return builder
}

func (builder *ExprBuilder) SubqueryExpr(subqueryExpr SubqueryExpr) *ExprBuilder {
	builder.internal.SubqueryExpr = &subqueryExpr

	return builder
}

func (builder *ExprBuilder) AggregationExpr(aggregationExpr AggregationExpr) *ExprBuilder {
	builder.internal.AggregationExpr = &aggregationExpr

	return builder
}

func (builder *ExprBuilder) VectorExpr(vectorExpr VectorExpr) *ExprBuilder {
	builder.internal.VectorExpr = &vectorExpr

	return builder
}

func (builder *ExprBuilder) BinaryExpr(binaryExpr BinaryExpr) *ExprBuilder {
	builder.internal.BinaryExpr = &binaryExpr

	return builder
}

func (builder *ExprBuilder) UnaryExpr(unaryExpr UnaryExpr) *ExprBuilder {
	builder.internal.UnaryExpr = &unaryExpr

	return builder
}

func (builder *ExprBuilder) FuncCallExpr(funcCallExpr FuncCallExpr) *ExprBuilder {
	builder.internal.FuncCallExpr = &funcCallExpr

	return builder
}
