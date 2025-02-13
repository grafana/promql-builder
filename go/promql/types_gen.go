// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	cog "github.com/grafana/promql-builder/go/cog"
)

// Represents a PromQL expression.
type Expr = NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr

// NewExpr creates a new Expr object.
func NewExpr() *Expr {
	return NewNumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr()
}

func (expr Expr) String() string {
	if expr.NumberLiteralExpr != nil {
		return expr.NumberLiteralExpr.String()
	}
	if expr.StringLiteralExpr != nil {
		return expr.StringLiteralExpr.String()
	}
	if expr.SubqueryExpr != nil {
		return expr.SubqueryExpr.String()
	}
	if expr.AggregationExpr != nil {
		return expr.AggregationExpr.String()
	}
	if expr.VectorExpr != nil {
		return expr.VectorExpr.String()
	}
	if expr.BinaryExpr != nil {
		return expr.BinaryExpr.String()
	}
	if expr.UnaryExpr != nil {
		return expr.UnaryExpr.String()
	}
	if expr.FuncCallExpr != nil {
		return expr.FuncCallExpr.String()
	}

	return ""
}

// Represents a number literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
type NumberLiteralExpr struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

// NewNumberLiteralExpr creates a new NumberLiteralExpr object.
func NewNumberLiteralExpr() *NumberLiteralExpr {
	return &NumberLiteralExpr{
		Type: "numberLiteralExpr",
	}
}

// Validate checks all the validation constraints that may be defined on `NumberLiteralExpr` fields for violations and returns them.
func (resource NumberLiteralExpr) Validate() error {
	return nil
}

func (expr NumberLiteralExpr) String() string {
	return fmt.Sprintf("%v", expr.Value)
}

// Represents a string literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
type StringLiteralExpr struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// NewStringLiteralExpr creates a new StringLiteralExpr object.
func NewStringLiteralExpr() *StringLiteralExpr {
	return &StringLiteralExpr{
		Type: "stringLiteralExpr",
	}
}

// Validate checks all the validation constraints that may be defined on `StringLiteralExpr` fields for violations and returns them.
func (resource StringLiteralExpr) Validate() error {
	return nil
}

func (expr StringLiteralExpr) String() string {
	return expr.Value
}

// Represents a subquery.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
type SubqueryExpr struct {
	Type string `json:"type"`
	Expr Expr   `json:"expr"`
	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	Offset string `json:"offset"`
	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	At string `json:"at"`
	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	Range string `json:"range"`
	// Empty string for default resolution.
	Resolution *string `json:"resolution,omitempty"`
}

// NewSubqueryExpr creates a new SubqueryExpr object.
func NewSubqueryExpr() *SubqueryExpr {
	return &SubqueryExpr{
		Type: "subqueryExpr",
		Expr: *NewExpr(),
	}
}

// Validate checks all the validation constraints that may be defined on `SubqueryExpr` fields for violations and returns them.
func (resource SubqueryExpr) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Expr.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("expr", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (expression SubqueryExpr) String() string {
	var buffer strings.Builder

	buffer.WriteString("(")
	buffer.WriteString(expression.Expr.String())
	buffer.WriteString(")")

	if expression.Range != "" {
		buffer.WriteString("[")
		buffer.WriteString(expression.Range)
		if expression.Resolution != nil {
			buffer.WriteString(":")
			buffer.WriteString(*expression.Resolution)
		}
		buffer.WriteString("]")
	}

	if expression.Offset != "" {
		buffer.WriteString(" offset ")
		buffer.WriteString(expression.Offset)
	}

	if expression.At != "" {
		buffer.WriteString(" @ ")
		buffer.WriteString(expression.At)
	}

	return buffer.String()
}

// Represents an aggregation.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
type AggregationExpr struct {
	Type  string        `json:"type"`
	Op    AggregationOp `json:"op"`
	Expr  Expr          `json:"expr"`
	Param *Expr         `json:"param,omitempty"`
	// By drops labels that are not listed in the by clause.
	By []string `json:"by"`
	// List of labels to remove from the result vector, while all other labels are preserved in the output.
	Without []string `json:"without"`
}

// NewAggregationExpr creates a new AggregationExpr object.
func NewAggregationExpr() *AggregationExpr {
	return &AggregationExpr{
		Type: "aggregationExpr",
		Expr: *NewExpr(),
	}
}

// Validate checks all the validation constraints that may be defined on `AggregationExpr` fields for violations and returns them.
func (resource AggregationExpr) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Expr.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("expr", err)...)
	}
	if resource.Param != nil {
		if err := resource.Param.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("param", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (agg AggregationExpr) String() string {
	var buffer strings.Builder

	buffer.WriteString(string(agg.Op))

	if len(agg.Without) != 0 {
		buffer.WriteString(" without")
		buffer.WriteString(" (")
		buffer.WriteString(strings.Join(agg.Without, ", "))
		buffer.WriteString(") ")
	}

	if len(agg.By) != 0 {
		buffer.WriteString(" by")
		buffer.WriteString(" (")
		buffer.WriteString(strings.Join(agg.By, ", "))
		buffer.WriteString(") ")
	}

	buffer.WriteString("(")
	if agg.Param != nil {
		buffer.WriteString(agg.Param.String())
		buffer.WriteString(", ")
	}

	buffer.WriteString(agg.Expr.String())
	buffer.WriteString(")")

	return buffer.String()
}

// Possible aggregation operators.
type AggregationOp string

const (
	AggregationOpSum         AggregationOp = "sum"
	AggregationOpMin         AggregationOp = "min"
	AggregationOpMax         AggregationOp = "max"
	AggregationOpAvg         AggregationOp = "avg"
	AggregationOpStddev      AggregationOp = "stddev"
	AggregationOpStdvar      AggregationOp = "stdvar"
	AggregationOpCount       AggregationOp = "count"
	AggregationOpGroup       AggregationOp = "group"
	AggregationOpCountValues AggregationOp = "count_values"
	AggregationOpBottomk     AggregationOp = "bottomk"
	AggregationOpTopk        AggregationOp = "topk"
	AggregationOpQuantile    AggregationOp = "quantile"
	AggregationOpLimitk      AggregationOp = "limitk"
	AggregationOpLimitRatio  AggregationOp = "limit_ratio"
)

// Represents both instant and range vectors
type VectorExpr struct {
	Type string `json:"type"`
	// Metric name.
	Metric string `json:"metric"`
	// Label selectors used to filter the timeseries.
	Labels []LabelSelector `json:"labels"`
	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	Offset string `json:"offset"`
	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	At string `json:"at"`
	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	Range string `json:"range"`
}

// NewVectorExpr creates a new VectorExpr object.
func NewVectorExpr() *VectorExpr {
	return &VectorExpr{
		Type: "vectorExpr",
	}
}

// Validate checks all the validation constraints that may be defined on `VectorExpr` fields for violations and returns them.
func (resource VectorExpr) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Labels {
		if err := resource.Labels[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("labels["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (expression VectorExpr) String() string {
	var buffer strings.Builder

	buffer.WriteString(expression.Metric)

	if len(expression.Labels) != 0 {
		buffer.WriteString("{")
		for i, label := range expression.Labels {
			buffer.WriteString(label.String())

			if i != len(expression.Labels)-1 {
				buffer.WriteString(",")
			}
		}
		buffer.WriteString("}")
	}

	if expression.Range != "" {
		buffer.WriteString("[")
		buffer.WriteString(expression.Range)
		buffer.WriteString("]")
	}

	if expression.Offset != "" {
		buffer.WriteString(" offset ")
		buffer.WriteString(expression.Offset)
	}

	if expression.At != "" {
		buffer.WriteString(" @ ")
		buffer.WriteString(expression.At)
	}

	return buffer.String()
}

type LabelSelector struct {
	// Name of the label to select.
	Name string `json:"name"`
	// Value to match against.
	Value string `json:"value"`
	// Operator used to perform the selection.
	Operator LabelMatchingOperator `json:"operator"`
}

// NewLabelSelector creates a new LabelSelector object.
func NewLabelSelector() *LabelSelector {
	return &LabelSelector{}
}

// Validate checks all the validation constraints that may be defined on `LabelSelector` fields for violations and returns them.
func (resource LabelSelector) Validate() error {
	var errs cog.BuildErrors
	if !(len([]rune(resource.Name)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"name",
			errors.New("must be >= 1"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (label LabelSelector) String() string {
	return fmt.Sprintf("%s%s\"%s\"", label.Name, label.Operator, label.Value)
}

// Possible label matching operators.
type LabelMatchingOperator string

const (
	LabelMatchingOperatorEqual          LabelMatchingOperator = "="
	LabelMatchingOperatorNotEqual       LabelMatchingOperator = "!="
	LabelMatchingOperatorMatchRegexp    LabelMatchingOperator = "=~"
	LabelMatchingOperatorNotMatchRegexp LabelMatchingOperator = "!~"
)

// Represents a binary operation expression.
type BinaryExpr struct {
	Type  string   `json:"type"`
	Op    BinaryOp `json:"op"`
	Left  Expr     `json:"left"`
	Right Expr     `json:"right"`
	// https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
	MatchType     *BinaryExprMatchType     `json:"matchType,omitempty"`
	MatchLabels   []string                 `json:"matchLabels,omitempty"`
	GroupModifier *BinaryExprGroupModifier `json:"groupModifier,omitempty"`
	GroupLabels   []string                 `json:"groupLabels,omitempty"`
}

// NewBinaryExpr creates a new BinaryExpr object.
func NewBinaryExpr() *BinaryExpr {
	return &BinaryExpr{
		Type:  "binaryExpr",
		Left:  *NewExpr(),
		Right: *NewExpr(),
	}
}

// Validate checks all the validation constraints that may be defined on `BinaryExpr` fields for violations and returns them.
func (resource BinaryExpr) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Left.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("left", err)...)
	}
	if err := resource.Right.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("right", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (expr BinaryExpr) String() string {
	var buffer strings.Builder

	buffer.WriteString("(")
	buffer.WriteString(expr.Left.String())
	buffer.WriteString(")")

	buffer.WriteString(" ")
	buffer.WriteString(string(expr.Op))
	buffer.WriteString(" ")

	if expr.MatchType != nil {
		if *expr.MatchType == "on" {
			buffer.WriteString("on")
		} else {
			buffer.WriteString("ignoring")
		}

		if len(expr.MatchLabels) != 0 {
			buffer.WriteString("(")
			buffer.WriteString(strings.Join(expr.MatchLabels, ", "))
			buffer.WriteString(") ")
		}
	}

	if expr.GroupModifier != nil {
		buffer.WriteString(" ")

		if *expr.GroupModifier == "left" {
			buffer.WriteString("group_left")
		} else {
			buffer.WriteString("group_right")
		}

		if len(expr.GroupLabels) != 0 {
			buffer.WriteString("(")
			buffer.WriteString(strings.Join(expr.GroupLabels, ", "))
			buffer.WriteString(") ")
		}
	}

	buffer.WriteString("(")
	buffer.WriteString(expr.Right.String())
	buffer.WriteString(")")

	return buffer.String()
}

// Possible binary operators.
type BinaryOp string

const (
	BinaryOpAdd    BinaryOp = "+"
	BinaryOpSub    BinaryOp = "-"
	BinaryOpMul    BinaryOp = "*"
	BinaryOpDiv    BinaryOp = "/"
	BinaryOpMod    BinaryOp = "%"
	BinaryOpPow    BinaryOp = "^"
	BinaryOpEql    BinaryOp = "=="
	BinaryOpNeq    BinaryOp = "!="
	BinaryOpGtr    BinaryOp = ">"
	BinaryOpLss    BinaryOp = "<"
	BinaryOpGte    BinaryOp = ">="
	BinaryOpLte    BinaryOp = "<="
	BinaryOpAnd    BinaryOp = "and"
	BinaryOpOr     BinaryOp = "or"
	BinaryOpUnless BinaryOp = "unless"
	BinaryOpAtan2  BinaryOp = "atan2"
)

// Represents an unary operation expression.
type UnaryExpr struct {
	Type string  `json:"type"`
	Op   UnaryOp `json:"op"`
	Expr Expr    `json:"expr"`
}

// NewUnaryExpr creates a new UnaryExpr object.
func NewUnaryExpr() *UnaryExpr {
	return &UnaryExpr{
		Type: "unaryExpr",
		Expr: *NewExpr(),
	}
}

// Validate checks all the validation constraints that may be defined on `UnaryExpr` fields for violations and returns them.
func (resource UnaryExpr) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Expr.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("expr", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (expr UnaryExpr) String() string {
	return string(expr.Op) + expr.Expr.String()
}

// Possible unary operators.
type UnaryOp string

const (
	UnaryOpPlus  UnaryOp = "+"
	UnaryOpMinus UnaryOp = "-"
)

// Represents a function call expression.
type FuncCallExpr struct {
	Type string `json:"type"`
	// Name of the function.
	Function string `json:"function"`
	// Arguments.
	Args []Expr `json:"args"`
}

// NewFuncCallExpr creates a new FuncCallExpr object.
func NewFuncCallExpr() *FuncCallExpr {
	return &FuncCallExpr{
		Type: "funcCallExpr",
	}
}

// Validate checks all the validation constraints that may be defined on `FuncCallExpr` fields for violations and returns them.
func (resource FuncCallExpr) Validate() error {
	var errs cog.BuildErrors
	if !(len([]rune(resource.Function)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"function",
			errors.New("must be >= 1"),
		)...)
	}

	for i1 := range resource.Args {
		if err := resource.Args[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("args["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (expr FuncCallExpr) String() string {
	var buffer strings.Builder

	buffer.WriteString(expr.Function)
	buffer.WriteString("(")
	for i, arg := range expr.Args {
		buffer.WriteString(arg.String())
		if i != len(expr.Args)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString(")")

	return buffer.String()
}

type BinaryExprMatchType string

const (
	BinaryExprMatchTypeOn     BinaryExprMatchType = "on"
	BinaryExprMatchTypeIgnore BinaryExprMatchType = "ignore"
)

type BinaryExprGroupModifier string

const (
	BinaryExprGroupModifierLeft  BinaryExprGroupModifier = "left"
	BinaryExprGroupModifierRight BinaryExprGroupModifier = "right"
)

type NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr struct {
	NumberLiteralExpr *NumberLiteralExpr `json:"NumberLiteralExpr,omitempty"`
	StringLiteralExpr *StringLiteralExpr `json:"StringLiteralExpr,omitempty"`
	SubqueryExpr      *SubqueryExpr      `json:"SubqueryExpr,omitempty"`
	AggregationExpr   *AggregationExpr   `json:"AggregationExpr,omitempty"`
	VectorExpr        *VectorExpr        `json:"VectorExpr,omitempty"`
	BinaryExpr        *BinaryExpr        `json:"BinaryExpr,omitempty"`
	UnaryExpr         *UnaryExpr         `json:"UnaryExpr,omitempty"`
	FuncCallExpr      *FuncCallExpr      `json:"FuncCallExpr,omitempty"`
}

// NewNumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr creates a new NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr object.
func NewNumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr() *NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr {
	return &NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr{}
}

// Validate checks all the validation constraints that may be defined on `NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr` fields for violations and returns them.
func (resource NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr) Validate() error {
	var errs cog.BuildErrors
	if resource.NumberLiteralExpr != nil {
		if err := resource.NumberLiteralExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("NumberLiteralExpr", err)...)
		}
	}
	if resource.StringLiteralExpr != nil {
		if err := resource.StringLiteralExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("StringLiteralExpr", err)...)
		}
	}
	if resource.SubqueryExpr != nil {
		if err := resource.SubqueryExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("SubqueryExpr", err)...)
		}
	}
	if resource.AggregationExpr != nil {
		if err := resource.AggregationExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("AggregationExpr", err)...)
		}
	}
	if resource.VectorExpr != nil {
		if err := resource.VectorExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("VectorExpr", err)...)
		}
	}
	if resource.BinaryExpr != nil {
		if err := resource.BinaryExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("BinaryExpr", err)...)
		}
	}
	if resource.UnaryExpr != nil {
		if err := resource.UnaryExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("UnaryExpr", err)...)
		}
	}
	if resource.FuncCallExpr != nil {
		if err := resource.FuncCallExpr.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("FuncCallExpr", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
