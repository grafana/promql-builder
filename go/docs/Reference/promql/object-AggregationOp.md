---
title: <span class="badge object-type-enum"></span> AggregationOp
---
# <span class="badge object-type-enum"></span> AggregationOp

Possible aggregation operators.

## Definition

```go
type AggregationOp string
const (
	AggregationOpSum AggregationOp = "sum"
	AggregationOpMin AggregationOp = "min"
	AggregationOpMax AggregationOp = "max"
	AggregationOpAvg AggregationOp = "avg"
	AggregationOpStddev AggregationOp = "stddev"
	AggregationOpStdvar AggregationOp = "stdvar"
	AggregationOpCount AggregationOp = "count"
	AggregationOpGroup AggregationOp = "group"
	AggregationOpCountValues AggregationOp = "count_values"
	AggregationOpBottomk AggregationOp = "bottomk"
	AggregationOpTopk AggregationOp = "topk"
	AggregationOpQuantile AggregationOp = "quantile"
	AggregationOpLimitk AggregationOp = "limitk"
	AggregationOpLimitRatio AggregationOp = "limit_ratio"
)

```
