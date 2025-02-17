// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*StringLiteralBuilder)(nil)

// Represents a PromQL expression.
type StringLiteralBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewStringLiteralBuilder() *StringLiteralBuilder {
	resource := NewExpr()
	builder := &StringLiteralBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	if builder.internal.StringLiteralExpr == nil {
		builder.internal.StringLiteralExpr = NewStringLiteralExpr()
	}
	builder.internal.StringLiteralExpr.Type = "stringLiteralExpr"

	return builder
}

// Shortcut to turn a string into a StringLiteral expression.
func S(value string) *StringLiteralBuilder {
	builder := NewStringLiteralBuilder()
	builder.Value(value)

	return builder
}

func (builder *StringLiteralBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	return *builder.internal, nil
}

func (builder StringLiteralBuilder) String() string {
	return builder.internal.String()
}

func (builder *StringLiteralBuilder) Value(value string) *StringLiteralBuilder {
	if builder.internal.StringLiteralExpr == nil {
		builder.internal.StringLiteralExpr = NewStringLiteralExpr()
	}
	builder.internal.StringLiteralExpr.Value = value

	return builder
}
