# promql

## Objects

 * <span class="badge object-type-struct"></span> [AggregationExpr](./object-AggregationExpr.md)
 * <span class="badge object-type-enum"></span> [AggregationOp](./object-AggregationOp.md)
 * <span class="badge object-type-struct"></span> [BinaryExpr](./object-BinaryExpr.md)
 * <span class="badge object-type-enum"></span> [BinaryExprGroupModifier](./object-BinaryExprGroupModifier.md)
 * <span class="badge object-type-enum"></span> [BinaryExprMatchType](./object-BinaryExprMatchType.md)
 * <span class="badge object-type-enum"></span> [BinaryOp](./object-BinaryOp.md)
 * <span class="badge object-type-ref"></span> [Expr](./object-Expr.md)
 * <span class="badge object-type-struct"></span> [FuncCallExpr](./object-FuncCallExpr.md)
 * <span class="badge object-type-enum"></span> [LabelMatchingOperator](./object-LabelMatchingOperator.md)
 * <span class="badge object-type-struct"></span> [LabelSelector](./object-LabelSelector.md)
 * <span class="badge object-type-struct"></span> [NumberLiteralExpr](./object-NumberLiteralExpr.md)
 * <span class="badge object-type-struct"></span> [NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr](./object-NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr.md)
 * <span class="badge object-type-struct"></span> [StringLiteralExpr](./object-StringLiteralExpr.md)
 * <span class="badge object-type-struct"></span> [SubqueryExpr](./object-SubqueryExpr.md)
 * <span class="badge object-type-struct"></span> [UnaryExpr](./object-UnaryExpr.md)
 * <span class="badge object-type-enum"></span> [UnaryOp](./object-UnaryOp.md)
 * <span class="badge object-type-struct"></span> [VectorExpr](./object-VectorExpr.md)
## Builders

 * <span class="badge builder"></span> [AggregationExprBuilder](./builder-AggregationExprBuilder.md)
 * <span class="badge builder"></span> [BinaryExprBuilder](./builder-BinaryExprBuilder.md)
 * <span class="badge builder"></span> [ExprBuilder](./builder-ExprBuilder.md)
 * <span class="badge builder"></span> [FuncCallExprBuilder](./builder-FuncCallExprBuilder.md)
 * <span class="badge builder"></span> [LabelSelectorBuilder](./builder-LabelSelectorBuilder.md)
 * <span class="badge builder"></span> [NumberLiteralBuilder](./builder-NumberLiteralBuilder.md)
 * <span class="badge builder"></span> [StringLiteralBuilder](./builder-StringLiteralBuilder.md)
 * <span class="badge builder"></span> [SubqueryExprBuilder](./builder-SubqueryExprBuilder.md)
 * <span class="badge builder"></span> [UnaryExprBuilder](./builder-UnaryExprBuilder.md)
 * <span class="badge builder"></span> [VectorExprBuilder](./builder-VectorExprBuilder.md)
## Functions

### <span class="badge function"></span> NewExpr

NewExpr creates a new Expr object.

```go
func NewExpr() *Expr
```

### <span class="badge function"></span> NewNumberLiteralExpr

NewNumberLiteralExpr creates a new NumberLiteralExpr object.

```go
func NewNumberLiteralExpr() *NumberLiteralExpr
```

### <span class="badge function"></span> NewStringLiteralExpr

NewStringLiteralExpr creates a new StringLiteralExpr object.

```go
func NewStringLiteralExpr() *StringLiteralExpr
```

### <span class="badge function"></span> NewSubqueryExpr

NewSubqueryExpr creates a new SubqueryExpr object.

```go
func NewSubqueryExpr() *SubqueryExpr
```

### <span class="badge function"></span> NewAggregationExpr

NewAggregationExpr creates a new AggregationExpr object.

```go
func NewAggregationExpr() *AggregationExpr
```

### <span class="badge function"></span> NewVectorExpr

NewVectorExpr creates a new VectorExpr object.

```go
func NewVectorExpr() *VectorExpr
```

### <span class="badge function"></span> NewLabelSelector

NewLabelSelector creates a new LabelSelector object.

```go
func NewLabelSelector() *LabelSelector
```

### <span class="badge function"></span> NewBinaryExpr

NewBinaryExpr creates a new BinaryExpr object.

```go
func NewBinaryExpr() *BinaryExpr
```

### <span class="badge function"></span> NewUnaryExpr

NewUnaryExpr creates a new UnaryExpr object.

```go
func NewUnaryExpr() *UnaryExpr
```

### <span class="badge function"></span> NewFuncCallExpr

NewFuncCallExpr creates a new FuncCallExpr object.

```go
func NewFuncCallExpr() *FuncCallExpr
```

### <span class="badge function"></span> NewNumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr

NewNumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr creates a new NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr object.

```go
func NewNumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr() *NumberLiteralExprOrStringLiteralExprOrSubqueryExprOrAggregationExprOrVectorExprOrBinaryExprOrUnaryExprOrFuncCallExpr
```

### <span class="badge function"></span> Sum

Calculate sum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Sum(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Min

Calculate minimum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Min(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Max

Calculate maximum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Max(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Avg

Calculate the average over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Avg(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Group

All values in the resulting vector are 1.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Group(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Stddev

Calculate population standard deviation over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Stddev(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Stdvar

Calculate population standard variance over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Stdvar(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Count

Count number of elements in the vector.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Count(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Quantile

Calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Quantile(vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> CountValues

Count number of elements with the same value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func CountValues(label string, vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Bottomk

Smallest k elements by sample value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Bottomk(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Topk

Largest k elements by sample value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Topk(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Limitk

Sample k elements.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func Limitk(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> LimitRatio

Sample elements with approximately r ratio if r > 0, and the complement of such samples if r = -(1.0 - r).

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```go
func LimitRatio(k float64, vector cog.Builder[Expr]) *AggregationExprBuilder
```

### <span class="badge function"></span> Add

Addition binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```go
func Add(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Sub

Subtraction binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```go
func Sub(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Mul

Multiplication binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```go
func Mul(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Div

Division binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```go
func Div(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Mod

Modulo binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```go
func Mod(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Pow

Power/exponentiation binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```go
func Pow(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Eq

"equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```go
func Eq(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Neq

"not-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```go
func Neq(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Gt

"greater-than" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```go
func Gt(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Lt

"less-than" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```go
func Lt(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Gte

"greater-or-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```go
func Gte(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Lte

"less-or-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```go
func Lte(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> And

"intersection" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```go
func And(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Or

"union" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```go
func Or(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Unless

"complement" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```go
func Unless(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Atan2

Arc tangent binary operator. Works in radians.

Trigonometric operators allow trigonometric functions to be executed on two vectors using vector matching, which isn't available with normal functions.

They act in the same manner as arithmetic operators.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#trigonometric-binary-operators

```go
func Atan2(left cog.Builder[Expr], right cog.Builder[Expr]) *BinaryExprBuilder
```

### <span class="badge function"></span> Abs

Returns the input vector with all sample values converted to their absolute value.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#abs

```go
func Abs(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Absent

Returns an empty vector if the vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the vector passed to it has no elements.

This is useful for alerting on when no time series exist for a given metric name and label combination.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent

```go
func Absent(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> AbsentOverTime

Returns an empty vector if the range vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the range vector passed to it has no elements.

This is useful for alerting on when no time series exist for a given metric name and label combination for a certain amount of time.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent_over_time

```go
func AbsentOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Ceil

Rounds the sample values of all elements in `v` up to the nearest integer value greater than or equal to v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#ceil

```go
func Ceil(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Changes

For each input time series, returns the number of times its value has changed within the provided time range as an instant vector.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#changes

```go
func Changes(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Clamp

Clamps the sample values of all elements in `v` to have a lower limit of min and an upper limit of max.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp

```go
func Clamp(v cog.Builder[Expr], min float64, max float64) *FuncCallExprBuilder
```

### <span class="badge function"></span> ClampMax

Clamps the sample values of all elements in `v` to have an upper limit of `max`.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_max

```go
func ClampMax(v cog.Builder[Expr], max float64) *FuncCallExprBuilder
```

### <span class="badge function"></span> ClampMin

Clamps the sample values of all elements in `v` to have an lower limit of `min`.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_min

```go
func ClampMin(v cog.Builder[Expr], min float64) *FuncCallExprBuilder
```

### <span class="badge function"></span> DayOfMonth

Returns the day of the month in UTC. Returned values are from 1 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month

```go
func DayOfMonth() *FuncCallExprBuilder
```

### <span class="badge function"></span> DayOfMonthFor

Returns the day of the month for each of the given times in UTC. Returned values are from 1 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month

```go
func DayOfMonthFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> DayOfWeek

Returns the day of the week in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week

```go
func DayOfWeek() *FuncCallExprBuilder
```

### <span class="badge function"></span> DayOfWeekFor

Returns the day of the week for each of the given times in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week

```go
func DayOfWeekFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> DayOfYear

Returns the day of the year in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year

```go
func DayOfYear() *FuncCallExprBuilder
```

### <span class="badge function"></span> DayOfYearFor

Returns the day of the year for each of the given times in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year

```go
func DayOfYearFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> DaysInMonth

Returns the number of days in the month. Returned values are from 28 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month

```go
func DaysInMonth() *FuncCallExprBuilder
```

### <span class="badge function"></span> DayInMonthFor

Returns the number of days in the month for each of the given times in UTC. Returned values are from 28 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month

```go
func DayInMonthFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Delta

Calculates the difference between the first and last value of each time series element in a range vector, returning an instant vector with the given deltas and equivalent labels.

The delta is extrapolated to cover the full time range as specified in the range vector selector, so that it is possible to get a non-integer result even if the sample values are all integers.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#delta

```go
func Delta(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Deriv

Calculates the per-second derivative of the time series in a range vector using simple linear regression.

The range vector must have at least two samples in order to perform the calculation. When +Inf or -Inf are found in the range vector, the slope and offset value calculated will be NaN.

deriv should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv

```go
func Deriv(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Exp

Calculates the exponential function for all elements in vector

See https://prometheus.io/docs/prometheus/latest/querying/functions/#exp

```go
func Exp(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Floor

Rounds the sample values of all elements in v down to the nearest integer value smaller than or equal to v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#floor

```go
func Floor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramAvg

Returns the arithmetic average of observed values stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_avg

```go
func HistogramAvg(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramCount

Returns the count of observations stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum

```go
func HistogramCount(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramSum

Returns the sum of observations stored in a native histogram.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum

```go
func HistogramSum(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramFraction

Returns the estimated fraction of observations between the provided lower and upper values. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_fraction

```go
func HistogramFraction(lower float64, upper float64, v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramQuantile

Calculates the φ-quantile (0 ≤ φ ≤ 1) from a classic histogram or from a native histogram.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile

```go
func HistogramQuantile(phi float64, v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramStddev

Returns the estimated standard deviation of observations in a native histogram, based on the geometric mean of the buckets where the observations lie.

Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stddev

```go
func HistogramStddev(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> HistogramStdvar

Returns the estimated standard variance of observations in a native histogram.

Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stdvar

```go
func HistogramStdvar(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Hour

Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour

```go
func Hour() *FuncCallExprBuilder
```

### <span class="badge function"></span> HourFor

Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour

```go
func HourFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Idelta

Calculates the difference between the last two samples in the range vector v, returning an instant vector with the given deltas and equivalent labels.

idelta should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta

```go
func Idelta(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Increase

Calculates the increase in the time series in the range vector

See https://prometheus.io/docs/prometheus/latest/querying/functions/#increase

```go
func Increase(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Irate

Calculates the per-second instant rate of increase of the time series in the range vector. This is based on the last two data points.

irate should only be used when graphing volatile, fast-moving counters. Use rate for alerts and slow-moving counters, as brief changes in the rate can reset the FOR clause and graphs consisting entirely of rare spikes are hard to read.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#irate

```go
func Irate(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> LabelReplace

matches the regular expression regex against the value of the label src_label. If it matches, the value of the label dst_label in the returned timeseries will be the expansion of replacement, together with the original labels in the input. Capturing groups in the regular expression can be referenced with $1, $2, etc. Named capturing groups in the regular expression can be referenced with $name (where name is the capturing group name). If the regular expression doesn't match then the timeseries is returned unchanged.

label_replace acts on float and histogram samples in the same way.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace

```go
func LabelReplace(v cog.Builder[Expr], dst_label string, replacement string, src_label string, regex string) *FuncCallExprBuilder
```

### <span class="badge function"></span> Ln

Calculates the natural logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#ln

```go
func Ln(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Log2

Calculates the binary logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#log2

```go
func Log2(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Log10

Calculates the decimal logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#log10

```go
func Log10(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Minute

Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute

```go
func Minute() *FuncCallExprBuilder
```

### <span class="badge function"></span> MinuteFor

Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute

```go
func MinuteFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Month

Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#month

```go
func Month() *FuncCallExprBuilder
```

### <span class="badge function"></span> MonthFor

Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#month

```go
func MonthFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> PredictLinear

Predicts the value of time series t seconds from now, based on the range vector v, using simple linear regression. The range vector must have at least two samples in order to perform the calculation.

predict_linear should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear

```go
func PredictLinear(v cog.Builder[Expr], t float64) *FuncCallExprBuilder
```

### <span class="badge function"></span> Rate

Calculates the per-second average rate of increase of the time series in the range vector.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#rate

```go
func Rate(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Resets

For each input time series, resets(v range-vector) returns the number of counter resets within the provided time range as an instant vector. Any decrease in the value between two consecutive float samples is interpreted as a counter reset. A reset in a native histogram is detected in a more complex way: Any decrease in any bucket, including the zero bucket, or in the count of observation constitutes a counter reset, but also the disappearance of any previously populated bucket, an increase in bucket resolution, or a decrease of the zero-bucket width.

`resets` should only be used with counters and counter-like native histograms.

If the range vector contains a mix of float and histogram samples for the same series, counter resets are detected separately and their numbers added up. The change from a float to a histogram sample is not considered a counter reset. Each float sample is compared to the next float sample, and each histogram is compared to the next histogram.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#resets

```go
func Resets(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Round

Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#round

```go
func Round(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> RoundTo

Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.

The to_nearest argument allows specifying the nearest multiple to which the sample values should be rounded. This multiple may also be a fraction.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#round

```go
func RoundTo(v cog.Builder[Expr], to_nearest float64) *FuncCallExprBuilder
```

### <span class="badge function"></span> Scalar

Given a single-element input vector, scalar() returns the sample value of that single element as a scalar.

If the input vector does not have exactly one element, scalar will return NaN.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#scalar

```go
func Scalar(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Sgn

Returns a vector with all sample values converted to their sign, defined as this: 1 if v is positive, -1 if v is negative and 0 if v is equal to zero.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sgn

```go
func Sgn(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Sort

Returns vector elements sorted by their sample values, in ascending order. Native histograms are sorted by their sum of observations.

Note that sort only affects the results of instant queries, as range query results always have a fixed output ordering.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort

```go
func Sort(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> SortDesc

Same as `sort()`, but sorts in descending order.

Like sort, sort_desc only affects the results of instant queries, as range query results always have a fixed output ordering.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort_desc

```go
func SortDesc(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Sqrt

Calculates the square root of all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sqrt

```go
func Sqrt(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Time

Returns the number of seconds since January 1, 1970 UTC. Note that this does not actually return the current time, but the time at which the expression is to be evaluated.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#time

```go
func Time() *FuncCallExprBuilder
```

### <span class="badge function"></span> Timestamp

Returns the timestamp of each of the samples of the given vector as the number of seconds since January 1, 1970 UTC. It also works with histogram samples.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp

```go
func Timestamp(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Vect

Returns the scalar s as a vector with no labels.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector

```go
func Vect(s float64) *FuncCallExprBuilder
```

### <span class="badge function"></span> Year

Returns the year for each of the given times in UTC.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#year

```go
func Year() *FuncCallExprBuilder
```

### <span class="badge function"></span> YearFor

Returns the year for each of the given times in UTC.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#year

```go
func YearFor(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> AvgOverTime

Calculates average value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func AvgOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> MinOverTime

Calculates the minimum value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func MinOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> MaxOverTime

Calculates the maximum value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func MaxOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> SumOverTime

Calculates the sum of all values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func SumOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> CountOverTime

Calculates the count of all values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func CountOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> QuantileOverTime

Calculates the φ-quantile (0 ≤ φ ≤ 1) of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func QuantileOverTime(phi float64, v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> StddevOverTime

Calculates the population standard deviation of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func StddevOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> StdvarOverTime

Calculates the population standard variance of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func StdvarOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> LastOverTime

Returns the most recent point value in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func LastOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> PresentOverTime

Returns the value 1 for any series in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```go
func PresentOverTime(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Acos

Calculates the arccosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Acos(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Acosh

Calculates the inverse hyperbolic cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Acosh(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Asin

Calculates the arcsine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Asin(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Asinh

Calculates the inverse hyperbolic sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Asinh(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Atan

Calculates the arctangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Atan(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Atanh

Calculates the inverse hyperbolic tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Atanh(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Cos

Calculates the cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Cos(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Cosh

Calculates the hyperbolic cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Cosh(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Sin

Calculates the sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Sin(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Sinh

Calculates the hyperbolic sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Sinh(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Tan

Calculates the tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Tan(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Tanh

Calculates the hyperbolic tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Tanh(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Deg

Converts radians to degrees for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Deg(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> Pi

Returns pi.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Pi() *FuncCallExprBuilder
```

### <span class="badge function"></span> Rad

Converts degrees to radians for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```go
func Rad(v cog.Builder[Expr]) *FuncCallExprBuilder
```

### <span class="badge function"></span> N

Shortcut to turn a number into a NumberLiteral expression.

```go
func N(value float64) *NumberLiteralBuilder
```

### <span class="badge function"></span> S

Shortcut to turn a string into a StringLiteral expression.

```go
func S(value string) *StringLiteralBuilder
```

### <span class="badge function"></span> Subquery

Creates a subquery.

Subquery allows you to run an instant query for a given range and resolution. The result of a subquery is a range vector.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery

```go
func Subquery(expression cog.Builder[Expr]) *SubqueryExprBuilder
```

### <span class="badge function"></span> Neg

Negation unary operator.

```go
func Neg(expr cog.Builder[Expr]) *UnaryExprBuilder
```

### <span class="badge function"></span> Id

Identity unary operator.

```go
func Id(expr cog.Builder[Expr]) *UnaryExprBuilder
```

### <span class="badge function"></span> Vector

Returns the scalar s as a vector with no labels.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector

```go
func Vector(s string) *VectorExprBuilder
```

