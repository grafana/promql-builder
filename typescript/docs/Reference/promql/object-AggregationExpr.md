---
title: <span class="badge object-type-interface"></span> AggregationExpr
---
# <span class="badge object-type-interface"></span> AggregationExpr

Represents an aggregation.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

## Definition

```typescript
export interface AggregationExpr {
	type: "aggregationExpr";
	op: promql.AggregationOp;
	expr: promql.Expr;
	param?: promql.Expr;
	// By drops labels that are not listed in the by clause.
	by: string[];
	// List of labels to remove from the result vector, while all other labels are preserved in the output.
	without: string[];
}

```
## See also

 * <span class="badge builder"></span> [AggregationExprBuilder](./builder-AggregationExprBuilder.md)
