---
title: <span class="badge builder"></span> LabelSelector
---
# <span class="badge builder"></span> LabelSelector

## Constructor

```python
LabelSelector()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> promql.LabelSelector
```

### <span class="badge object-method"></span> name

Name of the label to select.

```python
def name(name: str) -> typing.Self
```

### <span class="badge object-method"></span> operator

Operator used to perform the selection.

```python
def operator(operator: promql.LabelMatchingOperator) -> typing.Self
```

### <span class="badge object-method"></span> value

Value to match against.

```python
def value(value: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [LabelSelector](./object-LabelSelector.md)
