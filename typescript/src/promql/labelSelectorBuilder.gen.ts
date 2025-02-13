// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

export class LabelSelectorBuilder implements cog.Builder<promql.LabelSelector> {
    protected readonly internal: promql.LabelSelector;

    constructor() {
        this.internal = promql.defaultLabelSelector();
    }

    /**
     * Builds the object.
     */
    build(): promql.LabelSelector {
        return this.internal;
    }

    toString(): string {
        return promql.labelToString(this.internal);
    }

    // Name of the label to select.
    name(name: string): this {
        if (!(name.length >= 1)) {
            throw new Error("name.length must be >= 1");
        }
        this.internal.name = name;
        return this;
    }

    // Value to match against.
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Operator used to perform the selection.
    operator(operator: promql.LabelMatchingOperator): this {
        this.internal.operator = operator;
        return this;
    }
}

