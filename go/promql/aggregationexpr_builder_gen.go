// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*AggregationExprBuilder)(nil)

// Represents a PromQL expression.
type AggregationExprBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewAggregationExprBuilder() *AggregationExprBuilder {
	resource := NewExpr()
	builder := &AggregationExprBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	if builder.internal.AggregationExpr == nil {
		builder.internal.AggregationExpr = NewAggregationExpr()
	}
	builder.internal.AggregationExpr.Type = "aggregationExpr"

	return builder
}

// Calculate sum over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Sum(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpSum)
	builder.Expr(vector)

	return builder
}

// Calculate minimum over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Min(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpMin)
	builder.Expr(vector)

	return builder
}

// Calculate maximum over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Max(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpMax)
	builder.Expr(vector)

	return builder
}

// Calculate the average over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Avg(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpAvg)
	builder.Expr(vector)

	return builder
}

// All values in the resulting vector are 1.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Group(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpGroup)
	builder.Expr(vector)

	return builder
}

// Calculate population standard deviation over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Stddev(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpStddev)
	builder.Expr(vector)

	return builder
}

// Calculate population standard variance over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Stdvar(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpStdvar)
	builder.Expr(vector)

	return builder
}

// Count number of elements in the vector.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Count(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpCount)
	builder.Expr(vector)

	return builder
}

// Calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Quantile(vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpQuantile)
	builder.Expr(vector)

	return builder
}

// Count number of elements with the same value.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func CountValues(label string, vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpCountValues)
	builder.Expr(vector)
	builder.Param(S(label))

	return builder
}

// Smallest k elements by sample value.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Bottomk(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpBottomk)
	builder.Expr(vector)
	builder.Param(N(k))

	return builder
}

// Largest k elements by sample value.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Topk(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpTopk)
	builder.Expr(vector)
	builder.Param(N(k))

	return builder
}

// Sample k elements.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func Limitk(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpLimitk)
	builder.Expr(vector)
	builder.Param(N(k))

	return builder
}

// Sample elements with approximately r ratio if r > 0, and the complement of such samples if r = -(1.0 - r).
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
func LimitRatio(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder {
	builder := NewAggregationExprBuilder()
	builder.Op(AggregationOpLimitRatio)
	builder.Expr(vector)
	builder.Param(N(k))

	return builder
}

func (builder *AggregationExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	return *builder.internal, nil
}

func (builder AggregationExprBuilder) String() string {
	return builder.internal.String()
}

func (builder *AggregationExprBuilder) Op(op AggregationOp) *AggregationExprBuilder {
	if builder.internal.AggregationExpr == nil {
		builder.internal.AggregationExpr = NewAggregationExpr()
	}
	builder.internal.AggregationExpr.Op = op

	return builder
}

func (builder *AggregationExprBuilder) Expr(expr cog.Builder[Expr]) *AggregationExprBuilder {
	if builder.internal.AggregationExpr == nil {
		builder.internal.AggregationExpr = NewAggregationExpr()
	}
	exprResource, err := expr.Build()
	if err != nil {
		builder.errors["AggregationExpr.expr"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.AggregationExpr.Expr = exprResource

	return builder
}

func (builder *AggregationExprBuilder) Param(param cog.Builder[Expr]) *AggregationExprBuilder {
	if builder.internal.AggregationExpr == nil {
		builder.internal.AggregationExpr = NewAggregationExpr()
	}
	paramResource, err := param.Build()
	if err != nil {
		builder.errors["AggregationExpr.param"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.AggregationExpr.Param = &paramResource

	return builder
}

// By drops labels that are not listed in the by clause.
func (builder *AggregationExprBuilder) By(by []string) *AggregationExprBuilder {
	if builder.internal.AggregationExpr == nil {
		builder.internal.AggregationExpr = NewAggregationExpr()
	}
	builder.internal.AggregationExpr.By = by

	return builder
}

// List of labels to remove from the result vector, while all other labels are preserved in the output.
func (builder *AggregationExprBuilder) Without(without []string) *AggregationExprBuilder {
	if builder.internal.AggregationExpr == nil {
		builder.internal.AggregationExpr = NewAggregationExpr()
	}
	builder.internal.AggregationExpr.Without = without

	return builder
}
