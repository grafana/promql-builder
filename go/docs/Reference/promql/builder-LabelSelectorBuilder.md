---
title: <span class="badge builder"></span> LabelSelectorBuilder
---
# <span class="badge builder"></span> LabelSelectorBuilder

## Constructor

```go
func NewLabelSelectorBuilder() *LabelSelectorBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LabelSelectorBuilder) Build() (LabelSelector, error)
```

### <span class="badge object-method"></span> Name

Name of the label to select.

```go
func (builder *LabelSelectorBuilder) Name(name string) *LabelSelectorBuilder
```

### <span class="badge object-method"></span> Operator

Operator used to perform the selection.

```go
func (builder *LabelSelectorBuilder) Operator(operator promql.LabelMatchingOperator) *LabelSelectorBuilder
```

### <span class="badge object-method"></span> Value

Value to match against.

```go
func (builder *LabelSelectorBuilder) Value(value string) *LabelSelectorBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LabelSelector](./object-LabelSelector.md)
