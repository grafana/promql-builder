---
title: <span class="badge builder"></span> AggregationExprBuilder
---
# <span class="badge builder"></span> AggregationExprBuilder

## Constructor

```typescript
new AggregationExprBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> by

By drops labels that are not listed in the by clause.

```typescript
by(by: string[])
```

### <span class="badge object-method"></span> expr

```typescript
expr(expr: cog.Builder<promql.Expr>)
```

### <span class="badge object-method"></span> op

```typescript
op(op: promql.AggregationOp)
```

### <span class="badge object-method"></span> param

```typescript
param(param: cog.Builder<promql.Expr>)
```

### <span class="badge object-method"></span> without

List of labels to remove from the result vector, while all other labels are preserved in the output.

```typescript
without(without: string[])
```

## See also

 * <span class="badge object-type-interface"></span> [AggregationExpr](./object-AggregationExpr.md)
