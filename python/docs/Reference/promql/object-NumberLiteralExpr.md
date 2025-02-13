---
title: <span class="badge object-type-class"></span> NumberLiteralExpr
---
# <span class="badge object-type-class"></span> NumberLiteralExpr

Represents a number literal expression.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations

## Definition

```python
class NumberLiteralExpr:
    """
    Represents a number literal expression.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
    """

    type_val: typing.Literal["numberLiteralExpr"]
    value: float
```
## See also

 * <span class="badge builder"></span> [NumberLiteralExpr](./builder-NumberLiteralExpr.md)
