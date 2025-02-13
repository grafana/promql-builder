package promql

import "strings"

// Represents a PromQL expression.
Expr: #NumberLiteralExpr | #StringLiteralExpr | #SubqueryExpr | #AggregationExpr | #VectorExpr | #BinaryExpr | #UnaryExpr | #FuncCallExpr

// Represents a number literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
#NumberLiteralExpr: {
	type: "numberLiteralExpr"
	value: number
}

// Represents a string literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
#StringLiteralExpr: {
	type: "stringLiteralExpr"
	value: string
}

// Represents a subquery.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
#SubqueryExpr: {
	type: "subqueryExpr"
	expr: Expr

	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	offset: string

	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	at: string

	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	range: string

	// Empty string for default resolution.
	resolution?: string 
}

// Represents an aggregation.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
#AggregationExpr: {
	type: "aggregationExpr"
	op: #AggregationOp
	expr: Expr
	param?: Expr

	// By drops labels that are not listed in the by clause.
	by: [...string]

	// List of labels to remove from the result vector, while all other labels are preserved in the output.
	without: [...string]
}

// Possible aggregation operators.
#AggregationOp: "sum" | "min" | "max" | "avg" | "stddev" | "stdvar" | "count" | "group" | "count_values" | "bottomk" | "topk" | "quantile" | "limitk" | "limit_ratio" @cog(kind="enum", memberNames="sum|min|max|avg|stddev|stdvar|count|group|count_values|bottomk|topk|quantile|limitk|limit_ratio")

// Represents both instant and range vectors
#VectorExpr: {
	type: "vectorExpr"

	// Metric name.
	metric: string

	// Label selectors used to filter the timeseries.
	labels: [...#LabelSelector]

	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	offset: string
	offset: string

	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	at: string

	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	range: string
}

// Represents an unary operation expression.
#UnaryExpr: {
	type: "unaryExpr"
	op: #UnaryOp
	expr: Expr
}

// Possible unary operators.
#UnaryOp: "+" | "-" @cog(kind="enum", memberNames="plus|minus")

// Represents a binary operation expression.
#BinaryExpr: {
	type: "binaryExpr"
	op: #BinaryOp
	left: Expr
	right: Expr
	// https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
	matchType?: "on" | "ignore" @cog(kind="enum", memberNames="on|ignore")
	matchLabels?: [...string]
	groupModifier?: "left" | "right" @cog(kind="enum", memberNames="left|right")
	groupLabels?: [...string]
}

// Possible binary operators.
#BinaryOp: "+" |
  "-" |
  "*" |
  "/" |
  "%" |
  "^" |
  "==" |
  "!=" |
  ">" |
  "<" |
  ">=" |
  "<=" |
  "and" |
  "or" |
  "unless" |
  "atan2" @cog(kind="enum", memberNames="add|sub|mul|div|mod|pow|eql|neq|gtr|lss|gte|lte|and|or|unless|atan2")

// Represents a function call expression.
#FuncCallExpr: {
	type: "funcCallExpr"

	// Name of the function.
	function: string & strings.MinRunes(1)

	// Arguments.
	args: [...Expr]
}

// Possible label matching operators.
#LabelMatchingOperator: "=" | "!=" | "=~" | "!~" @cog(kind="enum", memberNames="equal|notEqual|matchRegexp|notMatchRegexp")

#LabelSelector: {
	// Name of the label to select.
	name: string & strings.MinRunes(1)
	// Value to match against.
	value: string
	// Operator used to perform the selection.
	operator: #LabelMatchingOperator
}
