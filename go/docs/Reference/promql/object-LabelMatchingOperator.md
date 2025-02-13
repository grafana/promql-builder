---
title: <span class="badge object-type-enum"></span> LabelMatchingOperator
---
# <span class="badge object-type-enum"></span> LabelMatchingOperator

Possible label matching operators.

## Definition

```go
type LabelMatchingOperator string
const (
	LabelMatchingOperatorEqual LabelMatchingOperator = "="
	LabelMatchingOperatorNotEqual LabelMatchingOperator = "!="
	LabelMatchingOperatorMatchRegexp LabelMatchingOperator = "=~"
	LabelMatchingOperatorNotMatchRegexp LabelMatchingOperator = "!~"
)

```
