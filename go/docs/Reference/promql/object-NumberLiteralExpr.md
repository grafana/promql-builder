---
title: <span class="badge object-type-struct"></span> NumberLiteralExpr
---
# <span class="badge object-type-struct"></span> NumberLiteralExpr

Represents a number literal expression.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations

## Definition

```go
type NumberLiteralExpr struct {
    Type string `json:"type"`
    Value float64 `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NumberLiteralExpr` fields for violations and returns them.

```go
func (numberLiteralExpr *NumberLiteralExpr) Validate() error
```

