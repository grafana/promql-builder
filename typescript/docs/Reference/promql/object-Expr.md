---
title: <span class="badge object-type-disjunction"></span> Expr
---
# <span class="badge object-type-disjunction"></span> Expr

Represents a PromQL expression.

## Definition

```typescript
export type Expr = promql.NumberLiteralExpr | promql.StringLiteralExpr | promql.SubqueryExpr | promql.AggregationExpr | promql.VectorExpr | promql.BinaryExpr | promql.UnaryExpr | promql.FuncCallExpr;

```
