---
title: <span class="badge builder"></span> UnaryExprBuilder
---
# <span class="badge builder"></span> UnaryExprBuilder

## Constructor

```go
func NewUnaryExprBuilder() *UnaryExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *UnaryExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> Expr

```go
func (builder *UnaryExprBuilder) Expr(expr cog.Builder[promql.Expr]) *UnaryExprBuilder
```

### <span class="badge object-method"></span> Op

```go
func (builder *UnaryExprBuilder) Op(op promql.UnaryOp) *UnaryExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
