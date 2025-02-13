---
title: <span class="badge object-type-struct"></span> AggregationExpr
---
# <span class="badge object-type-struct"></span> AggregationExpr

Represents an aggregation.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

## Definition

```go
type AggregationExpr struct {
    Type string `json:"type"`
    Op promql.AggregationOp `json:"op"`
    Expr promql.Expr `json:"expr"`
    Param *promql.Expr `json:"param,omitempty"`
    // By drops labels that are not listed in the by clause.
    By []string `json:"by"`
    // List of labels to remove from the result vector, while all other labels are preserved in the output.
    Without []string `json:"without"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AggregationExpr` fields for violations and returns them.

```go
func (aggregationExpr *AggregationExpr) Validate() error
```

