// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents a subquery.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
export class SubqueryExprBuilder implements cog.Builder<promql.SubqueryExpr> {
    protected readonly internal: promql.SubqueryExpr;

    constructor() {
        this.internal = promql.defaultSubqueryExpr();
        this.internal.type = "subqueryExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.SubqueryExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    expr(expr: cog.Builder<promql.Expr>): this {
        const exprResource = expr.build();
        this.internal.expr = exprResource;
        return this;
    }

    // The offset modifier allows changing the time offset for individual instant and range vectors in a query.
    // https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
    offset(offset: string): this {
        this.internal.offset = offset;
        return this;
    }

    // The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
    // The time supplied to the @ modifier is a unix timestamp.
    // https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
    at(at: string): this {
        this.internal.at = at;
        return this;
    }

    // Range of samples back from the current instant.
    // https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
    range(range: string): this {
        this.internal.range = range;
        return this;
    }

    // Empty string for default resolution.
    resolution(resolution: string): this {
        this.internal.resolution = resolution;
        return this;
    }
}

/**
 * Creates a subquery.
 * Subquery allows you to run an instant query for a given range and resolution. The result of a subquery is a range vector.
 * See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
 */
export function subquery(expression: cog.Builder<promql.Expr>): SubqueryExprBuilder {
	const builder = new SubqueryExprBuilder();
	builder.expr(expression);

	return builder;
}

