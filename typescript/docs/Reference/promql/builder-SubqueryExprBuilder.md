---
title: <span class="badge builder"></span> SubqueryExprBuilder
---
# <span class="badge builder"></span> SubqueryExprBuilder

## Constructor

```typescript
new SubqueryExprBuilder()
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

### <span class="badge object-method"></span> expr

```typescript
expr(expr: cog.Builder<promql.Expr>)
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

### <span class="badge object-method"></span> resolution

Empty string for default resolution.

```typescript
resolution(resolution: string)
```

## See also

 * <span class="badge object-type-interface"></span> [SubqueryExpr](./object-SubqueryExpr.md)
