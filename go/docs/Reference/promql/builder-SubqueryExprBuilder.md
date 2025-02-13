---
title: <span class="badge builder"></span> SubqueryExprBuilder
---
# <span class="badge builder"></span> SubqueryExprBuilder

## Constructor

```go
func NewSubqueryExprBuilder() *SubqueryExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SubqueryExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> At

The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.

The time supplied to the @ modifier is a unix timestamp.

https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier

```go
func (builder *SubqueryExprBuilder) At(at string) *SubqueryExprBuilder
```

### <span class="badge object-method"></span> Expr

```go
func (builder *SubqueryExprBuilder) Expr(expr cog.Builder[promql.Expr]) *SubqueryExprBuilder
```

### <span class="badge object-method"></span> Offset

The offset modifier allows changing the time offset for individual instant and range vectors in a query.

https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier

```go
func (builder *SubqueryExprBuilder) Offset(offset string) *SubqueryExprBuilder
```

### <span class="badge object-method"></span> Range

Range of samples back from the current instant.

https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors

```go
func (builder *SubqueryExprBuilder) Range(rangeArg string) *SubqueryExprBuilder
```

### <span class="badge object-method"></span> Resolution

Empty string for default resolution.

```go
func (builder *SubqueryExprBuilder) Resolution(resolution string) *SubqueryExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
