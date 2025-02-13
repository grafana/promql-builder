---
title: <span class="badge object-type-struct"></span> SubqueryExpr
---
# <span class="badge object-type-struct"></span> SubqueryExpr

Represents a subquery.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery

## Definition

```go
type SubqueryExpr struct {
    Type string `json:"type"`
    Expr promql.Expr `json:"expr"`
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
    // Empty string for default resolution.
    Resolution *string `json:"resolution,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SubqueryExpr` fields for violations and returns them.

```go
func (subqueryExpr *SubqueryExpr) Validate() error
```

