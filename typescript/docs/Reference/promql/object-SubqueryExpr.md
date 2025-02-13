---
title: <span class="badge object-type-interface"></span> SubqueryExpr
---
# <span class="badge object-type-interface"></span> SubqueryExpr

Represents a subquery.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery

## Definition

```typescript
export interface SubqueryExpr {
	type: "subqueryExpr";
	expr: promql.Expr;
	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	offset: string;
	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	at: string;
	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	range: string;
	// Empty string for default resolution.
	resolution?: string;
}

```
## See also

 * <span class="badge builder"></span> [SubqueryExprBuilder](./builder-SubqueryExprBuilder.md)
