// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents a number literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
export class NumberLiteralExprBuilder implements cog.Builder<promql.NumberLiteralExpr> {
    protected readonly internal: promql.NumberLiteralExpr;

    constructor() {
        this.internal = promql.defaultNumberLiteralExpr();
        this.internal.type = "numberLiteralExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.NumberLiteralExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    value(value: number): this {
        this.internal.value = value;
        return this;
    }
}

/**
 * Shortcut to turn a number into a NumberLiteralExpr expression.
 */
export function n(value: number): NumberLiteralExprBuilder {
	const builder = new NumberLiteralExprBuilder();
	builder.value(value);

	return builder;
}

