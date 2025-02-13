---
title: <span class="badge object-type-class"></span> LabelSelector
---
# <span class="badge object-type-class"></span> LabelSelector

## Definition

```python
class LabelSelector:
    # Name of the label to select.
    name: str
    # Value to match against.
    value: str
    # Operator used to perform the selection.
    operator: promql.LabelMatchingOperator
```
## See also

 * <span class="badge builder"></span> [LabelSelector](./builder-LabelSelector.md)
