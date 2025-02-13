---
title: <span class="badge object-type-class"></span> AggregationExpr
---
# <span class="badge object-type-class"></span> AggregationExpr

Represents an aggregation.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

## Definition

```python
class AggregationExpr:
    """
    Represents an aggregation.
    See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
    """

    type_val: typing.Literal["aggregationExpr"]
    op: promql.AggregationOp
    expr: promql.Expr
    param: typing.Optional[promql.Expr]
    # By drops labels that are not listed in the by clause.
    by: list[str]
    # List of labels to remove from the result vector, while all other labels are preserved in the output.
    without: list[str]
```
## See also

 * <span class="badge builder"></span> [AggregationExpr](./builder-AggregationExpr.md)
