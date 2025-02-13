---
title: <span class="badge object-type-class"></span> FuncCallExpr
---
# <span class="badge object-type-class"></span> FuncCallExpr

Represents a function call expression.

## Definition

```python
class FuncCallExpr:
    """
    Represents a function call expression.
    """

    type_val: typing.Literal["funcCallExpr"]
    # Name of the function.
    function: str
    # Arguments.
    args: list[promql.Expr]
```
## See also

 * <span class="badge builder"></span> [FuncCallExpr](./builder-FuncCallExpr.md)
