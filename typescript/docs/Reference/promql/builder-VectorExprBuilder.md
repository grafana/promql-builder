---
title: <span class="badge builder"></span> VectorExprBuilder
---
# <span class="badge builder"></span> VectorExprBuilder

## Constructor

```typescript
new VectorExprBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> at

The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.

The time supplied to the @ modifier is a unix timestamp.

https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier

```typescript
at(at: string)
```

### <span class="badge object-method"></span> label

```typescript
label(name: string, value: string)
```

### <span class="badge object-method"></span> labelMatchRegexp

```typescript
labelMatchRegexp(name: string, value: string)
```

### <span class="badge object-method"></span> labelNeq

```typescript
labelNeq(name: string, value: string)
```

### <span class="badge object-method"></span> labelNotMatchRegexp

```typescript
labelNotMatchRegexp(name: string, value: string)
```

### <span class="badge object-method"></span> labels

Label selectors used to filter the timeseries.

```typescript
labels(labels: cog.Builder<promql.LabelSelector>[])
```

### <span class="badge object-method"></span> metric

Metric name.

```typescript
metric(metric: string)
```

### <span class="badge object-method"></span> offset

The offset modifier allows changing the time offset for individual instant and range vectors in a query.

https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier

```typescript
offset(offset: string)
```

### <span class="badge object-method"></span> range

Range of samples back from the current instant.

https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors

```typescript
range(range: string)
```

## See also

 * <span class="badge object-type-interface"></span> [VectorExpr](./object-VectorExpr.md)
