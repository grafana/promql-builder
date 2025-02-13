---
title: <span class="badge object-type-struct"></span> StringLiteralExpr
---
# <span class="badge object-type-struct"></span> StringLiteralExpr

Represents a string literal expression.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals

## Definition

```go
type StringLiteralExpr struct {
    Type string `json:"type"`
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringLiteralExpr` fields for violations and returns them.

```go
func (stringLiteralExpr *StringLiteralExpr) Validate() error
```

