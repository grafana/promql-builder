// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*NumberLiteralBuilder)(nil)

// Represents a PromQL expression.
type NumberLiteralBuilder struct {
	internal *Expr
	errors   cog.BuildErrors
}

func NewNumberLiteralBuilder() *NumberLiteralBuilder {
	resource := NewExpr()
	builder := &NumberLiteralBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	if builder.internal.NumberLiteralExpr == nil {
		builder.internal.NumberLiteralExpr = NewNumberLiteralExpr()
	}
	builder.internal.NumberLiteralExpr.Type = "numberLiteralExpr"

	return builder
}

// Shortcut to turn a number into a NumberLiteral expression.
func N(value float64) *NumberLiteralBuilder {
	builder := NewNumberLiteralBuilder()

	builder.Value(value)

	return builder
}

func (builder *NumberLiteralBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	if len(builder.errors) > 0 {
		return Expr{}, cog.MakeBuildErrors("promql.numberLiteral", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NumberLiteralBuilder) RecordError(path string, err error) *NumberLiteralBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder NumberLiteralBuilder) String() string {
	return builder.internal.String()
}

func (builder *NumberLiteralBuilder) Value(value float64) *NumberLiteralBuilder {
	if builder.internal.NumberLiteralExpr == nil {
		builder.internal.NumberLiteralExpr = NewNumberLiteralExpr()
	}
	builder.internal.NumberLiteralExpr.Value = value

	return builder
}
