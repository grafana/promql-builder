---
title: <span class="badge builder"></span> VectorExprBuilder
---
# <span class="badge builder"></span> VectorExprBuilder

## Constructor

```go
func NewVectorExprBuilder() *VectorExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VectorExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> At

The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.

The time supplied to the @ modifier is a unix timestamp.

https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier

```go
func (builder *VectorExprBuilder) At(at string) *VectorExprBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *VectorExprBuilder) Label(name string, value string) *VectorExprBuilder
```

### <span class="badge object-method"></span> LabelMatchRegexp

```go
func (builder *VectorExprBuilder) LabelMatchRegexp(name string, value string) *VectorExprBuilder
```

### <span class="badge object-method"></span> LabelNeq

```go
func (builder *VectorExprBuilder) LabelNeq(name string, value string) *VectorExprBuilder
```

### <span class="badge object-method"></span> LabelNotMatchRegexp

```go
func (builder *VectorExprBuilder) LabelNotMatchRegexp(name string, value string) *VectorExprBuilder
```

### <span class="badge object-method"></span> Labels

Label selectors used to filter the timeseries.

```go
func (builder *VectorExprBuilder) Labels(labels []cog.Builder[promql.LabelSelector]) *VectorExprBuilder
```

### <span class="badge object-method"></span> Metric

Metric name.

```go
func (builder *VectorExprBuilder) Metric(metric string) *VectorExprBuilder
```

### <span class="badge object-method"></span> Offset

The offset modifier allows changing the time offset for individual instant and range vectors in a query.

https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier

```go
func (builder *VectorExprBuilder) Offset(offset string) *VectorExprBuilder
```

### <span class="badge object-method"></span> Range

Range of samples back from the current instant.

https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors

```go
func (builder *VectorExprBuilder) Range(rangeArg string) *VectorExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
