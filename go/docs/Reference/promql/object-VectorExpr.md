---
title: <span class="badge object-type-struct"></span> VectorExpr
---
# <span class="badge object-type-struct"></span> VectorExpr

Represents both instant and range vectors

## Definition

```go
type VectorExpr struct {
    Type string `json:"type"`
    // Metric name.
    Metric string `json:"metric"`
    // Label selectors used to filter the timeseries.
    Labels []promql.LabelSelector `json:"labels"`
    // The offset modifier allows changing the time offset for individual instant and range vectors in a query.
    // https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
    Offset string `json:"offset"`
    // The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
    // The time supplied to the @ modifier is a unix timestamp.
    // https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
    At string `json:"at"`
    // Range of samples back from the current instant.
    // https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
    Range string `json:"range"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `VectorExpr` fields for violations and returns them.

```go
func (vectorExpr *VectorExpr) Validate() error
```

