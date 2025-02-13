---
title: <span class="badge builder"></span> FuncCallExpr
---
# <span class="badge builder"></span> FuncCallExpr

## Constructor

```python
FuncCallExpr()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> promql.FuncCallExpr
```

### <span class="badge object-method"></span> arg

Arguments.

```python
def arg(arg: cogbuilder.Builder[promql.Expr]) -> typing.Self
```

### <span class="badge object-method"></span> args

Arguments.

```python
def args(args: list[cogbuilder.Builder[promql.Expr]]) -> typing.Self
```

### <span class="badge object-method"></span> function

Name of the function.

```python
def function(function: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [FuncCallExpr](./object-FuncCallExpr.md)
