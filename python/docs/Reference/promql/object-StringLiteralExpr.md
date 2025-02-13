---
title: <span class="badge object-type-class"></span> StringLiteralExpr
---
# <span class="badge object-type-class"></span> StringLiteralExpr

Represents a string literal expression.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals

## Definition

```python
class StringLiteralExpr:
    """
    Represents a string literal expression.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
    """

    type_val: typing.Literal["stringLiteralExpr"]
    value: str
```
## See also

 * <span class="badge builder"></span> [StringLiteralExpr](./builder-StringLiteralExpr.md)
