# promql

## Objects

 * <span class="badge object-type-class"></span> [AggregationExpr](./object-AggregationExpr.md)
 * <span class="badge object-type-enum"></span> [AggregationOp](./object-AggregationOp.md)
 * <span class="badge object-type-class"></span> [BinaryExpr](./object-BinaryExpr.md)
 * <span class="badge object-type-enum"></span> [BinaryOp](./object-BinaryOp.md)
 * <span class="badge object-type-disjunction"></span> [Expr](./object-Expr.md)
 * <span class="badge object-type-class"></span> [FuncCallExpr](./object-FuncCallExpr.md)
 * <span class="badge object-type-enum"></span> [LabelMatchingOperator](./object-LabelMatchingOperator.md)
 * <span class="badge object-type-class"></span> [LabelSelector](./object-LabelSelector.md)
 * <span class="badge object-type-class"></span> [NumberLiteralExpr](./object-NumberLiteralExpr.md)
 * <span class="badge object-type-class"></span> [StringLiteralExpr](./object-StringLiteralExpr.md)
 * <span class="badge object-type-class"></span> [SubqueryExpr](./object-SubqueryExpr.md)
 * <span class="badge object-type-class"></span> [UnaryExpr](./object-UnaryExpr.md)
 * <span class="badge object-type-enum"></span> [UnaryOp](./object-UnaryOp.md)
 * <span class="badge object-type-class"></span> [VectorExpr](./object-VectorExpr.md)
## Builders

 * <span class="badge builder"></span> [AggregationExpr](./builder-AggregationExpr.md)
 * <span class="badge builder"></span> [BinaryExpr](./builder-BinaryExpr.md)
 * <span class="badge builder"></span> [FuncCallExpr](./builder-FuncCallExpr.md)
 * <span class="badge builder"></span> [LabelSelector](./builder-LabelSelector.md)
 * <span class="badge builder"></span> [NumberLiteralExpr](./builder-NumberLiteralExpr.md)
 * <span class="badge builder"></span> [StringLiteralExpr](./builder-StringLiteralExpr.md)
 * <span class="badge builder"></span> [SubqueryExpr](./builder-SubqueryExpr.md)
 * <span class="badge builder"></span> [UnaryExpr](./builder-UnaryExpr.md)
 * <span class="badge builder"></span> [VectorExpr](./builder-VectorExpr.md)
## Functions

### <span class="badge function"></span> n

Shortcut to turn a number into a NumberLiteralExpr expression.

```python
def n(value: float) -> NumberLiteralExpr
```

### <span class="badge function"></span> s

Shortcut to turn a string into a StringLiteralExpr expression.

```python
def s(value: str) -> StringLiteralExpr
```

### <span class="badge function"></span> subquery

Creates a subquery.

Subquery allows you to run an instant query for a given range and resolution. The result of a subquery is a range vector.

See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery

```python
def subquery(expression: cogbuilder.Builder[promql.Expr]) -> SubqueryExpr
```

### <span class="badge function"></span> sum

Calculate sum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def sum(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> min

Calculate minimum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def min(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> max

Calculate maximum over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def max(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> avg

Calculate the average over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def avg(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> group

All values in the resulting vector are 1.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def group(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> stddev

Calculate population standard deviation over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def stddev(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> stdvar

Calculate population standard variance over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def stdvar(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> count

Count number of elements in the vector.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def count(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> quantile

Calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def quantile(vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> count_values

Count number of elements with the same value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def count_values(label: str, vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> bottomk

Smallest k elements by sample value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def bottomk(k: float, vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> topk

Largest k elements by sample value.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def topk(k: float, vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> limitk

Sample k elements.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def limitk(k: float, vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> limit_ratio

Sample elements with approximately r ratio if r > 0, and the complement of such samples if r = -(1.0 - r).

See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators

```python
def limit_ratio(k: float, vector: cogbuilder.Builder[promql.Expr]) -> AggregationExpr
```

### <span class="badge function"></span> vector

Returns the scalar s as a vector with no labels.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector

```python
def vector(s: str) -> VectorExpr
```

### <span class="badge function"></span> add

Addition binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```python
def add(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> sub

Subtraction binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```python
def sub(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> mul

Multiplication binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```python
def mul(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> div

Division binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```python
def div(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> mod

Modulo binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```python
def mod(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> pow

Power/exponentiation binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators

```python
def pow(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> eq

"equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```python
def eq(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> neq

"not-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```python
def neq(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> gt

"greater-than" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```python
def gt(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> lt

"less-than" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```python
def lt(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> gte

"greater-or-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```python
def gte(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> lte

"less-or-equal" comparison binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators

```python
def lte(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> and_val

"intersection" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```python
def and_val(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> or_val

"union" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```python
def or_val(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> unless

"complement" logical/set binary operator.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators

```python
def unless(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> atan2

Arc tangent binary operator. Works in radians.

Trigonometric operators allow trigonometric functions to be executed on two vectors using vector matching, which isn't available with normal functions.

They act in the same manner as arithmetic operators.

See https://prometheus.io/docs/prometheus/latest/querying/operators/#trigonometric-binary-operators

```python
def atan2(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]) -> BinaryExpr
```

### <span class="badge function"></span> neg

Negation unary operator.

```python
def neg(expr: cogbuilder.Builder[promql.Expr]) -> UnaryExpr
```

### <span class="badge function"></span> id

Identity unary operator.

```python
def id(expr: cogbuilder.Builder[promql.Expr]) -> UnaryExpr
```

### <span class="badge function"></span> abs

Returns the input vector with all sample values converted to their absolute value.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#abs

```python
def abs(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> absent

Returns an empty vector if the vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the vector passed to it has no elements.

This is useful for alerting on when no time series exist for a given metric name and label combination.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent

```python
def absent(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> absent_over_time

Returns an empty vector if the range vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the range vector passed to it has no elements.

This is useful for alerting on when no time series exist for a given metric name and label combination for a certain amount of time.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent_over_time

```python
def absent_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> ceil

Rounds the sample values of all elements in `v` up to the nearest integer value greater than or equal to v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#ceil

```python
def ceil(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> changes

For each input time series, returns the number of times its value has changed within the provided time range as an instant vector.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#changes

```python
def changes(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> clamp

Clamps the sample values of all elements in `v` to have a lower limit of min and an upper limit of max.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp

```python
def clamp(v: cogbuilder.Builder[promql.Expr], min: float, max: float) -> FuncCallExpr
```

### <span class="badge function"></span> clamp_max

Clamps the sample values of all elements in `v` to have an upper limit of `max`.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_max

```python
def clamp_max(v: cogbuilder.Builder[promql.Expr], max: float) -> FuncCallExpr
```

### <span class="badge function"></span> clamp_min

Clamps the sample values of all elements in `v` to have an lower limit of `min`.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_min

```python
def clamp_min(v: cogbuilder.Builder[promql.Expr], min: float) -> FuncCallExpr
```

### <span class="badge function"></span> day_of_month

Returns the day of the month in UTC. Returned values are from 1 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month

```python
def day_of_month() -> FuncCallExpr
```

### <span class="badge function"></span> day_of_month_for

Returns the day of the month for each of the given times in UTC. Returned values are from 1 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month

```python
def day_of_month_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> day_of_week

Returns the day of the week in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week

```python
def day_of_week() -> FuncCallExpr
```

### <span class="badge function"></span> day_of_week_for

Returns the day of the week for each of the given times in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week

```python
def day_of_week_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> day_of_year

Returns the day of the year in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year

```python
def day_of_year() -> FuncCallExpr
```

### <span class="badge function"></span> day_of_year_for

Returns the day of the year for each of the given times in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year

```python
def day_of_year_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> days_in_month

Returns the number of days in the month. Returned values are from 28 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month

```python
def days_in_month() -> FuncCallExpr
```

### <span class="badge function"></span> day_in_month_for

Returns the number of days in the month for each of the given times in UTC. Returned values are from 28 to 31.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month

```python
def day_in_month_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> delta

Calculates the difference between the first and last value of each time series element in a range vector, returning an instant vector with the given deltas and equivalent labels.

The delta is extrapolated to cover the full time range as specified in the range vector selector, so that it is possible to get a non-integer result even if the sample values are all integers.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#delta

```python
def delta(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> deriv

Calculates the per-second derivative of the time series in a range vector using simple linear regression.

The range vector must have at least two samples in order to perform the calculation. When +Inf or -Inf are found in the range vector, the slope and offset value calculated will be NaN.

deriv should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv

```python
def deriv(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> exp

Calculates the exponential function for all elements in vector

See https://prometheus.io/docs/prometheus/latest/querying/functions/#exp

```python
def exp(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> floor

Rounds the sample values of all elements in v down to the nearest integer value smaller than or equal to v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#floor

```python
def floor(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_avg

Returns the arithmetic average of observed values stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_avg

```python
def histogram_avg(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_count

Returns the count of observations stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum

```python
def histogram_count(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_sum

Returns the sum of observations stored in a native histogram.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum

```python
def histogram_sum(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_fraction

Returns the estimated fraction of observations between the provided lower and upper values. Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_fraction

```python
def histogram_fraction(lower: float, upper: float, v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_quantile

Calculates the φ-quantile (0 ≤ φ ≤ 1) from a classic histogram or from a native histogram.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile

```python
def histogram_quantile(phi: float, v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_stddev

Returns the estimated standard deviation of observations in a native histogram, based on the geometric mean of the buckets where the observations lie.

Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stddev

```python
def histogram_stddev(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> histogram_stdvar

Returns the estimated standard variance of observations in a native histogram.

Samples that are not native histograms are ignored and do not show up in the returned vector.

Note: This function only acts on native histograms.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stdvar

```python
def histogram_stdvar(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> hour

Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour

```python
def hour() -> FuncCallExpr
```

### <span class="badge function"></span> hour_for

Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour

```python
def hour_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> idelta

Calculates the difference between the last two samples in the range vector v, returning an instant vector with the given deltas and equivalent labels.

idelta should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta

```python
def idelta(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> increase

Calculates the increase in the time series in the range vector

See https://prometheus.io/docs/prometheus/latest/querying/functions/#increase

```python
def increase(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> irate

Calculates the per-second instant rate of increase of the time series in the range vector. This is based on the last two data points.

irate should only be used when graphing volatile, fast-moving counters. Use rate for alerts and slow-moving counters, as brief changes in the rate can reset the FOR clause and graphs consisting entirely of rare spikes are hard to read.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#irate

```python
def irate(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> label_replace

matches the regular expression regex against the value of the label src_label. If it matches, the value of the label dst_label in the returned timeseries will be the expansion of replacement, together with the original labels in the input. Capturing groups in the regular expression can be referenced with $1, $2, etc. Named capturing groups in the regular expression can be referenced with $name (where name is the capturing group name). If the regular expression doesn't match then the timeseries is returned unchanged.

label_replace acts on float and histogram samples in the same way.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace

```python
def label_replace(v: cogbuilder.Builder[promql.Expr], dst_label: str, replacement: str, src_label: str, regex: str) -> FuncCallExpr
```

### <span class="badge function"></span> ln

Calculates the natural logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#ln

```python
def ln(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> log2

Calculates the binary logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#log2

```python
def log2(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> log10

Calculates the decimal logarithm for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#log10

```python
def log10(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> minute

Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute

```python
def minute() -> FuncCallExpr
```

### <span class="badge function"></span> minute_for

Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute

```python
def minute_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> month

Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#month

```python
def month() -> FuncCallExpr
```

### <span class="badge function"></span> month_for

Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#month

```python
def month_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> predict_linear

Predicts the value of time series t seconds from now, based on the range vector v, using simple linear regression. The range vector must have at least two samples in order to perform the calculation.

predict_linear should only be used with gauges.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear

```python
def predict_linear(v: cogbuilder.Builder[promql.Expr], t: float) -> FuncCallExpr
```

### <span class="badge function"></span> rate

Calculates the per-second average rate of increase of the time series in the range vector.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#rate

```python
def rate(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> resets

For each input time series, resets(v range-vector) returns the number of counter resets within the provided time range as an instant vector. Any decrease in the value between two consecutive float samples is interpreted as a counter reset. A reset in a native histogram is detected in a more complex way: Any decrease in any bucket, including the zero bucket, or in the count of observation constitutes a counter reset, but also the disappearance of any previously populated bucket, an increase in bucket resolution, or a decrease of the zero-bucket width.

`resets` should only be used with counters and counter-like native histograms.

If the range vector contains a mix of float and histogram samples for the same series, counter resets are detected separately and their numbers added up. The change from a float to a histogram sample is not considered a counter reset. Each float sample is compared to the next float sample, and each histogram is compared to the next histogram.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#resets

```python
def resets(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> round

Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#round

```python
def round(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> round_to

Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.

The to_nearest argument allows specifying the nearest multiple to which the sample values should be rounded. This multiple may also be a fraction.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#round

```python
def round_to(v: cogbuilder.Builder[promql.Expr], to_nearest: float) -> FuncCallExpr
```

### <span class="badge function"></span> scalar

Given a single-element input vector, scalar() returns the sample value of that single element as a scalar.

If the input vector does not have exactly one element, scalar will return NaN.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#scalar

```python
def scalar(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sgn

Returns a vector with all sample values converted to their sign, defined as this: 1 if v is positive, -1 if v is negative and 0 if v is equal to zero.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sgn

```python
def sgn(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sort

Returns vector elements sorted by their sample values, in ascending order. Native histograms are sorted by their sum of observations.

Note that sort only affects the results of instant queries, as range query results always have a fixed output ordering.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort

```python
def sort(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sort_desc

Same as `sort()`, but sorts in descending order.

Like sort, sort_desc only affects the results of instant queries, as range query results always have a fixed output ordering.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort_desc

```python
def sort_desc(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sqrt

Calculates the square root of all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#sqrt

```python
def sqrt(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> time

Returns the number of seconds since January 1, 1970 UTC. Note that this does not actually return the current time, but the time at which the expression is to be evaluated.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#time

```python
def time() -> FuncCallExpr
```

### <span class="badge function"></span> timestamp

Returns the timestamp of each of the samples of the given vector as the number of seconds since January 1, 1970 UTC. It also works with histogram samples.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp

```python
def timestamp(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> vect

Returns the scalar s as a vector with no labels.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector

```python
def vect(s: float) -> FuncCallExpr
```

### <span class="badge function"></span> year

Returns the year for each of the given times in UTC.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#year

```python
def year() -> FuncCallExpr
```

### <span class="badge function"></span> year_for

Returns the year for each of the given times in UTC.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#year

```python
def year_for(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> avg_over_time

Calculates average value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def avg_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> min_over_time

Calculates the minimum value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def min_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> max_over_time

Calculates the maximum value of all points in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def max_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sum_over_time

Calculates the sum of all values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def sum_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> count_over_time

Calculates the count of all values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def count_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> quantile_over_time

Calculates the φ-quantile (0 ≤ φ ≤ 1) of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def quantile_over_time(phi: float, v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> stddev_over_time

Calculates the population standard deviation of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def stddev_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> stdvar_over_time

Calculates the population standard variance of the values in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def stdvar_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> last_over_time

Returns the most recent point value in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def last_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> present_over_time

Returns the value 1 for any series in the specified interval.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time

```python
def present_over_time(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> acos

Calculates the arccosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def acos(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> acosh

Calculates the inverse hyperbolic cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def acosh(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> asin

Calculates the arcsine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def asin(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> asinh

Calculates the inverse hyperbolic sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def asinh(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> atan

Calculates the arctangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def atan(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> atanh

Calculates the inverse hyperbolic tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def atanh(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> cos

Calculates the cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def cos(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> cosh

Calculates the hyperbolic cosine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def cosh(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sin

Calculates the sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def sin(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> sinh

Calculates the hyperbolic sine of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def sinh(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> tan

Calculates the tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def tan(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> tanh

Calculates the hyperbolic tangent of all elements in v

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def tanh(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> deg

Converts radians to degrees for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def deg(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

### <span class="badge function"></span> pi

Returns pi.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def pi() -> FuncCallExpr
```

### <span class="badge function"></span> rad

Converts degrees to radians for all elements in v.

See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions

```python
def rad(v: cogbuilder.Builder[promql.Expr]) -> FuncCallExpr
```

