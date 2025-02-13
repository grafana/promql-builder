---
title: <span class="badge builder"></span> BinaryExpr
---
# <span class="badge builder"></span> BinaryExpr

## Constructor

```python
BinaryExpr()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> promql.BinaryExpr
```

### <span class="badge object-method"></span> group_left

See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches

```python
def group_left(labels: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> group_right

See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches

```python
def group_right(labels: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> ignoring

Allows ignoring certain labels when matching.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches

```python
def ignoring(labels: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> left

```python
def left(left: cogbuilder.Builder[promql.Expr]) -> typing.Self
```

### <span class="badge object-method"></span> on

Allows reducing the set of considered labels to a provided list when matching.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches

```python
def on(labels: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> op

```python
def op(op: promql.BinaryOp) -> typing.Self
```

### <span class="badge object-method"></span> right

```python
def right(right: cogbuilder.Builder[promql.Expr]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [BinaryExpr](./object-BinaryExpr.md)
