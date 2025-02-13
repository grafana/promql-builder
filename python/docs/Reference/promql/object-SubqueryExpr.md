---
title: <span class="badge object-type-class"></span> SubqueryExpr
---
# <span class="badge object-type-class"></span> SubqueryExpr

Represents a subquery.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery

## Definition

```python
class SubqueryExpr:
    """
    Represents a subquery.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
    """

    type_val: typing.Literal["subqueryExpr"]
    expr: promql.Expr
    # The offset modifier allows changing the time offset for individual instant and range vectors in a query.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
    offset: str
    # The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
    # The time supplied to the @ modifier is a unix timestamp.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
    at: str
    # Range of samples back from the current instant.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
    range_val: str
    # Empty string for default resolution.
    resolution: typing.Optional[str]
```
## See also

 * <span class="badge builder"></span> [SubqueryExpr](./builder-SubqueryExpr.md)
