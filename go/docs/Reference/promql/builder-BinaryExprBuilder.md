---
title: <span class="badge builder"></span> BinaryExprBuilder
---
# <span class="badge builder"></span> BinaryExprBuilder

## Constructor

```go
func NewBinaryExprBuilder() *BinaryExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BinaryExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> GroupLeft

See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches

```go
func (builder *BinaryExprBuilder) GroupLeft(labels []string) *BinaryExprBuilder
```

### <span class="badge object-method"></span> GroupRight

See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches

```go
func (builder *BinaryExprBuilder) GroupRight(labels []string) *BinaryExprBuilder
```

### <span class="badge object-method"></span> Ignoring

Allows ignoring certain labels when matching.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches

```go
func (builder *BinaryExprBuilder) Ignoring(labels []string) *BinaryExprBuilder
```

### <span class="badge object-method"></span> Left

```go
func (builder *BinaryExprBuilder) Left(left cog.Builder[promql.Expr]) *BinaryExprBuilder
```

### <span class="badge object-method"></span> On

Allows reducing the set of considered labels to a provided list when matching.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches

```go
func (builder *BinaryExprBuilder) On(labels []string) *BinaryExprBuilder
```

### <span class="badge object-method"></span> Op

```go
func (builder *BinaryExprBuilder) Op(op promql.BinaryOp) *BinaryExprBuilder
```

### <span class="badge object-method"></span> Right

```go
func (builder *BinaryExprBuilder) Right(right cog.Builder[promql.Expr]) *BinaryExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
