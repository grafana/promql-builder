// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents a binary operation expression.
export class BinaryExprBuilder implements cog.Builder<promql.BinaryExpr> {
    protected readonly internal: promql.BinaryExpr;

    constructor() {
        this.internal = promql.defaultBinaryExpr();
        this.internal.type = "binaryExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.BinaryExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    op(op: promql.BinaryOp): this {
        this.internal.op = op;
        return this;
    }

    left(left: cog.Builder<promql.Expr>): this {
        const leftResource = left.build();
        this.internal.left = leftResource;
        return this;
    }

    right(right: cog.Builder<promql.Expr>): this {
        const rightResource = right.build();
        this.internal.right = rightResource;
        return this;
    }

    // Allows ignoring certain labels when matching.
    // See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches
    ignoring(labels: string[]): this {
        this.internal.matchType = "ignore";
        this.internal.matchLabels = labels;
        return this;
    }

    // Allows reducing the set of considered labels to a provided list when matching.
    // See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches
    on(labels: string[]): this {
        this.internal.matchType = "on";
        this.internal.matchLabels = labels;
        return this;
    }

    // See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches
    groupLeft(labels: string[]): this {
        this.internal.groupModifier = "left";
        this.internal.groupLabels = labels;
        return this;
    }

    // See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches
    groupRight(labels: string[]): this {
        this.internal.groupModifier = "right";
        this.internal.groupLabels = labels;
        return this;
    }
}

/**
 * Addition binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
 */
export function add(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("+");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * Subtraction binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
 */
export function sub(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("-");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * Multiplication binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
 */
export function mul(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("*");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * Division binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
 */
export function div(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("/");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * Modulo binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
 */
export function mod(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("%");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * Power/exponentiation binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
 */
export function pow(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("^");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "equal" comparison binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
 */
export function eq(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("==");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "not-equal" comparison binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
 */
export function neq(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("!=");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "greater-than" comparison binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
 */
export function gt(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op(">");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "less-than" comparison binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
 */
export function lt(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("<");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "greater-or-equal" comparison binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
 */
export function gte(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op(">=");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "less-or-equal" comparison binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
 */
export function lte(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("<=");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "intersection" logical/set binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
 */
export function and(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("and");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "union" logical/set binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
 */
export function or(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("or");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * "complement" logical/set binary operator.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
 */
export function unless(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("unless");
	builder.left(left);
	builder.right(right);

	return builder;
}

/**
 * Arc tangent binary operator. Works in radians.
 * Trigonometric operators allow trigonometric functions to be executed on two vectors using vector matching, which isn't available with normal functions.
 * They act in the same manner as arithmetic operators.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#trigonometric-binary-operators
 */
export function atan2(left: cog.Builder<promql.Expr>,right: cog.Builder<promql.Expr>): BinaryExprBuilder {
	const builder = new BinaryExprBuilder();
	builder.op("atan2");
	builder.left(left);
	builder.right(right);

	return builder;
}

