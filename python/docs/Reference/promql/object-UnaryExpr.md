---
title: <span class="badge object-type-class"></span> UnaryExpr
---
# <span class="badge object-type-class"></span> UnaryExpr

Represents an unary operation expression.

## Definition

```python
class UnaryExpr:
    """
    Represents an unary operation expression.
    """

    type_val: typing.Literal["unaryExpr"]
    op: promql.UnaryOp
    expr: promql.Expr
```
## See also

 * <span class="badge builder"></span> [UnaryExpr](./builder-UnaryExpr.md)
