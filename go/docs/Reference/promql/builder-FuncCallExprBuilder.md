---
title: <span class="badge builder"></span> FuncCallExprBuilder
---
# <span class="badge builder"></span> FuncCallExprBuilder

## Constructor

```go
func NewFuncCallExprBuilder() *FuncCallExprBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FuncCallExprBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> Arg

Arguments.

```go
func (builder *FuncCallExprBuilder) Arg(arg cog.Builder[promql.Expr]) *FuncCallExprBuilder
```

### <span class="badge object-method"></span> Args

Arguments.

```go
func (builder *FuncCallExprBuilder) Args(args []cog.Builder[promql.Expr]) *FuncCallExprBuilder
```

### <span class="badge object-method"></span> Function

Name of the function.

```go
func (builder *FuncCallExprBuilder) Function(function string) *FuncCallExprBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
