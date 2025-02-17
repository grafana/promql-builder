// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package promql

import (
	cog "github.com/grafana/promql-builder/go/cog"
)

var _ cog.Builder[Expr] = (*FuncCallExprBuilder)(nil)

// Represents a PromQL expression.
type FuncCallExprBuilder struct {
	internal *Expr
	errors   map[string]cog.BuildErrors
}

func NewFuncCallExprBuilder() *FuncCallExprBuilder {
	resource := NewExpr()
	builder := &FuncCallExprBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	if builder.internal.FuncCallExpr == nil {
		builder.internal.FuncCallExpr = NewFuncCallExpr()
	}
	builder.internal.FuncCallExpr.Type = "funcCallExpr"

	return builder
}

// Returns the input vector with all sample values converted to their absolute value.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#abs
func Abs(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("abs")
	builder.Arg(v)

	return builder
}

// Returns an empty vector if the vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the vector passed to it has no elements.
// This is useful for alerting on when no time series exist for a given metric name and label combination.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent
func Absent(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("absent")
	builder.Arg(v)

	return builder
}

// Returns an empty vector if the range vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the range vector passed to it has no elements.
// This is useful for alerting on when no time series exist for a given metric name and label combination for a certain amount of time.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent_over_time
func AbsentOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("absent_over_time")
	builder.Arg(v)

	return builder
}

// Rounds the sample values of all elements in `v` up to the nearest integer value greater than or equal to v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#ceil
func Ceil(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("ceil")
	builder.Arg(v)

	return builder
}

// For each input time series, returns the number of times its value has changed within the provided time range as an instant vector.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#changes
func Changes(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("changes")
	builder.Arg(v)

	return builder
}

// Clamps the sample values of all elements in `v` to have a lower limit of min and an upper limit of max.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp
func Clamp(v cog.Builder[Expr], min float64, max float64) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("clamp")
	builder.Arg(v)
	builder.Arg(N(min))
	builder.Arg(N(max))

	return builder
}

// Clamps the sample values of all elements in `v` to have an upper limit of `max`.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_max
func ClampMax(v cog.Builder[Expr], max float64) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("clamp_max")
	builder.Arg(v)
	builder.Arg(N(max))

	return builder
}

// Clamps the sample values of all elements in `v` to have an lower limit of `min`.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_min
func ClampMin(v cog.Builder[Expr], min float64) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("clamp_min")
	builder.Arg(v)
	builder.Arg(N(min))

	return builder
}

// Returns the day of the month in UTC. Returned values are from 1 to 31.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month
func DayOfMonth() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("day_of_month")

	return builder
}

// Returns the day of the month for each of the given times in UTC. Returned values are from 1 to 31.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month
func DayOfMonthFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("day_of_month")
	builder.Arg(v)

	return builder
}

// Returns the day of the week in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week
func DayOfWeek() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("day_of_week")

	return builder
}

// Returns the day of the week for each of the given times in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week
func DayOfWeekFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("day_of_week")
	builder.Arg(v)

	return builder
}

// Returns the day of the year in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year
func DayOfYear() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("day_of_year")

	return builder
}

// Returns the day of the year for each of the given times in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year
func DayOfYearFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("day_of_year")
	builder.Arg(v)

	return builder
}

// Returns the number of days in the month. Returned values are from 28 to 31.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month
func DaysInMonth() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("days_in_month")

	return builder
}

// Returns the number of days in the month for each of the given times in UTC. Returned values are from 28 to 31.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month
func DayInMonthFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("days_in_month")
	builder.Arg(v)

	return builder
}

// Calculates the difference between the first and last value of each time series element in a range vector, returning an instant vector with the given deltas and equivalent labels.
// The delta is extrapolated to cover the full time range as specified in the range vector selector, so that it is possible to get a non-integer result even if the sample values are all integers.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#delta
func Delta(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("delta")
	builder.Arg(v)

	return builder
}

// Calculates the per-second derivative of the time series in a range vector using simple linear regression.
// The range vector must have at least two samples in order to perform the calculation. When +Inf or -Inf are found in the range vector, the slope and offset value calculated will be NaN.
// deriv should only be used with gauges.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv
func Deriv(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("deriv")
	builder.Arg(v)

	return builder
}

// Calculates the exponential function for all elements in vector
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#exp
func Exp(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("exp")
	builder.Arg(v)

	return builder
}

// Rounds the sample values of all elements in v down to the nearest integer value smaller than or equal to v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#floor
func Floor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("floor")
	builder.Arg(v)

	return builder
}

// Returns the arithmetic average of observed values stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.
// Note: This function only acts on native histograms.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_avg
func HistogramAvg(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_avg")
	builder.Arg(v)

	return builder
}

// Returns the count of observations stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.
// Note: This function only acts on native histograms.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum
func HistogramCount(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_count")
	builder.Arg(v)

	return builder
}

// Returns the sum of observations stored in a native histogram.
// Note: This function only acts on native histograms.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum
func HistogramSum(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_sum")
	builder.Arg(v)

	return builder
}

// Returns the estimated fraction of observations between the provided lower and upper values. Samples that are not native histograms are ignored and do not show up in the returned vector.
// Note: This function only acts on native histograms.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_fraction
func HistogramFraction(lower float64, upper float64, v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_fraction")
	builder.Arg(N(lower))
	builder.Arg(N(upper))
	builder.Arg(v)

	return builder
}

// Calculates the φ-quantile (0 ≤ φ ≤ 1) from a classic histogram or from a native histogram.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile
func HistogramQuantile(phi float64, v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_quantile")
	builder.Arg(N(phi))
	builder.Arg(v)

	return builder
}

// Returns the estimated standard deviation of observations in a native histogram, based on the geometric mean of the buckets where the observations lie.
// Samples that are not native histograms are ignored and do not show up in the returned vector.
// Note: This function only acts on native histograms.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stddev
func HistogramStddev(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_stddev")
	builder.Arg(v)

	return builder
}

// Returns the estimated standard variance of observations in a native histogram.
// Samples that are not native histograms are ignored and do not show up in the returned vector.
// Note: This function only acts on native histograms.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stdvar
func HistogramStdvar(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("histogram_stdvar")
	builder.Arg(v)

	return builder
}

// Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour
func Hour() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("hour")

	return builder
}

// Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour
func HourFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("hour")
	builder.Arg(v)

	return builder
}

// Calculates the difference between the last two samples in the range vector v, returning an instant vector with the given deltas and equivalent labels.
// idelta should only be used with gauges.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta
func Idelta(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("idelta")
	builder.Arg(v)

	return builder
}

// Calculates the increase in the time series in the range vector
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#increase
func Increase(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("increase")
	builder.Arg(v)

	return builder
}

// Calculates the per-second instant rate of increase of the time series in the range vector. This is based on the last two data points.
// irate should only be used when graphing volatile, fast-moving counters. Use rate for alerts and slow-moving counters, as brief changes in the rate can reset the FOR clause and graphs consisting entirely of rare spikes are hard to read.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#irate
func Irate(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("irate")
	builder.Arg(v)

	return builder
}

// matches the regular expression regex against the value of the label src_label. If it matches, the value of the label dst_label in the returned timeseries will be the expansion of replacement, together with the original labels in the input. Capturing groups in the regular expression can be referenced with $1, $2, etc. Named capturing groups in the regular expression can be referenced with $name (where name is the capturing group name). If the regular expression doesn't match then the timeseries is returned unchanged.
// label_replace acts on float and histogram samples in the same way.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace
func LabelReplace(v cog.Builder[Expr], dstLabel string, replacement string, srcLabel string, regex string) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("label_replace")
	builder.Arg(v)
	builder.Arg(S(dstLabel))
	builder.Arg(S(replacement))
	builder.Arg(S(srcLabel))
	builder.Arg(S(regex))

	return builder
}

// Calculates the natural logarithm for all elements in v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#ln
func Ln(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("ln")
	builder.Arg(v)

	return builder
}

// Calculates the binary logarithm for all elements in v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#log2
func Log2(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("log2")
	builder.Arg(v)

	return builder
}

// Calculates the decimal logarithm for all elements in v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#log10
func Log10(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("log10")
	builder.Arg(v)

	return builder
}

// Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute
func Minute() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("minute")

	return builder
}

// Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute
func MinuteFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("minute")
	builder.Arg(v)

	return builder
}

// Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#month
func Month() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("month")

	return builder
}

// Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#month
func MonthFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("month")
	builder.Arg(v)

	return builder
}

// Predicts the value of time series t seconds from now, based on the range vector v, using simple linear regression. The range vector must have at least two samples in order to perform the calculation.
// predict_linear should only be used with gauges.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear
func PredictLinear(v cog.Builder[Expr], t float64) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("predict_linear")
	builder.Arg(v)
	builder.Arg(N(t))

	return builder
}

// Calculates the per-second average rate of increase of the time series in the range vector.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#rate
func Rate(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("rate")
	builder.Arg(v)

	return builder
}

// For each input time series, resets(v range-vector) returns the number of counter resets within the provided time range as an instant vector. Any decrease in the value between two consecutive float samples is interpreted as a counter reset. A reset in a native histogram is detected in a more complex way: Any decrease in any bucket, including the zero bucket, or in the count of observation constitutes a counter reset, but also the disappearance of any previously populated bucket, an increase in bucket resolution, or a decrease of the zero-bucket width.
// `resets` should only be used with counters and counter-like native histograms.
// If the range vector contains a mix of float and histogram samples for the same series, counter resets are detected separately and their numbers added up. The change from a float to a histogram sample is not considered a counter reset. Each float sample is compared to the next float sample, and each histogram is compared to the next histogram.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#resets
func Resets(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("resets")
	builder.Arg(v)

	return builder
}

// Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#round
func Round(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("round")
	builder.Arg(v)

	return builder
}

// Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.
// The to_nearest argument allows specifying the nearest multiple to which the sample values should be rounded. This multiple may also be a fraction.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#round
func RoundTo(v cog.Builder[Expr], toNearest float64) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("round")
	builder.Arg(v)
	builder.Arg(N(toNearest))

	return builder
}

// Given a single-element input vector, scalar() returns the sample value of that single element as a scalar.
// If the input vector does not have exactly one element, scalar will return NaN.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#scalar
func Scalar(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("scalar")
	builder.Arg(v)

	return builder
}

// Returns a vector with all sample values converted to their sign, defined as this: 1 if v is positive, -1 if v is negative and 0 if v is equal to zero.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#sgn
func Sgn(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sgn")
	builder.Arg(v)

	return builder
}

// Returns vector elements sorted by their sample values, in ascending order. Native histograms are sorted by their sum of observations.
// Note that sort only affects the results of instant queries, as range query results always have a fixed output ordering.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort
func Sort(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sort")
	builder.Arg(v)

	return builder
}

// Same as `sort()`, but sorts in descending order.
// Like sort, sort_desc only affects the results of instant queries, as range query results always have a fixed output ordering.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort_desc
func SortDesc(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sort_desc")
	builder.Arg(v)

	return builder
}

// Calculates the square root of all elements in v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#sqrt
func Sqrt(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sqrt")
	builder.Arg(v)

	return builder
}

// Returns the number of seconds since January 1, 1970 UTC. Note that this does not actually return the current time, but the time at which the expression is to be evaluated.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#time
func Time() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("time")

	return builder
}

// Returns the timestamp of each of the samples of the given vector as the number of seconds since January 1, 1970 UTC. It also works with histogram samples.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp
func Timestamp(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("timestamp")
	builder.Arg(v)

	return builder
}

// Returns the scalar s as a vector with no labels.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector
func Vect(s float64) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("vector")
	builder.Arg(N(s))

	return builder
}

// Returns the year for each of the given times in UTC.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#year
func Year() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("year")

	return builder
}

// Returns the year for each of the given times in UTC.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#year
func YearFor(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("year")
	builder.Arg(v)

	return builder
}

// Calculates average value of all points in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func AvgOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("avg_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the minimum value of all points in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func MinOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("min_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the maximum value of all points in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func MaxOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("max_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the sum of all values in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func SumOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sum_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the count of all values in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func CountOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("count_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the φ-quantile (0 ≤ φ ≤ 1) of the values in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func QuantileOverTime(phi float64, v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("quantile_over_time")
	builder.Arg(N(phi))
	builder.Arg(v)

	return builder
}

// Calculates the population standard deviation of the values in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func StddevOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("stddev_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the population standard variance of the values in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func StdvarOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("stdvar_over_time")
	builder.Arg(v)

	return builder
}

// Returns the most recent point value in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func LastOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("last_over_time")
	builder.Arg(v)

	return builder
}

// Returns the value 1 for any series in the specified interval.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
func PresentOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("present_over_time")
	builder.Arg(v)

	return builder
}

// Calculates the arccosine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Acos(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("acos")
	builder.Arg(v)

	return builder
}

// Calculates the inverse hyperbolic cosine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Acosh(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("acosh")
	builder.Arg(v)

	return builder
}

// Calculates the arcsine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Asin(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("asin")
	builder.Arg(v)

	return builder
}

// Calculates the inverse hyperbolic sine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Asinh(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("asinh")
	builder.Arg(v)

	return builder
}

// Calculates the arctangent of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Atan(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("atan")
	builder.Arg(v)

	return builder
}

// Calculates the inverse hyperbolic tangent of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Atanh(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("atanh")
	builder.Arg(v)

	return builder
}

// Calculates the cosine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Cos(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("cos")
	builder.Arg(v)

	return builder
}

// Calculates the hyperbolic cosine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Cosh(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("cosh")
	builder.Arg(v)

	return builder
}

// Calculates the sine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Sin(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sin")
	builder.Arg(v)

	return builder
}

// Calculates the hyperbolic sine of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Sinh(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("sinh")
	builder.Arg(v)

	return builder
}

// Calculates the tangent of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Tan(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("tan")
	builder.Arg(v)

	return builder
}

// Calculates the hyperbolic tangent of all elements in v
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Tanh(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("tanh")
	builder.Arg(v)

	return builder
}

// Converts radians to degrees for all elements in v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Deg(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("deg")
	builder.Arg(v)

	return builder
}

// Returns pi.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Pi() *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("pi")

	return builder
}

// Converts degrees to radians for all elements in v.
// See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
func Rad(v cog.Builder[Expr]) *FuncCallExprBuilder {
	builder := NewFuncCallExprBuilder()
	builder.Function("rad")
	builder.Arg(v)

	return builder
}

func (builder *FuncCallExprBuilder) Build() (Expr, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	return *builder.internal, nil
}

func (builder FuncCallExprBuilder) String() string {
	return builder.internal.String()
}

// Name of the function.
func (builder *FuncCallExprBuilder) Function(function string) *FuncCallExprBuilder {
	if builder.internal.FuncCallExpr == nil {
		builder.internal.FuncCallExpr = NewFuncCallExpr()
	}
	builder.internal.FuncCallExpr.Function = function

	return builder
}

// Arguments.
func (builder *FuncCallExprBuilder) Args(args []cog.Builder[Expr]) *FuncCallExprBuilder {
	if builder.internal.FuncCallExpr == nil {
		builder.internal.FuncCallExpr = NewFuncCallExpr()
	}
	argsResources := make([]Expr, 0, len(args))
	for _, r1 := range args {
		argsDepth1, err := r1.Build()
		if err != nil {
			builder.errors["FuncCallExpr.args"] = err.(cog.BuildErrors)
			return builder
		}
		argsResources = append(argsResources, argsDepth1)
	}
	builder.internal.FuncCallExpr.Args = argsResources

	return builder
}

// Arguments.
func (builder *FuncCallExprBuilder) Arg(arg cog.Builder[Expr]) *FuncCallExprBuilder {
	if builder.internal.FuncCallExpr == nil {
		builder.internal.FuncCallExpr = NewFuncCallExpr()
	}
	argResource, err := arg.Build()
	if err != nil {
		builder.errors["FuncCallExpr.args"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.FuncCallExpr.Args = append(builder.internal.FuncCallExpr.Args, argResource)

	return builder
}
