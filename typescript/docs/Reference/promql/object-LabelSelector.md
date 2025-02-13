---
title: <span class="badge object-type-interface"></span> LabelSelector
---
# <span class="badge object-type-interface"></span> LabelSelector

## Definition

```typescript
export interface LabelSelector {
	// Name of the label to select.
	name: string;
	// Value to match against.
	value: string;
	// Operator used to perform the selection.
	operator: promql.LabelMatchingOperator;
}

```
## See also

 * <span class="badge builder"></span> [LabelSelectorBuilder](./builder-LabelSelectorBuilder.md)
