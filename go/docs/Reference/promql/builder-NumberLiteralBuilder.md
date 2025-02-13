---
title: <span class="badge builder"></span> NumberLiteralBuilder
---
# <span class="badge builder"></span> NumberLiteralBuilder

## Constructor

```go
func NewNumberLiteralBuilder() *NumberLiteralBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NumberLiteralBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> Value

```go
func (builder *NumberLiteralBuilder) Value(value float64) *NumberLiteralBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
