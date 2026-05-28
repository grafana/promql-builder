// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*SubqueryExprBuilder)(nil)

// Represents a PromQL expression.
type SubqueryExprBuilder struct {
	internal *Expr
	errors   cog.BuildErrors
}

func NewSubqueryExprBuilder() *SubqueryExprBuilder {
	resource := NewExpr()
	builder := &SubqueryExprBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	if builder.internal.SubqueryExpr == nil {
		builder.internal.SubqueryExpr = NewSubqueryExpr()
	}
	builder.internal.SubqueryExpr.Type = "subqueryExpr"

	return builder
}

// Creates a subquery.
// Subquery allows you to run an instant query for a given range and resolution. The result of a subquery is a range vector.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
func Subquery(expression cog.Builder[Expr]) *SubqueryExprBuilder {
	builder := NewSubqueryExprBuilder()

	builder.Expr(expression)

	return builder
}

func (builder *SubqueryExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	if len(builder.errors) > 0 {
		return Expr{}, cog.MakeBuildErrors("promql.subqueryExpr", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SubqueryExprBuilder) RecordError(path string, err error) *SubqueryExprBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder SubqueryExprBuilder) String() string {
	return builder.internal.String()
}

func (builder *SubqueryExprBuilder) Expr(expr cog.Builder[Expr]) *SubqueryExprBuilder {
	if builder.internal.SubqueryExpr == nil {
		builder.internal.SubqueryExpr = NewSubqueryExpr()
	}
	exprResource, err := expr.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.SubqueryExpr.Expr = exprResource

	return builder
}

// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
func (builder *SubqueryExprBuilder) Offset(offset string) *SubqueryExprBuilder {
	if builder.internal.SubqueryExpr == nil {
		builder.internal.SubqueryExpr = NewSubqueryExpr()
	}
	builder.internal.SubqueryExpr.Offset = offset

	return builder
}

// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
// The time supplied to the @ modifier is a unix timestamp.
// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
func (builder *SubqueryExprBuilder) At(at string) *SubqueryExprBuilder {
	if builder.internal.SubqueryExpr == nil {
		builder.internal.SubqueryExpr = NewSubqueryExpr()
	}
	builder.internal.SubqueryExpr.At = at

	return builder
}

// Range of samples back from the current instant.
// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
func (builder *SubqueryExprBuilder) Range(rangeArg string) *SubqueryExprBuilder {
	if builder.internal.SubqueryExpr == nil {
		builder.internal.SubqueryExpr = NewSubqueryExpr()
	}
	builder.internal.SubqueryExpr.Range = rangeArg

	return builder
}

// Empty string for default resolution.
func (builder *SubqueryExprBuilder) Resolution(resolution string) *SubqueryExprBuilder {
	if builder.internal.SubqueryExpr == nil {
		builder.internal.SubqueryExpr = NewSubqueryExpr()
	}
	builder.internal.SubqueryExpr.Resolution = &resolution

	return builder
}
