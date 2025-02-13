---
title: <span class="badge object-type-enum"></span> BinaryOp
---
# <span class="badge object-type-enum"></span> BinaryOp

Possible binary operators.

## Definition

```go
type BinaryOp string
const (
	BinaryOpAdd BinaryOp = "+"
	BinaryOpSub BinaryOp = "-"
	BinaryOpMul BinaryOp = "*"
	BinaryOpDiv BinaryOp = "/"
	BinaryOpMod BinaryOp = "%"
	BinaryOpPow BinaryOp = "^"
	BinaryOpEql BinaryOp = "=="
	BinaryOpNeq BinaryOp = "!="
	BinaryOpGtr BinaryOp = ">"
	BinaryOpLss BinaryOp = "<"
	BinaryOpGte BinaryOp = ">="
	BinaryOpLte BinaryOp = "<="
	BinaryOpAnd BinaryOp = "and"
	BinaryOpOr BinaryOp = "or"
	BinaryOpUnless BinaryOp = "unless"
	BinaryOpAtan2 BinaryOp = "atan2"
)

```
