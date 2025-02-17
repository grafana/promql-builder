// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*NumberLiteralBuilder)(nil)

// Represents a PromQL expression.
type NumberLiteralBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewNumberLiteralBuilder() *NumberLiteralBuilder {
	resource := NewExpr()
	builder := &NumberLiteralBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
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

	return *builder.internal, nil
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
