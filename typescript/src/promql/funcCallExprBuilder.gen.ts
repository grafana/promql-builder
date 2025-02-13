// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as promql from '../promql';

// Represents a function call expression.
export class FuncCallExprBuilder implements cog.Builder<promql.FuncCallExpr> {
    protected readonly internal: promql.FuncCallExpr;

    constructor() {
        this.internal = promql.defaultFuncCallExpr();
        this.internal.type = "funcCallExpr";
    }

    /**
     * Builds the object.
     */
    build(): promql.FuncCallExpr {
        return this.internal;
    }

    toString(): string {
        return promql.toString(this.internal);
    }

    // Name of the function.
    functionVal(functionVal: string): this {
        if (!(functionVal.length >= 1)) {
            throw new Error("functionVal.length must be >= 1");
        }
        this.internal.function = functionVal;
        return this;
    }

    // Arguments.
    args(args: cog.Builder<promql.Expr>[]): this {
        const argsResources = args.map(builder1 => builder1.build());
        this.internal.args = argsResources;
        return this;
    }

    // Arguments.
    arg(arg: cog.Builder<promql.Expr>): this {
        if (!this.internal.args) {
            this.internal.args = [];
        }
        const argResource = arg.build();
        this.internal.args.push(argResource);
        return this;
    }
}

/**
 * Returns the input vector with all sample values converted to their absolute value.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#abs
 */
export function abs(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("abs");
	builder.arg(v);

	return builder;
}

/**
 * Returns an empty vector if the vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the vector passed to it has no elements.
 * This is useful for alerting on when no time series exist for a given metric name and label combination.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent
 */
export function absent(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("absent");
	builder.arg(v);

	return builder;
}

/**
 * Returns an empty vector if the range vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the range vector passed to it has no elements.
 * This is useful for alerting on when no time series exist for a given metric name and label combination for a certain amount of time.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent_over_time
 */
export function absentOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("absent_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Rounds the sample values of all elements in `v` up to the nearest integer value greater than or equal to v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#ceil
 */
export function ceil(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("ceil");
	builder.arg(v);

	return builder;
}

/**
 * For each input time series, returns the number of times its value has changed within the provided time range as an instant vector.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#changes
 */
export function changes(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("changes");
	builder.arg(v);

	return builder;
}

/**
 * Clamps the sample values of all elements in `v` to have a lower limit of min and an upper limit of max.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp
 */
export function clamp(v: cog.Builder<promql.Expr>,min: number,max: number): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("clamp");
	builder.arg(v);
	builder.arg(promql.n(min));
	builder.arg(promql.n(max));

	return builder;
}

/**
 * Clamps the sample values of all elements in `v` to have an upper limit of `max`.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_max
 */
export function clampMax(v: cog.Builder<promql.Expr>,max: number): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("clamp_max");
	builder.arg(v);
	builder.arg(promql.n(max));

	return builder;
}

/**
 * Clamps the sample values of all elements in `v` to have an lower limit of `min`.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_min
 */
export function clampMin(v: cog.Builder<promql.Expr>,min: number): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("clamp_min");
	builder.arg(v);
	builder.arg(promql.n(min));

	return builder;
}

/**
 * Returns the day of the month in UTC. Returned values are from 1 to 31.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month
 */
export function dayOfMonth(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("day_of_month");

	return builder;
}

/**
 * Returns the day of the month for each of the given times in UTC. Returned values are from 1 to 31.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month
 */
export function dayOfMonthFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("day_of_month");
	builder.arg(v);

	return builder;
}

/**
 * Returns the day of the week in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week
 */
export function dayOfWeek(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("day_of_week");

	return builder;
}

/**
 * Returns the day of the week for each of the given times in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week
 */
export function dayOfWeekFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("day_of_week");
	builder.arg(v);

	return builder;
}

/**
 * Returns the day of the year in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year
 */
export function dayOfYear(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("day_of_year");

	return builder;
}

/**
 * Returns the day of the year for each of the given times in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year
 */
export function dayOfYearFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("day_of_year");
	builder.arg(v);

	return builder;
}

/**
 * Returns the number of days in the month. Returned values are from 28 to 31.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month
 */
export function daysInMonth(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("days_in_month");

	return builder;
}

/**
 * Returns the number of days in the month for each of the given times in UTC. Returned values are from 28 to 31.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month
 */
export function dayInMonthFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("days_in_month");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the difference between the first and last value of each time series element in a range vector, returning an instant vector with the given deltas and equivalent labels.
 * The delta is extrapolated to cover the full time range as specified in the range vector selector, so that it is possible to get a non-integer result even if the sample values are all integers.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#delta
 */
export function delta(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("delta");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the per-second derivative of the time series in a range vector using simple linear regression.
 * The range vector must have at least two samples in order to perform the calculation. When +Inf or -Inf are found in the range vector, the slope and offset value calculated will be NaN.
 * deriv should only be used with gauges.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv
 */
export function deriv(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("deriv");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the exponential function for all elements in vector
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#exp
 */
export function exp(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("exp");
	builder.arg(v);

	return builder;
}

/**
 * Rounds the sample values of all elements in v down to the nearest integer value smaller than or equal to v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#floor
 */
export function floor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("floor");
	builder.arg(v);

	return builder;
}

/**
 * Returns the arithmetic average of observed values stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.
 * Note: This function only acts on native histograms.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_avg
 */
export function histogramAvg(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_avg");
	builder.arg(v);

	return builder;
}

/**
 * Returns the count of observations stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.
 * Note: This function only acts on native histograms.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum
 */
export function histogramCount(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_count");
	builder.arg(v);

	return builder;
}

/**
 * Returns the sum of observations stored in a native histogram.
 * Note: This function only acts on native histograms.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum
 */
export function histogramSum(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_sum");
	builder.arg(v);

	return builder;
}

/**
 * Returns the estimated fraction of observations between the provided lower and upper values. Samples that are not native histograms are ignored and do not show up in the returned vector.
 * Note: This function only acts on native histograms.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_fraction
 */
export function histogramFraction(lower: number,upper: number,v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_fraction");
	builder.arg(promql.n(lower));
	builder.arg(promql.n(upper));
	builder.arg(v);

	return builder;
}

/**
 * Calculates the φ-quantile (0 ≤ φ ≤ 1) from a classic histogram or from a native histogram.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile
 */
export function histogramQuantile(phi: number,v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_quantile");
	builder.arg(promql.n(phi));
	builder.arg(v);

	return builder;
}

/**
 * Returns the estimated standard deviation of observations in a native histogram, based on the geometric mean of the buckets where the observations lie.
 * Samples that are not native histograms are ignored and do not show up in the returned vector.
 * Note: This function only acts on native histograms.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stddev
 */
export function histogramStddev(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_stddev");
	builder.arg(v);

	return builder;
}

/**
 * Returns the estimated standard variance of observations in a native histogram.
 * Samples that are not native histograms are ignored and do not show up in the returned vector.
 * Note: This function only acts on native histograms.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stdvar
 */
export function histogramStdvar(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("histogram_stdvar");
	builder.arg(v);

	return builder;
}

/**
 * Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour
 */
export function hour(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("hour");

	return builder;
}

/**
 * Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour
 */
export function hourFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("hour");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the difference between the last two samples in the range vector v, returning an instant vector with the given deltas and equivalent labels.
 * idelta should only be used with gauges.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta
 */
export function idelta(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("idelta");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the increase in the time series in the range vector
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#increase
 */
export function increase(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("increase");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the per-second instant rate of increase of the time series in the range vector. This is based on the last two data points.
 * irate should only be used when graphing volatile, fast-moving counters. Use rate for alerts and slow-moving counters, as brief changes in the rate can reset the FOR clause and graphs consisting entirely of rare spikes are hard to read.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#irate
 */
export function irate(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("irate");
	builder.arg(v);

	return builder;
}

/**
 * matches the regular expression regex against the value of the label src_label. If it matches, the value of the label dst_label in the returned timeseries will be the expansion of replacement, together with the original labels in the input. Capturing groups in the regular expression can be referenced with $1, $2, etc. Named capturing groups in the regular expression can be referenced with $name (where name is the capturing group name). If the regular expression doesn't match then the timeseries is returned unchanged.
 * label_replace acts on float and histogram samples in the same way.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace
 */
export function labelReplace(v: cog.Builder<promql.Expr>,dstLabel: string,replacement: string,srcLabel: string,regex: string): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("label_replace");
	builder.arg(v);
	builder.arg(promql.s(dstLabel));
	builder.arg(promql.s(replacement));
	builder.arg(promql.s(srcLabel));
	builder.arg(promql.s(regex));

	return builder;
}

/**
 * Calculates the natural logarithm for all elements in v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#ln
 */
export function ln(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("ln");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the binary logarithm for all elements in v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#log2
 */
export function log2(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("log2");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the decimal logarithm for all elements in v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#log10
 */
export function log10(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("log10");
	builder.arg(v);

	return builder;
}

/**
 * Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute
 */
export function minute(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("minute");

	return builder;
}

/**
 * Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute
 */
export function minuteFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("minute");
	builder.arg(v);

	return builder;
}

/**
 * Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#month
 */
export function month(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("month");

	return builder;
}

/**
 * Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#month
 */
export function monthFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("month");
	builder.arg(v);

	return builder;
}

/**
 * Predicts the value of time series t seconds from now, based on the range vector v, using simple linear regression. The range vector must have at least two samples in order to perform the calculation.
 * predict_linear should only be used with gauges.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear
 */
export function predictLinear(v: cog.Builder<promql.Expr>,t: number): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("predict_linear");
	builder.arg(v);
	builder.arg(promql.n(t));

	return builder;
}

/**
 * Calculates the per-second average rate of increase of the time series in the range vector.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#rate
 */
export function rate(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("rate");
	builder.arg(v);

	return builder;
}

/**
 * For each input time series, resets(v range-vector) returns the number of counter resets within the provided time range as an instant vector. Any decrease in the value between two consecutive float samples is interpreted as a counter reset. A reset in a native histogram is detected in a more complex way: Any decrease in any bucket, including the zero bucket, or in the count of observation constitutes a counter reset, but also the disappearance of any previously populated bucket, an increase in bucket resolution, or a decrease of the zero-bucket width.
 * `resets` should only be used with counters and counter-like native histograms.
 * If the range vector contains a mix of float and histogram samples for the same series, counter resets are detected separately and their numbers added up. The change from a float to a histogram sample is not considered a counter reset. Each float sample is compared to the next float sample, and each histogram is compared to the next histogram.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#resets
 */
export function resets(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("resets");
	builder.arg(v);

	return builder;
}

/**
 * Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#round
 */
export function round(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("round");
	builder.arg(v);

	return builder;
}

/**
 * Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.
 * The to_nearest argument allows specifying the nearest multiple to which the sample values should be rounded. This multiple may also be a fraction.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#round
 */
export function roundTo(v: cog.Builder<promql.Expr>,toNearest: number): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("round");
	builder.arg(v);
	builder.arg(promql.n(toNearest));

	return builder;
}

/**
 * Given a single-element input vector, scalar() returns the sample value of that single element as a scalar.
 * If the input vector does not have exactly one element, scalar will return NaN.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#scalar
 */
export function scalar(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("scalar");
	builder.arg(v);

	return builder;
}

/**
 * Returns a vector with all sample values converted to their sign, defined as this: 1 if v is positive, -1 if v is negative and 0 if v is equal to zero.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#sgn
 */
export function sgn(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sgn");
	builder.arg(v);

	return builder;
}

/**
 * Returns vector elements sorted by their sample values, in ascending order. Native histograms are sorted by their sum of observations.
 * Note that sort only affects the results of instant queries, as range query results always have a fixed output ordering.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort
 */
export function sort(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sort");
	builder.arg(v);

	return builder;
}

/**
 * Same as `sort()`, but sorts in descending order.
 * Like sort, sort_desc only affects the results of instant queries, as range query results always have a fixed output ordering.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort_desc
 */
export function sortDesc(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sort_desc");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the square root of all elements in v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#sqrt
 */
export function sqrt(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sqrt");
	builder.arg(v);

	return builder;
}

/**
 * Returns the number of seconds since January 1, 1970 UTC. Note that this does not actually return the current time, but the time at which the expression is to be evaluated.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#time
 */
export function time(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("time");

	return builder;
}

/**
 * Returns the timestamp of each of the samples of the given vector as the number of seconds since January 1, 1970 UTC. It also works with histogram samples.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp
 */
export function timestamp(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("timestamp");
	builder.arg(v);

	return builder;
}

/**
 * Returns the scalar s as a vector with no labels.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector
 */
export function vect(s: number): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("vector");
	builder.arg(promql.n(s));

	return builder;
}

/**
 * Returns the year for each of the given times in UTC.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#year
 */
export function year(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("year");

	return builder;
}

/**
 * Returns the year for each of the given times in UTC.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#year
 */
export function yearFor(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("year");
	builder.arg(v);

	return builder;
}

/**
 * Calculates average value of all points in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function avgOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("avg_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the minimum value of all points in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function minOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("min_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the maximum value of all points in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function maxOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("max_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the sum of all values in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function sumOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sum_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the count of all values in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function countOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("count_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the φ-quantile (0 ≤ φ ≤ 1) of the values in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function quantileOverTime(phi: number,v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("quantile_over_time");
	builder.arg(promql.n(phi));
	builder.arg(v);

	return builder;
}

/**
 * Calculates the population standard deviation of the values in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function stddevOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("stddev_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the population standard variance of the values in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function stdvarOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("stdvar_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Returns the most recent point value in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function lastOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("last_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Returns the value 1 for any series in the specified interval.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
 */
export function presentOverTime(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("present_over_time");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the arccosine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function acos(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("acos");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the inverse hyperbolic cosine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function acosh(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("acosh");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the arcsine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function asin(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("asin");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the inverse hyperbolic sine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function asinh(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("asinh");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the arctangent of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function atan(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("atan");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the inverse hyperbolic tangent of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function atanh(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("atanh");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the cosine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function cos(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("cos");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the hyperbolic cosine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function cosh(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("cosh");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the sine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function sin(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sin");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the hyperbolic sine of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function sinh(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("sinh");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the tangent of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function tan(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("tan");
	builder.arg(v);

	return builder;
}

/**
 * Calculates the hyperbolic tangent of all elements in v
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function tanh(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("tanh");
	builder.arg(v);

	return builder;
}

/**
 * Converts radians to degrees for all elements in v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function deg(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("deg");
	builder.arg(v);

	return builder;
}

/**
 * Returns pi.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function pi(): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("pi");

	return builder;
}

/**
 * Converts degrees to radians for all elements in v.
 * See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
 */
export function rad(v: cog.Builder<promql.Expr>): FuncCallExprBuilder {
	const builder = new FuncCallExprBuilder();
	builder.functionVal("rad");
	builder.arg(v);

	return builder;
}

