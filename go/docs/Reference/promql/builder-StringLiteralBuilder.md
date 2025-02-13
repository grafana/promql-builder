---
title: <span class="badge builder"></span> StringLiteralBuilder
---
# <span class="badge builder"></span> StringLiteralBuilder

## Constructor

```go
func NewStringLiteralBuilder() *StringLiteralBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StringLiteralBuilder) Build() (Expr, error)
```

### <span class="badge object-method"></span> Value

```go
func (builder *StringLiteralBuilder) Value(value string) *StringLiteralBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
