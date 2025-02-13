---
title: <span class="badge object-type-struct"></span> LabelSelector
---
# <span class="badge object-type-struct"></span> LabelSelector

## Definition

```go
type LabelSelector struct {
    // Name of the label to select.
    Name string `json:"name"`
    // Value to match against.
    Value string `json:"value"`
    // Operator used to perform the selection.
    Operator promql.LabelMatchingOperator `json:"operator"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LabelSelector` fields for violations and returns them.

```go
func (labelSelector *LabelSelector) Validate() error
```

## See also

 * <span class="badge builder"></span> [LabelSelectorBuilder](./builder-LabelSelectorBuilder.md)
