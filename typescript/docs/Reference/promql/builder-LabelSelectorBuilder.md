---
title: <span class="badge builder"></span> LabelSelectorBuilder
---
# <span class="badge builder"></span> LabelSelectorBuilder

## Constructor

```typescript
new LabelSelectorBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> name

Name of the label to select.

```typescript
name(name: string)
```

### <span class="badge object-method"></span> operator

Operator used to perform the selection.

```typescript
operator(operator: promql.LabelMatchingOperator)
```

### <span class="badge object-method"></span> value

Value to match against.

```typescript
value(value: string)
```

## See also

 * <span class="badge object-type-interface"></span> [LabelSelector](./object-LabelSelector.md)
