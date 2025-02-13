---
title: <span class="badge builder"></span> AggregationExprBuilder
---
# <span class="badge builder"></span> AggregationExprBuilder

## Constructor

```go
func NewAggregationExprBuilder() *AggregationExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AggregationExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> By

By drops labels that are not listed in the by clause.

```go
func (builder *AggregationExprBuilder) By(by []string) *AggregationExprBuilder
```

### <span class="badge object-method"></span> Expr

```go
func (builder *AggregationExprBuilder) Expr(expr cog.Builder[promql.Expr]) *AggregationExprBuilder
```

### <span class="badge object-method"></span> Op

```go
func (builder *AggregationExprBuilder) Op(op promql.AggregationOp) *AggregationExprBuilder
```

### <span class="badge object-method"></span> Param

```go
func (builder *AggregationExprBuilder) Param(param cog.Builder[promql.Expr]) *AggregationExprBuilder
```

### <span class="badge object-method"></span> Without

List of labels to remove from the result vector, while all other labels are preserved in the output.

```go
func (builder *AggregationExprBuilder) Without(without []string) *AggregationExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
