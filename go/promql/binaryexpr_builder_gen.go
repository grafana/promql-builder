// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*BinaryExprBuilder)(nil)

// Represents a PromQL expression.
type BinaryExprBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewBinaryExprBuilder() *BinaryExprBuilder {
	resource := NewExpr()
	builder := &BinaryExprBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	builder.internal.BinaryExpr.Type = "binaryExpr"

	return builder
}

// Addition binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
func Add(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpAdd)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// Subtraction binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
func Sub(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpSub)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// Multiplication binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
func Mul(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpMul)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// Division binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
func Div(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpDiv)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// Modulo binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
func Mod(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpMod)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// Power/exponentiation binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
func Pow(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpPow)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "equal" comparison binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
func Eq(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpEql)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "not-equal" comparison binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
func Neq(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpNeq)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "greater-than" comparison binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
func Gt(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpGtr)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "less-than" comparison binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
func Lt(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpLss)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "greater-or-equal" comparison binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
func Gte(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpGte)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "less-or-equal" comparison binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
func Lte(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpLte)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "intersection" logical/set binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
func And(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpAnd)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "union" logical/set binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
func Or(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpOr)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// "complement" logical/set binary operator.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
func Unless(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpUnless)
	builder.Left(left)
	builder.Right(right)

	return builder
}

// Arc tangent binary operator. Works in radians.
// Trigonometric operators allow trigonometric functions to be executed on two vectors using vector matching, which isn't available with normal functions.
// They act in the same manner as arithmetic operators.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#trigonometric-binary-operators
func Atan2(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder {
	builder := NewBinaryExprBuilder()
	builder.Op(BinaryOpAtan2)
	builder.Left(left)
	builder.Right(right)

	return builder
}

func (builder *BinaryExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	return *builder.internal, nil
}

func (builder BinaryExprBuilder) String() string {
	return builder.internal.String()
}

func (builder *BinaryExprBuilder) Op(op BinaryOp) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	builder.internal.BinaryExpr.Op = op

	return builder
}

func (builder *BinaryExprBuilder) Left(left cog.Builder[Expr]) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	leftResource, err := left.Build()
	if err != nil {
		builder.errors["BinaryExpr.left"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.BinaryExpr.Left = leftResource

	return builder
}

func (builder *BinaryExprBuilder) Right(right cog.Builder[Expr]) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	rightResource, err := right.Build()
	if err != nil {
		builder.errors["BinaryExpr.right"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.BinaryExpr.Right = rightResource

	return builder
}

// Allows ignoring certain labels when matching.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches
func (builder *BinaryExprBuilder) Ignoring(labels []string) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	valMatchType := BinaryExprMatchTypeIgnore
	builder.internal.BinaryExpr.MatchType = &valMatchType
	builder.internal.BinaryExpr.MatchLabels = labels

	return builder
}

// Allows reducing the set of considered labels to a provided list when matching.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches
func (builder *BinaryExprBuilder) On(labels []string) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	valMatchType := BinaryExprMatchTypeOn
	builder.internal.BinaryExpr.MatchType = &valMatchType
	builder.internal.BinaryExpr.MatchLabels = labels

	return builder
}

// See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches
func (builder *BinaryExprBuilder) GroupLeft(labels []string) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	valGroupModifier := BinaryExprGroupModifierLeft
	builder.internal.BinaryExpr.GroupModifier = &valGroupModifier
	builder.internal.BinaryExpr.GroupLabels = labels

	return builder
}

// See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches
func (builder *BinaryExprBuilder) GroupRight(labels []string) *BinaryExprBuilder {
	if builder.internal.BinaryExpr == nil {
		builder.internal.BinaryExpr = NewBinaryExpr()
	}
	valGroupModifier := BinaryExprGroupModifierRight
	builder.internal.BinaryExpr.GroupModifier = &valGroupModifier
	builder.internal.BinaryExpr.GroupLabels = labels

	return builder
}
