---
title: <span class="badge object-type-interface"></span> NumberLiteralExpr
---
# <span class="badge object-type-interface"></span> NumberLiteralExpr

Represents a number literal expression.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations

## Definition

```typescript
export interface NumberLiteralExpr {
	type: "numberLiteralExpr";
	value: number;
}

```
## See also

 * <span class="badge builder"></span> [NumberLiteralExprBuilder](./builder-NumberLiteralExprBuilder.md)
