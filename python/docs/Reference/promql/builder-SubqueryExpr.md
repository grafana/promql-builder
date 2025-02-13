---
title: <span class="badge builder"></span> SubqueryExpr
---
# <span class="badge builder"></span> SubqueryExpr

## Constructor

```python
SubqueryExpr()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> promql.SubqueryExpr
```

### <span class="badge object-method"></span> at

The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.

The time supplied to the @ modifier is a unix timestamp.

https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier

```python
def at(at: str) -> typing.Self
```

### <span class="badge object-method"></span> expr

```python
def expr(expr: cogbuilder.Builder[promql.Expr]) -> typing.Self
```

### <span class="badge object-method"></span> offset

The offset modifier allows changing the time offset for individual instant and range vectors in a query.

https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier

```python
def offset(offset: str) -> typing.Self
```

### <span class="badge object-method"></span> range_val

Range of samples back from the current instant.

https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors

```python
def range_val(range_val: str) -> typing.Self
```

### <span class="badge object-method"></span> resolution

Empty string for default resolution.

```python
def resolution(resolution: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SubqueryExpr](./object-SubqueryExpr.md)
