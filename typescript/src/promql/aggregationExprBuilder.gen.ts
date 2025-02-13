// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents an aggregation.
// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
export class AggregationExprBuilder implements cog.Builder<promql.AggregationExpr> {
    protected readonly internal: promql.AggregationExpr;

    constructor() {
        this.internal = promql.defaultAggregationExpr();
        this.internal.type = "aggregationExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.AggregationExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    op(op: promql.AggregationOp): this {
        this.internal.op = op;
        return this;
    }

    expr(expr: cog.Builder<promql.Expr>): this {
        const exprResource = expr.build();
        this.internal.expr = exprResource;
        return this;
    }

    param(param: cog.Builder<promql.Expr>): this {
        const paramResource = param.build();
        this.internal.param = paramResource;
        return this;
    }

    // By drops labels that are not listed in the by clause.
    by(by: string[]): this {
        this.internal.by = by;
        return this;
    }

    // List of labels to remove from the result vector, while all other labels are preserved in the output.
    without(without: string[]): this {
        this.internal.without = without;
        return this;
    }
}

/**
 * Calculate sum over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function sum(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("sum");
	builder.expr(vector);

	return builder;
}

/**
 * Calculate minimum over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function min(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("min");
	builder.expr(vector);

	return builder;
}

/**
 * Calculate maximum over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function max(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("max");
	builder.expr(vector);

	return builder;
}

/**
 * Calculate the average over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function avg(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("avg");
	builder.expr(vector);

	return builder;
}

/**
 * All values in the resulting vector are 1.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function group(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("group");
	builder.expr(vector);

	return builder;
}

/**
 * Calculate population standard deviation over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function stddev(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("stddev");
	builder.expr(vector);

	return builder;
}

/**
 * Calculate population standard variance over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function stdvar(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("stdvar");
	builder.expr(vector);

	return builder;
}

/**
 * Count number of elements in the vector.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function count(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("count");
	builder.expr(vector);

	return builder;
}

/**
 * Calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function quantile(vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("quantile");
	builder.expr(vector);

	return builder;
}

/**
 * Count number of elements with the same value.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function countValues(label: string,vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("count_values");
	builder.expr(vector);
	builder.param(promql.s(label));

	return builder;
}

/**
 * Smallest k elements by sample value.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function bottomk(k: number,vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("bottomk");
	builder.expr(vector);
	builder.param(promql.n(k));

	return builder;
}

/**
 * Largest k elements by sample value.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function topk(k: number,vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("topk");
	builder.expr(vector);
	builder.param(promql.n(k));

	return builder;
}

/**
 * Sample k elements.
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function limitk(k: number,vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("limitk");
	builder.expr(vector);
	builder.param(promql.n(k));

	return builder;
}

/**
 * Sample elements with approximately r ratio if r > 0, and the complement of such samples if r = -(1.0 - r).
 * See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
 */
export function limitRatio(k: number,vector: cog.Builder<promql.Expr>): AggregationExprBuilder {
	const builder = new AggregationExprBuilder();
	builder.op("limit_ratio");
	builder.expr(vector);
	builder.param(promql.n(k));

	return builder;
}

