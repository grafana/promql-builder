// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents a string literal expression.
// See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
export class StringLiteralExprBuilder implements cog.Builder<promql.StringLiteralExpr> {
    protected readonly internal: promql.StringLiteralExpr;

    constructor() {
        this.internal = promql.defaultStringLiteralExpr();
        this.internal.type = "stringLiteralExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.StringLiteralExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

/**
 * Shortcut to turn a string into a StringLiteralExpr expression.
 */
export function s(value: string): StringLiteralExprBuilder {
	const builder = new StringLiteralExprBuilder();
	builder.value(value);

	return builder;
}

