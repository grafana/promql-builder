---
title: <span class="badge object-type-interface"></span> BinaryExpr
---
# <span class="badge object-type-interface"></span> BinaryExpr

Represents a binary operation expression.

## Definition

```typescript
export interface BinaryExpr {
	type: "binaryExpr";
	op: promql.BinaryOp;
	left: promql.Expr;
	right: promql.Expr;
	// https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
	matchType?: "on" | "ignore";
	matchLabels?: string[];
	groupModifier?: "left" | "right";
	groupLabels?: string[];
}

```
## See also

 * <span class="badge builder"></span> [BinaryExprBuilder](./builder-BinaryExprBuilder.md)
