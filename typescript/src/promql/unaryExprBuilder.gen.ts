// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents an unary operation expression.
export class UnaryExprBuilder implements cog.Builder<promql.UnaryExpr> {
    protected readonly internal: promql.UnaryExpr;

    constructor() {
        this.internal = promql.defaultUnaryExpr();
        this.internal.type = "unaryExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.UnaryExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    op(op: promql.UnaryOp): this {
        this.internal.op = op;
        return this;
    }

    expr(expr: cog.Builder<promql.Expr>): this {
        const exprResource = expr.build();
        this.internal.expr = exprResource;
        return this;
    }
}

/**
 * Negation unary operator.
 */
export function neg(expr: cog.Builder<promql.Expr>): UnaryExprBuilder {
	const builder = new UnaryExprBuilder();
	builder.op("-");
	builder.expr(expr);

	return builder;
}

/**
 * Identity unary operator.
 */
export function id(expr: cog.Builder<promql.Expr>): UnaryExprBuilder {
	const builder = new UnaryExprBuilder();
	builder.op("+");
	builder.expr(expr);

	return builder;
}

