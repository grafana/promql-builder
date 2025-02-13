// Code generated - EDITING IS FUTILE. DO NOT EDIT.

// Represents a PromQL expression.
export type Expr = NumberLiteralExpr | StringLiteralExpr | SubqueryExpr | AggregationExpr | VectorExpr | BinaryExpr | UnaryExpr | FuncCallExpr;

export const defaultExpr = (): Expr => (defaultNumberLiteralExpr());

export const toString = (expr: Expr): string => {
	switch (expr.type) {
		case "aggregationExpr":
			return aggregationExprToString(expr);
		case "binaryExpr":
			return binaryExprToString(expr);
		case "funcCallExpr":
			return funcCallExprToString(expr);
		case "numberLiteralExpr":
			return numberLiteralExprToString(expr);
		case "stringLiteralExpr":
			return stringLiteralExprToString(expr);
		case "subqueryExpr":
			return subqueryExprToString(expr);
		case "unaryExpr":
			return unaryExprToString(expr);
		case "vectorExpr":
			return vectorExprToString(expr);
		default:
			return 'unknown expression type';
	}
};

// Represents a number literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
export interface NumberLiteralExpr {
	type: "numberLiteralExpr";
	value: number;
}

export const defaultNumberLiteralExpr = (): NumberLiteralExpr => ({
	type: "numberLiteralExpr",
	value: 0,
});

const numberLiteralExprToString = (expr: NumberLiteralExpr): string => {
    return `${JSON.stringify(expr.value)}`;
};

// Represents a string literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
export interface StringLiteralExpr {
	type: "stringLiteralExpr";
	value: string;
}

export const defaultStringLiteralExpr = (): StringLiteralExpr => ({
	type: "stringLiteralExpr",
	value: "",
});

const stringLiteralExprToString = (expr: StringLiteralExpr): string => {
    return expr.value;
};

// Represents a subquery.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
export interface SubqueryExpr {
	type: "subqueryExpr";
	expr: Expr;
	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	offset: string;
	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	at: string;
	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	range: string;
	// Empty string for default resolution.
	resolution?: string;
}

export const defaultSubqueryExpr = (): SubqueryExpr => ({
	type: "subqueryExpr",
	expr: defaultExpr(),
	offset: "",
	at: "",
	range: "",
});

const subqueryExprToString = (expr: SubqueryExpr): string => {
    let buffer = `(${toString(expr.expr)})`;

    if (expr.range !== '') {
        buffer += `[${expr.range}`;

        if (expr.resolution) {
            buffer += `:${expr.resolution}`;
        }

        buffer += ']';
    }

    if (expr.offset !== '') {
        buffer += ` offset ${expr.offset}`;
    }

    if (expr.at !== '') {
        buffer += ` @ ${expr.at}`;
    }

    return buffer;
};

// Represents an aggregation.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
export interface AggregationExpr {
	type: "aggregationExpr";
	op: AggregationOp;
	expr: Expr;
	param?: Expr;
	// By drops labels that are not listed in the by clause.
	by: string[];
	// List of labels to remove from the result vector, while all other labels are preserved in the output.
	without: string[];
}

export const defaultAggregationExpr = (): AggregationExpr => ({
	type: "aggregationExpr",
	op: "sum",
	expr: defaultExpr(),
	by: [],
	without: [],
});

const aggregationExprToString = (expr: AggregationExpr): string => {
    let buffer = `${expr.op}`;

    if (expr.by.length !== 0) {
        buffer += ` by(${expr.by.join(', ')}) `;
    }

    if (expr.without.length !== 0) {
        buffer += ` without(${expr.without.join(', ')}) `;
    }

    buffer += '(';

    if (expr.param) {
        buffer += `${toString(expr.param)}, `;
    }

    buffer += toString(expr.expr);
    buffer += ')';

    return buffer;
};

// Possible aggregation operators.
export type AggregationOp = "sum" | "min" | "max" | "avg" | "stddev" | "stdvar" | "count" | "group" | "count_values" | "bottomk" | "topk" | "quantile" | "limitk" | "limit_ratio";

export const defaultAggregationOp = (): AggregationOp => ("sum");

// Represents both instant and range vectors
export interface VectorExpr {
	type: "vectorExpr";
	// Metric name.
	metric: string;
	// Label selectors used to filter the timeseries.
	labels: LabelSelector[];
	// The offset modifier allows changing the time offset for individual instant and range vectors in a query.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
	offset: string;
	// The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
	// The time supplied to the @ modifier is a unix timestamp.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
	at: string;
	// Range of samples back from the current instant.
	// https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
	range: string;
}

export const defaultVectorExpr = (): VectorExpr => ({
	type: "vectorExpr",
	metric: "",
	labels: [],
	offset: "",
	at: "",
	range: "",
});

const vectorExprToString = (expr: VectorExpr): string => {
    let buffer = expr.metric;

    if (expr.labels.length !== 0) {
        const labels = expr.labels.map(labelToString);
        buffer += `{${labels.join(',')}}`;
    }

    if (expr.range !== '') {
        buffer += `[${expr.range}]`;
    }

    if (expr.offset !== '') {
        buffer += ` offset ${expr.offset}`;
    }

    if (expr.at !== '') {
        buffer += ` @ ${expr.at}`;
    }

    return buffer;
};

export interface LabelSelector {
	// Name of the label to select.
	name: string;
	// Value to match against.
	value: string;
	// Operator used to perform the selection.
	operator: LabelMatchingOperator;
}

export const defaultLabelSelector = (): LabelSelector => ({
	name: "",
	value: "",
	operator: "=",
});

export const labelToString = (label: LabelSelector): string => {
    return `${label.name}${label.operator}${JSON.stringify(label.value)}`;
};

// Possible label matching operators.
export type LabelMatchingOperator = "=" | "!=" | "=~" | "!~";

export const defaultLabelMatchingOperator = (): LabelMatchingOperator => ("=");

// Represents a binary operation expression.
export interface BinaryExpr {
	type: "binaryExpr";
	op: BinaryOp;
	left: Expr;
	right: Expr;
	// https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
	matchType?: "on" | "ignore";
	matchLabels?: string[];
	groupModifier?: "left" | "right";
	groupLabels?: string[];
}

export const defaultBinaryExpr = (): BinaryExpr => ({
	type: "binaryExpr",
	op: "+",
	left: defaultExpr(),
	right: defaultExpr(),
});

const binaryExprToString = (expr: BinaryExpr): string => {
    let buffer = '';

    buffer += `(${toString(expr.left)})`;
    buffer += ` ${expr.op} `;

    if (expr.matchType) {
        buffer += expr.matchType === 'on' ? 'on' : 'ignoring';

        if (expr.matchLabels && expr.matchLabels.length !== 0) {
            buffer += `(${expr.matchLabels.join(', ')}) `;
        }
    }

    if (expr.groupModifier) {
        buffer += expr.groupModifier === 'left' ? 'group_left' : 'group_right';

        if (expr.groupLabels && expr.groupLabels.length !== 0) {
            buffer += `(${expr.groupLabels.join(', ')}) `;
        }
    }

    buffer += `(${toString(expr.right)})`;

    return buffer;
};

// Possible binary operators.
export type BinaryOp = "+" | "-" | "*" | "/" | "%" | "^" | "==" | "!=" | ">" | "<" | ">=" | "<=" | "and" | "or" | "unless" | "atan2";

export const defaultBinaryOp = (): BinaryOp => ("+");

// Represents an unary operation expression.
export interface UnaryExpr {
	type: "unaryExpr";
	op: UnaryOp;
	expr: Expr;
}

export const defaultUnaryExpr = (): UnaryExpr => ({
	type: "unaryExpr",
	op: "+",
	expr: defaultExpr(),
});

const unaryExprToString = (expr: UnaryExpr): string => {
    return `${expr.op}${expr.expr.toString()}`;
};

// Possible unary operators.
export type UnaryOp = "+" | "-";

export const defaultUnaryOp = (): UnaryOp => ("+");

// Represents a function call expression.
export interface FuncCallExpr {
	type: "funcCallExpr";
	// Name of the function.
	function: string;
	// Arguments.
	args: Expr[];
}

export const defaultFuncCallExpr = (): FuncCallExpr => ({
	type: "funcCallExpr",
	function: "",
	args: [],
});

const funcCallExprToString = (expr: FuncCallExpr): string => {
    let buffer = `${expr.function}(`;
    const args = expr.args.map(toString);
    buffer += `${args.join(', ')}`;
    buffer += ')';

    return buffer;
};

