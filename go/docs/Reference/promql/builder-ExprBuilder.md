---
title: <span class="badge builder"></span> ExprBuilder
---
# <span class="badge builder"></span> ExprBuilder

## Constructor

```go
func NewExprBuilder() *ExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> AggregationExpr

```go
func (builder *ExprBuilder) AggregationExpr(aggregationExpr promql.AggregationExpr) *ExprBuilder
```

### <span class="badge object-method"></span> BinaryExpr

```go
func (builder *ExprBuilder) BinaryExpr(binaryExpr promql.BinaryExpr) *ExprBuilder
```

### <span class="badge object-method"></span> FuncCallExpr

```go
func (builder *ExprBuilder) FuncCallExpr(funcCallExpr promql.FuncCallExpr) *ExprBuilder
```

### <span class="badge object-method"></span> NumberLiteralExpr

```go
func (builder *ExprBuilder) NumberLiteralExpr(numberLiteralExpr promql.NumberLiteralExpr) *ExprBuilder
```

### <span class="badge object-method"></span> StringLiteralExpr

```go
func (builder *ExprBuilder) StringLiteralExpr(stringLiteralExpr promql.StringLiteralExpr) *ExprBuilder
```

### <span class="badge object-method"></span> SubqueryExpr

```go
func (builder *ExprBuilder) SubqueryExpr(subqueryExpr promql.SubqueryExpr) *ExprBuilder
```

### <span class="badge object-method"></span> UnaryExpr

```go
func (builder *ExprBuilder) UnaryExpr(unaryExpr promql.UnaryExpr) *ExprBuilder
```

### <span class="badge object-method"></span> VectorExpr

```go
func (builder *ExprBuilder) VectorExpr(vectorExpr promql.VectorExpr) *ExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
