---
title: <span class="badge object-type-interface"></span> VectorExpr
---
# <span class="badge object-type-interface"></span> VectorExpr

Represents both instant and range vectors

## Definition

```typescript
export interface VectorExpr {
	type: "vectorExpr";
	// Metric name.
	metric: string;
	// Label selectors used to filter the timeseries.
	labels: promql.LabelSelector[];
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
}

```
## See also

 * <span class="badge builder"></span> [VectorExprBuilder](./builder-VectorExprBuilder.md)
