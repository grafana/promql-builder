---
title: <span class="badge object-type-class"></span> VectorExpr
---
# <span class="badge object-type-class"></span> VectorExpr

Represents both instant and range vectors

## Definition

```python
class VectorExpr:
    """
    Represents both instant and range vectors
    """

    type_val: typing.Literal["vectorExpr"]
    # Metric name.
    metric: str
    # Label selectors used to filter the timeseries.
    labels: list[promql.LabelSelector]
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
```
## See also

 * <span class="badge builder"></span> [VectorExpr](./builder-VectorExpr.md)
