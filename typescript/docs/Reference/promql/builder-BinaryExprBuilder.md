---
title: <span class="badge builder"></span> BinaryExprBuilder
---
# <span class="badge builder"></span> BinaryExprBuilder

## Constructor

```typescript
new BinaryExprBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> groupLeft

See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches

```typescript
groupLeft(labels: string[])
```

### <span class="badge object-method"></span> groupRight

See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches

```typescript
groupRight(labels: string[])
```

### <span class="badge object-method"></span> ignoring

Allows ignoring certain labels when matching.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches

```typescript
ignoring(labels: string[])
```

### <span class="badge object-method"></span> left

```typescript
left(left: cog.Builder<promql.Expr>)
```

### <span class="badge object-method"></span> on

Allows reducing the set of considered labels to a provided list when matching.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches

```typescript
on(labels: string[])
```

### <span class="badge object-method"></span> op

```typescript
op(op: promql.BinaryOp)
```

### <span class="badge object-method"></span> right

```typescript
right(right: cog.Builder<promql.Expr>)
```

## See also

 * <span class="badge object-type-interface"></span> [BinaryExpr](./object-BinaryExpr.md)
