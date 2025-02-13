---
title: <span class="badge builder"></span> AggregationExpr
---
# <span class="badge builder"></span> AggregationExpr

## Constructor

```python
AggregationExpr()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> promql.AggregationExpr
```

### <span class="badge object-method"></span> by

By drops labels that are not listed in the by clause.

```python
def by(by: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> expr

```python
def expr(expr: cogbuilder.Builder[promql.Expr]) -> typing.Self
```

### <span class="badge object-method"></span> op

```python
def op(op: promql.AggregationOp) -> typing.Self
```

### <span class="badge object-method"></span> param

```python
def param(param: cogbuilder.Builder[promql.Expr]) -> typing.Self
```

### <span class="badge object-method"></span> without

List of labels to remove from the result vector, while all other labels are preserved in the output.

```python
def without(without: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AggregationExpr](./object-AggregationExpr.md)
