# promql

## Objects

 * <span class="badge object-type-interface"></span> [AggregationExpr](./object-AggregationExpr.md)
 * <span class="badge object-type-enum"></span> [AggregationOp](./object-AggregationOp.md)
 * <span class="badge object-type-interface"></span> [BinaryExpr](./object-BinaryExpr.md)
 * <span class="badge object-type-enum"></span> [BinaryOp](./object-BinaryOp.md)
 * <span class="badge object-type-disjunction"></span> [Expr](./object-Expr.md)
 * <span class="badge object-type-interface"></span> [FuncCallExpr](./object-FuncCallExpr.md)
 * <span class="badge object-type-enum"></span> [LabelMatchingOperator](./object-LabelMatchingOperator.md)
 * <span class="badge object-type-interface"></span> [LabelSelector](./object-LabelSelector.md)
 * <span class="badge object-type-interface"></span> [NumberLiteralExpr](./object-NumberLiteralExpr.md)
 * <span class="badge object-type-interface"></span> [StringLiteralExpr](./object-StringLiteralExpr.md)
 * <span class="badge object-type-interface"></span> [SubqueryExpr](./object-SubqueryExpr.md)
 * <span class="badge object-type-interface"></span> [UnaryExpr](./object-UnaryExpr.md)
 * <span class="badge object-type-enum"></span> [UnaryOp](./object-UnaryOp.md)
 * <span class="badge object-type-interface"></span> [VectorExpr](./object-VectorExpr.md)
## Builders

 * <span class="badge builder"></span> [AggregationExprBuilder](./builder-AggregationExprBuilder.md)
 * <span class="badge builder"></span> [BinaryExprBuilder](./builder-BinaryExprBuilder.md)
 * <span class="badge builder"></span> [FuncCallExprBuilder](./builder-FuncCallExprBuilder.md)
 * <span class="badge builder"></span> [LabelSelectorBuilder](./builder-LabelSelectorBuilder.md)
 * <span class="badge builder"></span> [NumberLiteralExprBuilder](./builder-NumberLiteralExprBuilder.md)
 * <span class="badge builder"></span> [StringLiteralExprBuilder](./builder-StringLiteralExprBuilder.md)
 * <span class="badge builder"></span> [SubqueryExprBuilder](./builder-SubqueryExprBuilder.md)
 * <span class="badge builder"></span> [UnaryExprBuilder](./builder-UnaryExprBuilder.md)
 * <span class="badge builder"></span> [VectorExprBuilder](./builder-VectorExprBuilder.md)
## Functions

### <span class="badge function"></span> n

Shortcut to turn a number into a NumberLiteralExpr expression.

```typescript
n(value: number)
```

### <span class="badge function"></span> s

Shortcut to turn a string into a StringLiteralExpr expression.

```typescript
s(value: string)
```

### <span class="badge function"></span> subquery

Creates a subquery.

Subquery allows you to run an instant query for a given range and resolution. The result of a subquery is a range vector.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery

```typescript
subquery(expression: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sum

Calculate sum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
sum(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> min

Calculate minimum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
min(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> max

Calculate maximum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
max(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> avg

Calculate the average over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
avg(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> group

All values in the resulting vector are 1.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
group(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> stddev

Calculate population standard deviation over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
stddev(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> stdvar

Calculate population standard variance over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
stdvar(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> count

Count number of elements in the vector.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
count(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> quantile

Calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
quantile(vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> countValues

Count number of elements with the same value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
countValues(label: string, vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> bottomk

Smallest k elements by sample value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
bottomk(k: number, vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> topk

Largest k elements by sample value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
topk(k: number, vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> limitk

Sample k elements.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
limitk(k: number, vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> limitRatio

Sample elements with approximately r ratio if r > 0, and the complement of such samples if r = -(1.0 - r).

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```typescript
limitRatio(k: number, vector: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> vector

Returns the scalar s as a vector with no labels.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector

```typescript
vector(s: string)
```

### <span class="badge function"></span> add

Addition binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```typescript
add(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sub

Subtraction binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```typescript
sub(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> mul

Multiplication binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```typescript
mul(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> div

Division binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```typescript
div(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> mod

Modulo binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```typescript
mod(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> pow

Power/exponentiation binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```typescript
pow(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> eq

"equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```typescript
eq(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> neq

"not-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```typescript
neq(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> gt

"greater-than" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```typescript
gt(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> lt

"less-than" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```typescript
lt(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> gte

"greater-or-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```typescript
gte(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> lte

"less-or-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```typescript
lte(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> and

"intersection" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```typescript
and(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> or

"union" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```typescript
or(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> unless

"complement" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```typescript
unless(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> atan2

Arc tangent binary operator. Works in radians.

Trigonometric operators allow trigonometric functions to be executed on two vectors using vector matching, which isn't available with normal functions.

They act in the same manner as arithmetic operators.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#trigonometric-binary-operators

```typescript
atan2(left: cog.Builder<promql.Expr>, right: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> neg

Negation unary operator.

```typescript
neg(expr: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> id

Identity unary operator.

```typescript
id(expr: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> abs

Returns the input vector with all sample values converted to their absolute value.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#abs

```typescript
abs(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> absent

Returns an empty vector if the vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the vector passed to it has no elements.

This is useful for alerting on when no time series exist for a given metric name and label combination.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent

```typescript
absent(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> absentOverTime

Returns an empty vector if the range vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the range vector passed to it has no elements.

This is useful for alerting on when no time series exist for a given metric name and label combination for a certain amount of time.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent_over_time

```typescript
absentOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> ceil

Rounds the sample values of all elements in `v` up to the nearest integer value greater than or equal to v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#ceil

```typescript
ceil(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> changes

For each input time series, returns the number of times its value has changed within the provided time range as an instant vector.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#changes

```typescript
changes(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> clamp

Clamps the sample values of all elements in `v` to have a lower limit of min and an upper limit of max.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp

```typescript
clamp(v: cog.Builder<promql.Expr>, min: number, max: number)
```

### <span class="badge function"></span> clampMax

Clamps the sample values of all elements in `v` to have an upper limit of `max`.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_max

```typescript
clampMax(v: cog.Builder<promql.Expr>, max: number)
```

### <span class="badge function"></span> clampMin

Clamps the sample values of all elements in `v` to have an lower limit of `min`.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_min

```typescript
clampMin(v: cog.Builder<promql.Expr>, min: number)
```

### <span class="badge function"></span> dayOfMonth

Returns the day of the month in UTC. Returned values are from 1 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month

```typescript
dayOfMonth()
```

### <span class="badge function"></span> dayOfMonthFor

Returns the day of the month for each of the given times in UTC. Returned values are from 1 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month

```typescript
dayOfMonthFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> dayOfWeek

Returns the day of the week in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week

```typescript
dayOfWeek()
```

### <span class="badge function"></span> dayOfWeekFor

Returns the day of the week for each of the given times in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week

```typescript
dayOfWeekFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> dayOfYear

Returns the day of the year in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year

```typescript
dayOfYear()
```

### <span class="badge function"></span> dayOfYearFor

Returns the day of the year for each of the given times in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year

```typescript
dayOfYearFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> daysInMonth

Returns the number of days in the month. Returned values are from 28 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month

```typescript
daysInMonth()
```

### <span class="badge function"></span> dayInMonthFor

Returns the number of days in the month for each of the given times in UTC. Returned values are from 28 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month

```typescript
dayInMonthFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> delta

Calculates the difference between the first and last value of each time series element in a range vector, returning an instant vector with the given deltas and equivalent labels.

The delta is extrapolated to cover the full time range as specified in the range vector selector, so that it is possible to get a non-integer result even if the sample values are all integers.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#delta

```typescript
delta(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> deriv

Calculates the per-second derivative of the time series in a range vector using simple linear regression.

The range vector must have at least two samples in order to perform the calculation. When +Inf or -Inf are found in the range vector, the slope and offset value calculated will be NaN.

deriv should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv

```typescript
deriv(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> exp

Calculates the exponential function for all elements in vector

See https://prometheus.io/docs/prometheus/latest/querying/functions/#exp

```typescript
exp(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> floor

Rounds the sample values of all elements in v down to the nearest integer value smaller than or equal to v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#floor

```typescript
floor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramAvg

Returns the arithmetic average of observed values stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_avg

```typescript
histogramAvg(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramCount

Returns the count of observations stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum

```typescript
histogramCount(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramSum

Returns the sum of observations stored in a native histogram.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum

```typescript
histogramSum(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramFraction

Returns the estimated fraction of observations between the provided lower and upper values. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_fraction

```typescript
histogramFraction(lower: number, upper: number, v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramQuantile

Calculates the φ-quantile (0 ≤ φ ≤ 1) from a classic histogram or from a native histogram.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile

```typescript
histogramQuantile(phi: number, v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramStddev

Returns the estimated standard deviation of observations in a native histogram, based on the geometric mean of the buckets where the observations lie.

Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stddev

```typescript
histogramStddev(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> histogramStdvar

Returns the estimated standard variance of observations in a native histogram.

Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stdvar

```typescript
histogramStdvar(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> hour

Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour

```typescript
hour()
```

### <span class="badge function"></span> hourFor

Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour

```typescript
hourFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> idelta

Calculates the difference between the last two samples in the range vector v, returning an instant vector with the given deltas and equivalent labels.

idelta should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta

```typescript
idelta(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> increase

Calculates the increase in the time series in the range vector

See https://prometheus.io/docs/prometheus/latest/querying/functions/#increase

```typescript
increase(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> irate

Calculates the per-second instant rate of increase of the time series in the range vector. This is based on the last two data points.

irate should only be used when graphing volatile, fast-moving counters. Use rate for alerts and slow-moving counters, as brief changes in the rate can reset the FOR clause and graphs consisting entirely of rare spikes are hard to read.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#irate

```typescript
irate(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> labelReplace

matches the regular expression regex against the value of the label src_label. If it matches, the value of the label dst_label in the returned timeseries will be the expansion of replacement, together with the original labels in the input. Capturing groups in the regular expression can be referenced with $1, $2, etc. Named capturing groups in the regular expression can be referenced with $name (where name is the capturing group name). If the regular expression doesn't match then the timeseries is returned unchanged.

label_replace acts on float and histogram samples in the same way.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace

```typescript
labelReplace(v: cog.Builder<promql.Expr>, dst_label: string, replacement: string, src_label: string, regex: string)
```

### <span class="badge function"></span> ln

Calculates the natural logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#ln

```typescript
ln(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> log2

Calculates the binary logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#log2

```typescript
log2(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> log10

Calculates the decimal logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#log10

```typescript
log10(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> minute

Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute

```typescript
minute()
```

### <span class="badge function"></span> minuteFor

Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute

```typescript
minuteFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> month

Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#month

```typescript
month()
```

### <span class="badge function"></span> monthFor

Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#month

```typescript
monthFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> predictLinear

Predicts the value of time series t seconds from now, based on the range vector v, using simple linear regression. The range vector must have at least two samples in order to perform the calculation.

predict_linear should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear

```typescript
predictLinear(v: cog.Builder<promql.Expr>, t: number)
```

### <span class="badge function"></span> rate

Calculates the per-second average rate of increase of the time series in the range vector.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#rate

```typescript
rate(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> resets

For each input time series, resets(v range-vector) returns the number of counter resets within the provided time range as an instant vector. Any decrease in the value between two consecutive float samples is interpreted as a counter reset. A reset in a native histogram is detected in a more complex way: Any decrease in any bucket, including the zero bucket, or in the count of observation constitutes a counter reset, but also the disappearance of any previously populated bucket, an increase in bucket resolution, or a decrease of the zero-bucket width.

`resets` should only be used with counters and counter-like native histograms.

If the range vector contains a mix of float and histogram samples for the same series, counter resets are detected separately and their numbers added up. The change from a float to a histogram sample is not considered a counter reset. Each float sample is compared to the next float sample, and each histogram is compared to the next histogram.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#resets

```typescript
resets(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> round

Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#round

```typescript
round(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> roundTo

Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.

The to_nearest argument allows specifying the nearest multiple to which the sample values should be rounded. This multiple may also be a fraction.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#round

```typescript
roundTo(v: cog.Builder<promql.Expr>, to_nearest: number)
```

### <span class="badge function"></span> scalar

Given a single-element input vector, scalar() returns the sample value of that single element as a scalar.

If the input vector does not have exactly one element, scalar will return NaN.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#scalar

```typescript
scalar(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sgn

Returns a vector with all sample values converted to their sign, defined as this: 1 if v is positive, -1 if v is negative and 0 if v is equal to zero.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sgn

```typescript
sgn(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sort

Returns vector elements sorted by their sample values, in ascending order. Native histograms are sorted by their sum of observations.

Note that sort only affects the results of instant queries, as range query results always have a fixed output ordering.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort

```typescript
sort(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sortDesc

Same as `sort()`, but sorts in descending order.

Like sort, sort_desc only affects the results of instant queries, as range query results always have a fixed output ordering.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort_desc

```typescript
sortDesc(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sqrt

Calculates the square root of all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sqrt

```typescript
sqrt(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> time

Returns the number of seconds since January 1, 1970 UTC. Note that this does not actually return the current time, but the time at which the expression is to be evaluated.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#time

```typescript
time()
```

### <span class="badge function"></span> timestamp

Returns the timestamp of each of the samples of the given vector as the number of seconds since January 1, 1970 UTC. It also works with histogram samples.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp

```typescript
timestamp(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> vect

Returns the scalar s as a vector with no labels.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector

```typescript
vect(s: number)
```

### <span class="badge function"></span> year

Returns the year for each of the given times in UTC.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#year

```typescript
year()
```

### <span class="badge function"></span> yearFor

Returns the year for each of the given times in UTC.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#year

```typescript
yearFor(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> avgOverTime

Calculates average value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
avgOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> minOverTime

Calculates the minimum value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
minOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> maxOverTime

Calculates the maximum value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
maxOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sumOverTime

Calculates the sum of all values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
sumOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> countOverTime

Calculates the count of all values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
countOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> quantileOverTime

Calculates the φ-quantile (0 ≤ φ ≤ 1) of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
quantileOverTime(phi: number, v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> stddevOverTime

Calculates the population standard deviation of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
stddevOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> stdvarOverTime

Calculates the population standard variance of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
stdvarOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> lastOverTime

Returns the most recent point value in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
lastOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> presentOverTime

Returns the value 1 for any series in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```typescript
presentOverTime(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> acos

Calculates the arccosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
acos(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> acosh

Calculates the inverse hyperbolic cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
acosh(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> asin

Calculates the arcsine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
asin(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> asinh

Calculates the inverse hyperbolic sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
asinh(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> atan

Calculates the arctangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
atan(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> atanh

Calculates the inverse hyperbolic tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
atanh(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> cos

Calculates the cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
cos(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> cosh

Calculates the hyperbolic cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
cosh(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sin

Calculates the sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
sin(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> sinh

Calculates the hyperbolic sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
sinh(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> tan

Calculates the tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
tan(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> tanh

Calculates the hyperbolic tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
tanh(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> deg

Converts radians to degrees for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
deg(v: cog.Builder<promql.Expr>)
```

### <span class="badge function"></span> pi

Returns pi.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
pi()
```

### <span class="badge function"></span> rad

Converts degrees to radians for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```typescript
rad(v: cog.Builder<promql.Expr>)
```

