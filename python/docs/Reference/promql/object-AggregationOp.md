---
title: <span class="badge object-type-enum"></span> AggregationOp
---
# <span class="badge object-type-enum"></span> AggregationOp

Possible aggregation operators.

## Definition

```python
class AggregationOp(enum.StrEnum):
    """
    Possible aggregation operators.
    """

    SUM = "sum"
    MIN = "min"
    MAX = "max"
    AVG = "avg"
    STDDEV = "stddev"
    STDVAR = "stdvar"
    COUNT = "count"
    GROUP = "group"
    COUNT_VALUES = "count_values"
    BOTTOMK = "bottomk"
    TOPK = "topk"
    QUANTILE = "quantile"
    LIMITK = "limitk"
    LIMIT_RATIO = "limit_ratio"
```
