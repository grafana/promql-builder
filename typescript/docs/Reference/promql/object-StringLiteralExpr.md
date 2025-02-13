---
title: <span class="badge object-type-interface"></span> StringLiteralExpr
---
# <span class="badge object-type-interface"></span> StringLiteralExpr

Represents a string literal expression.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals

## Definition

```typescript
export interface StringLiteralExpr {
	type: "stringLiteralExpr";
	value: string;
}

```
## See also

 * <span class="badge builder"></span> [StringLiteralExprBuilder](./builder-StringLiteralExprBuilder.md)
