---
title: <span class="badge builder"></span> VectorExpr
---
# <span class="badge builder"></span> VectorExpr

## Constructor

```python
VectorExpr()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> promql.VectorExpr
```

### <span class="badge object-method"></span> at

The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.

The time supplied to the @ modifier is a unix timestamp.

https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier

```python
def at(at: str) -> typing.Self
```

### <span class="badge object-method"></span> label

```python
def label(name: str, value: str) -> typing.Self
```

### <span class="badge object-method"></span> label_match_regexp

```python
def label_match_regexp(name: str, value: str) -> typing.Self
```

### <span class="badge object-method"></span> label_neq

```python
def label_neq(name: str, value: str) -> typing.Self
```

### <span class="badge object-method"></span> label_not_match_regexp

```python
def label_not_match_regexp(name: str, value: str) -> typing.Self
```

### <span class="badge object-method"></span> labels

Label selectors used to filter the timeseries.

```python
def labels(labels: list[cogbuilder.Builder[promql.LabelSelector]]) -> typing.Self
```

### <span class="badge object-method"></span> metric

Metric name.

```python
def metric(metric: str) -> typing.Self
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

## See also

 * <span class="badge object-type-class"></span> [VectorExpr](./object-VectorExpr.md)
