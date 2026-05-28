// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[LabelSelector] = (*LabelSelectorBuilder)(nil)

type LabelSelectorBuilder struct {
	internal *LabelSelector
	errors   cog.BuildErrors
}

func NewLabelSelectorBuilder() *LabelSelectorBuilder {
	resource := NewLabelSelector()
	builder := &LabelSelectorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func LabelEq(name string, value string) *LabelSelectorBuilder {
	builder := NewLabelSelectorBuilder()

	builder.Name(name)

	builder.Value(value)

	builder.Operator(LabelMatchingOperatorEqual)

	return builder
}

func LabelNeq(name string, value string) *LabelSelectorBuilder {
	builder := NewLabelSelectorBuilder()

	builder.Name(name)

	builder.Value(value)

	builder.Operator(LabelMatchingOperatorNotEqual)

	return builder
}

func LabelMatchRegexp(name string, value string) *LabelSelectorBuilder {
	builder := NewLabelSelectorBuilder()

	builder.Name(name)

	builder.Value(value)

	builder.Operator(LabelMatchingOperatorMatchRegexp)

	return builder
}

func LabelNotMatchRegexp(name string, value string) *LabelSelectorBuilder {
	builder := NewLabelSelectorBuilder()

	builder.Name(name)

	builder.Value(value)

	builder.Operator(LabelMatchingOperatorNotMatchRegexp)

	return builder
}

func (builder *LabelSelectorBuilder) Build() (LabelSelector, error) {
	if err := builder.internal.Validate(); err != nil {
		return LabelSelector{}, err
	}

	if len(builder.errors) > 0 {
		return LabelSelector{}, cog.MakeBuildErrors("promql.labelSelector", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LabelSelectorBuilder) RecordError(path string, err error) *LabelSelectorBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder LabelSelectorBuilder) String() string {
	return builder.internal.String()
}

// Name of the label to select.
func (builder *LabelSelectorBuilder) Name(name string) *LabelSelectorBuilder {
	builder.internal.Name = name

	return builder
}

// Value to match against.
func (builder *LabelSelectorBuilder) Value(value string) *LabelSelectorBuilder {
	builder.internal.Value = value

	return builder
}

// Operator used to perform the selection.
func (builder *LabelSelectorBuilder) Operator(operator LabelMatchingOperator) *LabelSelectorBuilder {
	builder.internal.Operator = operator

	return builder
}
