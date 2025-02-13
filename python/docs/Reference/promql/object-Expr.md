---
title: <span class="badge object-type-disjunction"></span> Expr
---
# <span class="badge object-type-disjunction"></span> Expr

Represents a PromQL expression.

## Definition

```python
# Represents a PromQL expression.
Expr: typing.TypeAlias = typing.Union[promql.NumberLiteralExpr, promql.StringLiteralExpr, promql.SubqueryExpr, promql.AggregationExpr, promql.VectorExpr, promql.BinaryExpr, promql.UnaryExpr, promql.FuncCallExpr]
```
