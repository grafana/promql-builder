---
title: <span class="badge object-type-struct"></span> UnaryExpr
---
# <span class="badge object-type-struct"></span> UnaryExpr

Represents an unary operation expression.

## Definition

```go
type UnaryExpr struct {
    Type string `json:"type"`
    Op promql.UnaryOp `json:"op"`
    Expr promql.Expr `json:"expr"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `UnaryExpr` fields for violations and returns them.

```go
func (unaryExpr *UnaryExpr) Validate() error
```

