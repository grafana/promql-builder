---
title: <span class="badge object-type-struct"></span> BinaryExpr
---
# <span class="badge object-type-struct"></span> BinaryExpr

Represents a binary operation expression.

## Definition

```go
type BinaryExpr struct {
    Type string `json:"type"`
    Op promql.BinaryOp `json:"op"`
    Left promql.Expr `json:"left"`
    Right promql.Expr `json:"right"`
    // https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
    MatchType *promql.BinaryExprMatchType `json:"matchType,omitempty"`
    MatchLabels []string `json:"matchLabels,omitempty"`
    GroupModifier *promql.BinaryExprGroupModifier `json:"groupModifier,omitempty"`
    GroupLabels []string `json:"groupLabels,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BinaryExpr` fields for violations and returns them.

```go
func (binaryExpr *BinaryExpr) Validate() error
```

