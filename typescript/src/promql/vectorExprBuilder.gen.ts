// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents both instant and range vectors
export class VectorExprBuilder implements cog.Builder<promql.VectorExpr> {
    protected readonly internal: promql.VectorExpr;

    constructor() {
        this.internal = promql.defaultVectorExpr();
        this.internal.type = "vectorExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.VectorExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    // Metric name.
    metric(metric: string): this {
        this.internal.metric = metric;
        return this;
    }

    // Label selectors used to filter the timeseries.
    labels(labels: cog.Builder<promql.LabelSelector>[]): this {
        const labelsResources = labels.map(builder1 => builder1.build());
        this.internal.labels = labelsResources;
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

    label(name: string,value: string): this {
        if (!this.internal.labels) {
            this.internal.labels = [];
        }
        this.internal.labels.push({
        name: name,
        operator: "=",
        value: value,
    });
        return this;
    }

    labelNeq(name: string,value: string): this {
        if (!this.internal.labels) {
            this.internal.labels = [];
        }
        this.internal.labels.push({
        name: name,
        operator: "!=",
        value: value,
    });
        return this;
    }

    labelMatchRegexp(name: string,value: string): this {
        if (!this.internal.labels) {
            this.internal.labels = [];
        }
        this.internal.labels.push({
        name: name,
        operator: "=~",
        value: value,
    });
        return this;
    }

    labelNotMatchRegexp(name: string,value: string): this {
        if (!this.internal.labels) {
            this.internal.labels = [];
        }
        this.internal.labels.push({
        name: name,
        operator: "!~",
        value: value,
    });
        return this;
    }
}

/**
 * Returns the scalar s as a vector with no labels.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector
 */
export function vector(s: string): VectorExprBuilder {
	const builder = new VectorExprBuilder();
	builder.metric(s);

	return builder;
}

