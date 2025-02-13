---
title: <span class="badge object-type-enum"></span> LabelMatchingOperator
---
# <span class="badge object-type-enum"></span> LabelMatchingOperator

Possible label matching operators.

## Definition

```python
class LabelMatchingOperator(enum.StrEnum):
    """
    Possible label matching operators.
    """

    EQUAL = "="
    NOT_EQUAL = "!="
    MATCH_REGEXP = "=~"
    NOT_MATCH_REGEXP = "!~"
```
