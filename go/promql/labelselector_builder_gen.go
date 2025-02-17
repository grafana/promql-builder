// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[LabelSelector] = (*LabelSelectorBuilder)(nil)

type LabelSelectorBuilder struct {
	internal *LabelSelector
	errors   map[string]cog.BuildErrors
}

func NewLabelSelectorBuilder() *LabelSelectorBuilder {
	resource := NewLabelSelector()
	builder := &LabelSelectorBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *LabelSelectorBuilder) Build() (LabelSelector, error) {
	if err := builder.internal.Validate(); err != nil {
		return LabelSelector{}, err
	}

	return *builder.internal, nil
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
