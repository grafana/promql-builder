---
title: <span class="badge object-type-class"></span> BinaryExpr
---
# <span class="badge object-type-class"></span> BinaryExpr

Represents a binary operation expression.

## Definition

```python
class BinaryExpr:
    """
    Represents a binary operation expression.
    """

    type_val: typing.Literal["binaryExpr"]
    op: promql.BinaryOp
    left: promql.Expr
    right: promql.Expr
    # https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
    match_type: typing.Optional[typing.Literal["on", "ignore"]]
    match_labels: typing.Optional[list[str]]
    group_modifier: typing.Optional[typing.Literal["left", "right"]]
    group_labels: typing.Optional[list[str]]
```
## See also

 * <span class="badge builder"></span> [BinaryExpr](./builder-BinaryExpr.md)
