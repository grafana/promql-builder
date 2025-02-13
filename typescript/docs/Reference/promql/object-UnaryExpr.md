---
title: <span class="badge object-type-interface"></span> UnaryExpr
---
# <span class="badge object-type-interface"></span> UnaryExpr

Represents an unary operation expression.

## Definition

```typescript
export interface UnaryExpr {
	type: "unaryExpr";
	op: promql.UnaryOp;
	expr: promql.Expr;
}

```
## See also

 * <span class="badge builder"></span> [UnaryExprBuilder](./builder-UnaryExprBuilder.md)
