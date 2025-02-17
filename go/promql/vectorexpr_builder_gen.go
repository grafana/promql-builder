// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*VectorExprBuilder)(nil)

// Represents a PromQL expression.
type VectorExprBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewVectorExprBuilder() *VectorExprBuilder {
	resource := NewExpr()
	builder := &VectorExprBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Type = "vectorExpr"

	return builder
}

// Returns the scalar s as a vector with no labels.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector
func Vector(s string) *VectorExprBuilder {
	builder := NewVectorExprBuilder()
	builder.Metric(s)

	return builder
}

func (builder *VectorExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	return *builder.internal, nil
}

func (builder VectorExprBuilder) String() string {
	return builder.internal.String()
}

// Metric name.
func (builder *VectorExprBuilder) Metric(metric string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Metric = metric

	return builder
}

// Label selectors used to filter the timeseries.
func (builder *VectorExprBuilder) Labels(labels []cog.Builder[LabelSelector]) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	labelsResources := make([]LabelSelector, 0, len(labels))
	for _, r1 := range labels {
		labelsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["VectorExpr.labels"] = err.(cog.BuildErrors)
			return builder
		}
		labelsResources = append(labelsResources, labelsDepth1)
	}
	builder.internal.VectorExpr.Labels = labelsResources

	return builder
}

// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
func (builder *VectorExprBuilder) Offset(offset string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Offset = offset

	return builder
}

// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
// The time supplied to the @ modifier is a unix timestamp.
// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
func (builder *VectorExprBuilder) At(at string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.At = at

	return builder
}

// Range of samples back from the current instant.
// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
func (builder *VectorExprBuilder) Range(rangeArg string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Range = rangeArg

	return builder
}

func (builder *VectorExprBuilder) Label(name string, value string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Labels = append(builder.internal.VectorExpr.Labels, LabelSelector{
		Name:     name,
		Operator: "=",
		Value:    value,
	})

	return builder
}

func (builder *VectorExprBuilder) LabelNeq(name string, value string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Labels = append(builder.internal.VectorExpr.Labels, LabelSelector{
		Name:     name,
		Operator: "!=",
		Value:    value,
	})

	return builder
}

func (builder *VectorExprBuilder) LabelMatchRegexp(name string, value string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Labels = append(builder.internal.VectorExpr.Labels, LabelSelector{
		Name:     name,
		Operator: "=~",
		Value:    value,
	})

	return builder
}

func (builder *VectorExprBuilder) LabelNotMatchRegexp(name string, value string) *VectorExprBuilder {
	if builder.internal.VectorExpr == nil {
		builder.internal.VectorExpr = NewVectorExpr()
	}
	builder.internal.VectorExpr.Labels = append(builder.internal.VectorExpr.Labels, LabelSelector{
		Name:     name,
		Operator: "!~",
		Value:    value,
	})

	return builder
}
