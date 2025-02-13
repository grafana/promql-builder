---
title: <span class="badge object-type-interface"></span> FuncCallExpr
---
# <span class="badge object-type-interface"></span> FuncCallExpr

Represents a function call expression.

## Definition

```typescript
export interface FuncCallExpr {
	type: "funcCallExpr";
	// Name of the function.
	function: string;
	// Arguments.
	args: promql.Expr[];
}

```
## See also

 * <span class="badge builder"></span> [FuncCallExprBuilder](./builder-FuncCallExprBuilder.md)
