---
title: <span class="badge object-type-struct"></span> NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr
---
# <span class="badge object-type-struct"></span> NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr

## Definition

```go
type NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr struct {
    NumberLiteralExpr *promql.NumberLiteralExpr `json:"NumberLiteralExpr,omitempty"`
    StringLiteralExpr *promql.StringLiteralExpr `json:"StringLiteralExpr,omitempty"`
    SubqueryExpr *promql.SubqueryExpr `json:"SubqueryExpr,omitempty"`
    AggregationExpr *promql.AggregationExpr `json:"AggregationExpr,omitempty"`
    VectorExpr *promql.VectorExpr `json:"VectorExpr,omitempty"`
    BinaryExpr *promql.BinaryExpr `json:"BinaryExpr,omitempty"`
    UnaryExpr *promql.UnaryExpr `json:"UnaryExpr,omitempty"`
    FuncCallExpr *promql.FuncCallExpr `json:"FuncCallExpr,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr` fields for violations and returns them.

```go
func (numberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr *NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr) Validate() error
```

