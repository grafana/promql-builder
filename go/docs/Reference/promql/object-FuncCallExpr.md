---
title: <span class="badge object-type-struct"></span> FuncCallExpr
---
# <span class="badge object-type-struct"></span> FuncCallExpr

Represents a function call expression.

## Definition

```go
type FuncCallExpr struct {
    Type string `json:"type"`
    // Name of the function.
    Function string `json:"function"`
    // Arguments.
    Args []promql.Expr `json:"args"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `FuncCallExpr` fields for violations and returns them.

```go
func (funcCallExpr *FuncCallExpr) Validate() error
```

