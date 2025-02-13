# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import promql


class NumberLiteralExpr(cogbuilder.Builder[promql.NumberLiteralExpr]):    
    """
    Represents a number literal expression.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
    """
    
    _internal: promql.NumberLiteralExpr

    def __init__(self):
        self._internal = promql.NumberLiteralExpr()        
        self._internal.type_val = "numberLiteralExpr"

    def build(self) -> promql.NumberLiteralExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def value(self, value: float) -> typing.Self:    
        self._internal.value = value
    
        return self
    

"""
Shortcut to turn a number into a NumberLiteralExpr expression.
"""
def n(value: float):
    builder = NumberLiteralExpr()
    builder.value(value)

    return builder


class StringLiteralExpr(cogbuilder.Builder[promql.StringLiteralExpr]):    
    """
    Represents a string literal expression.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
    """
    
    _internal: promql.StringLiteralExpr

    def __init__(self):
        self._internal = promql.StringLiteralExpr()        
        self._internal.type_val = "stringLiteralExpr"

    def build(self) -> promql.StringLiteralExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    

"""
Shortcut to turn a string into a StringLiteralExpr expression.
"""
def s(value: str):
    builder = StringLiteralExpr()
    builder.value(value)

    return builder


class SubqueryExpr(cogbuilder.Builder[promql.SubqueryExpr]):    
    """
    Represents a subquery.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
    """
    
    _internal: promql.SubqueryExpr

    def __init__(self):
        self._internal = promql.SubqueryExpr()        
        self._internal.type_val = "subqueryExpr"

    def build(self) -> promql.SubqueryExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def expr(self, expr: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        expr_resource = expr.build()
        self._internal.expr = expr_resource
    
        return self
    
    def offset(self, offset: str) -> typing.Self:    
        """
        The offset modifier allows changing the time offset for individual instant and range vectors in a query.
        https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
        """
            
        self._internal.offset = offset
    
        return self
    
    def at(self, at: str) -> typing.Self:    
        """
        The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
        The time supplied to the @ modifier is a unix timestamp.
        https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
        """
            
        self._internal.at = at
    
        return self
    
    def range(self, range_val: str) -> typing.Self:    
        """
        Range of samples back from the current instant.
        https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
        """
            
        self._internal.range_val = range_val
    
        return self
    
    def resolution(self, resolution: str) -> typing.Self:    
        """
        Empty string for default resolution.
        """
            
        self._internal.resolution = resolution
    
        return self
    

"""
Creates a subquery.
Subquery allows you to run an instant query for a given range and resolution. The result of a subquery is a range vector.
See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
"""
def subquery(expression: cogbuilder.Builder[promql.Expr]):
    builder = SubqueryExpr()
    builder.expr(expression)

    return builder


class AggregationExpr(cogbuilder.Builder[promql.AggregationExpr]):    
    """
    Represents an aggregation.
    See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
    """
    
    _internal: promql.AggregationExpr

    def __init__(self):
        self._internal = promql.AggregationExpr()        
        self._internal.type_val = "aggregationExpr"

    def build(self) -> promql.AggregationExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def op(self, op: promql.AggregationOp) -> typing.Self:    
        self._internal.op = op
    
        return self
    
    def expr(self, expr: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        expr_resource = expr.build()
        self._internal.expr = expr_resource
    
        return self
    
    def param(self, param: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        param_resource = param.build()
        self._internal.param = param_resource
    
        return self
    
    def by(self, by: list[str]) -> typing.Self:    
        """
        By drops labels that are not listed in the by clause.
        """
            
        self._internal.by = by
    
        return self
    
    def without(self, without: list[str]) -> typing.Self:    
        """
        List of labels to remove from the result vector, while all other labels are preserved in the output.
        """
            
        self._internal.without = without
    
        return self
    

"""
Calculate sum over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def sum(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.SUM)
    builder.expr(vector)

    return builder

"""
Calculate minimum over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def min(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.MIN)
    builder.expr(vector)

    return builder

"""
Calculate maximum over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def max(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.MAX)
    builder.expr(vector)

    return builder

"""
Calculate the average over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def avg(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.AVG)
    builder.expr(vector)

    return builder

"""
All values in the resulting vector are 1.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def group(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.GROUP)
    builder.expr(vector)

    return builder

"""
Calculate population standard deviation over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def stddev(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.STDDEV)
    builder.expr(vector)

    return builder

"""
Calculate population standard variance over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def stdvar(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.STDVAR)
    builder.expr(vector)

    return builder

"""
Count number of elements in the vector.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def count(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.COUNT)
    builder.expr(vector)

    return builder

"""
Calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def quantile(vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.QUANTILE)
    builder.expr(vector)

    return builder

"""
Count number of elements with the same value.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def count_values(label: str, vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.COUNT_VALUES)
    builder.expr(vector)
    builder.param(s(label))

    return builder

"""
Smallest k elements by sample value.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def bottomk(k: float, vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.BOTTOMK)
    builder.expr(vector)
    builder.param(n(k))

    return builder

"""
Largest k elements by sample value.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def topk(k: float, vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.TOPK)
    builder.expr(vector)
    builder.param(n(k))

    return builder

"""
Sample k elements.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def limitk(k: float, vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.LIMITK)
    builder.expr(vector)
    builder.param(n(k))

    return builder

"""
Sample elements with approximately r ratio if r > 0, and the complement of such samples if r = -(1.0 - r).
See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
"""
def limit_ratio(k: float, vector: cogbuilder.Builder[promql.Expr]):
    builder = AggregationExpr()
    builder.op(promql.AggregationOp.LIMIT_RATIO)
    builder.expr(vector)
    builder.param(n(k))

    return builder


class VectorExpr(cogbuilder.Builder[promql.VectorExpr]):    
    """
    Represents both instant and range vectors
    """
    
    _internal: promql.VectorExpr

    def __init__(self):
        self._internal = promql.VectorExpr()        
        self._internal.type_val = "vectorExpr"

    def build(self) -> promql.VectorExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def metric(self, metric: str) -> typing.Self:    
        """
        Metric name.
        """
            
        self._internal.metric = metric
    
        return self
    
    def labels(self, labels: list[cogbuilder.Builder[promql.LabelSelector]]) -> typing.Self:    
        """
        Label selectors used to filter the timeseries.
        """
            
        labels_resources = [r1.build() for r1 in labels]
        self._internal.labels = labels_resources
    
        return self
    
    def offset(self, offset: str) -> typing.Self:    
        """
        The offset modifier allows changing the time offset for individual instant and range vectors in a query.
        https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
        """
            
        self._internal.offset = offset
    
        return self
    
    def at(self, at: str) -> typing.Self:    
        """
        The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
        The time supplied to the @ modifier is a unix timestamp.
        https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
        """
            
        self._internal.at = at
    
        return self
    
    def range(self, range_val: str) -> typing.Self:    
        """
        Range of samples back from the current instant.
        https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
        """
            
        self._internal.range_val = range_val
    
        return self
    
    def label(self, name: str, value: str) -> typing.Self:    
        if self._internal.labels is None:
            self._internal.labels = []
        
        self._internal.labels.append(promql.LabelSelector(
            name=name,
            operator=promql.LabelMatchingOperator.EQUAL,
            value=value,
        ))
    
        return self
    
    def label_neq(self, name: str, value: str) -> typing.Self:    
        if self._internal.labels is None:
            self._internal.labels = []
        
        self._internal.labels.append(promql.LabelSelector(
            name=name,
            operator=promql.LabelMatchingOperator.NOT_EQUAL,
            value=value,
        ))
    
        return self
    
    def label_match_regexp(self, name: str, value: str) -> typing.Self:    
        if self._internal.labels is None:
            self._internal.labels = []
        
        self._internal.labels.append(promql.LabelSelector(
            name=name,
            operator=promql.LabelMatchingOperator.MATCH_REGEXP,
            value=value,
        ))
    
        return self
    
    def label_not_match_regexp(self, name: str, value: str) -> typing.Self:    
        if self._internal.labels is None:
            self._internal.labels = []
        
        self._internal.labels.append(promql.LabelSelector(
            name=name,
            operator=promql.LabelMatchingOperator.NOT_MATCH_REGEXP,
            value=value,
        ))
    
        return self
    

"""
Returns the scalar s as a vector with no labels.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector
"""
def vector(s: str):
    builder = VectorExpr()
    builder.metric(s)

    return builder


class LabelSelector(cogbuilder.Builder[promql.LabelSelector]):
    _internal: promql.LabelSelector

    def __init__(self):
        self._internal = promql.LabelSelector()

    def build(self) -> promql.LabelSelector:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def name(self, name: str) -> typing.Self:    
        """
        Name of the label to select.
        """
            
        if not len(name) >= 1:
            raise ValueError("len(name) must be >= 1")
        self._internal.name = name
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        """
        Value to match against.
        """
            
        self._internal.value = value
    
        return self
    
    def operator(self, operator: promql.LabelMatchingOperator) -> typing.Self:    
        """
        Operator used to perform the selection.
        """
            
        self._internal.operator = operator
    
        return self
    


class BinaryExpr(cogbuilder.Builder[promql.BinaryExpr]):    
    """
    Represents a binary operation expression.
    """
    
    _internal: promql.BinaryExpr

    def __init__(self):
        self._internal = promql.BinaryExpr()        
        self._internal.type_val = "binaryExpr"

    def build(self) -> promql.BinaryExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def op(self, op: promql.BinaryOp) -> typing.Self:    
        self._internal.op = op
    
        return self
    
    def left(self, left: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        left_resource = left.build()
        self._internal.left = left_resource
    
        return self
    
    def right(self, right: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        right_resource = right.build()
        self._internal.right = right_resource
    
        return self
    
    def ignoring(self, labels: list[str]) -> typing.Self:    
        """
        Allows ignoring certain labels when matching.
        See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches
        """
            
        self._internal.match_type = "ignore"    
        self._internal.match_labels = labels
    
        return self
    
    def on(self, labels: list[str]) -> typing.Self:    
        """
        Allows reducing the set of considered labels to a provided list when matching.
        See https://prometheus.io/docs/prometheus/latest/querying/operators/#one-to-one-vector-matches
        """
            
        self._internal.match_type = "on"    
        self._internal.match_labels = labels
    
        return self
    
    def group_left(self, labels: list[str]) -> typing.Self:    
        """
        See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches
        """
            
        self._internal.group_modifier = "left"    
        self._internal.group_labels = labels
    
        return self
    
    def group_right(self, labels: list[str]) -> typing.Self:    
        """
        See https://prometheus.io/docs/prometheus/latest/querying/operators/#many-to-one-and-one-to-many-vector-matches
        """
            
        self._internal.group_modifier = "right"    
        self._internal.group_labels = labels
    
        return self
    

"""
Addition binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
"""
def add(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.ADD)
    builder.left(left)
    builder.right(right)

    return builder

"""
Subtraction binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
"""
def sub(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.SUB)
    builder.left(left)
    builder.right(right)

    return builder

"""
Multiplication binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
"""
def mul(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.MUL)
    builder.left(left)
    builder.right(right)

    return builder

"""
Division binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
"""
def div(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.DIV)
    builder.left(left)
    builder.right(right)

    return builder

"""
Modulo binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
"""
def mod(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.MOD)
    builder.left(left)
    builder.right(right)

    return builder

"""
Power/exponentiation binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#arithmetic-binary-operators
"""
def pow(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.POW)
    builder.left(left)
    builder.right(right)

    return builder

"""
"equal" comparison binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
"""
def eq(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.EQL)
    builder.left(left)
    builder.right(right)

    return builder

"""
"not-equal" comparison binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
"""
def neq(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.NEQ)
    builder.left(left)
    builder.right(right)

    return builder

"""
"greater-than" comparison binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
"""
def gt(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.GTR)
    builder.left(left)
    builder.right(right)

    return builder

"""
"less-than" comparison binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
"""
def lt(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.LSS)
    builder.left(left)
    builder.right(right)

    return builder

"""
"greater-or-equal" comparison binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
"""
def gte(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.GTE)
    builder.left(left)
    builder.right(right)

    return builder

"""
"less-or-equal" comparison binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#comparison-binary-operators
"""
def lte(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.LTE)
    builder.left(left)
    builder.right(right)

    return builder

"""
"intersection" logical/set binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
"""
def and_val(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.AND)
    builder.left(left)
    builder.right(right)

    return builder

"""
"union" logical/set binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
"""
def or_val(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.OR)
    builder.left(left)
    builder.right(right)

    return builder

"""
"complement" logical/set binary operator.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#logical-set-binary-operators
"""
def unless(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.UNLESS)
    builder.left(left)
    builder.right(right)

    return builder

"""
Arc tangent binary operator. Works in radians.
Trigonometric operators allow trigonometric functions to be executed on two vectors using vector matching, which isn't available with normal functions.
They act in the same manner as arithmetic operators.
See https://prometheus.io/docs/prometheus/latest/querying/operators/#trigonometric-binary-operators
"""
def atan2(left: cogbuilder.Builder[promql.Expr], right: cogbuilder.Builder[promql.Expr]):
    builder = BinaryExpr()
    builder.op(promql.BinaryOp.ATAN2)
    builder.left(left)
    builder.right(right)

    return builder


class UnaryExpr(cogbuilder.Builder[promql.UnaryExpr]):    
    """
    Represents an unary operation expression.
    """
    
    _internal: promql.UnaryExpr

    def __init__(self):
        self._internal = promql.UnaryExpr()        
        self._internal.type_val = "unaryExpr"

    def build(self) -> promql.UnaryExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def op(self, op: promql.UnaryOp) -> typing.Self:    
        self._internal.op = op
    
        return self
    
    def expr(self, expr: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        expr_resource = expr.build()
        self._internal.expr = expr_resource
    
        return self
    

"""
Negation unary operator.
"""
def neg(expr: cogbuilder.Builder[promql.Expr]):
    builder = UnaryExpr()
    builder.op(promql.UnaryOp.MINUS)
    builder.expr(expr)

    return builder

"""
Identity unary operator.
"""
def id(expr: cogbuilder.Builder[promql.Expr]):
    builder = UnaryExpr()
    builder.op(promql.UnaryOp.PLUS)
    builder.expr(expr)

    return builder


class FuncCallExpr(cogbuilder.Builder[promql.FuncCallExpr]):    
    """
    Represents a function call expression.
    """
    
    _internal: promql.FuncCallExpr

    def __init__(self):
        self._internal = promql.FuncCallExpr()        
        self._internal.type_val = "funcCallExpr"

    def build(self) -> promql.FuncCallExpr:
        """
        Builds the object.
        """
        return self._internal    
    
    def __str__(self):
        return str(self._internal)    
    
    def function(self, function: str) -> typing.Self:    
        """
        Name of the function.
        """
            
        if not len(function) >= 1:
            raise ValueError("len(function) must be >= 1")
        self._internal.function = function
    
        return self
    
    def args(self, args: list[cogbuilder.Builder[promql.Expr]]) -> typing.Self:    
        """
        Arguments.
        """
            
        args_resources = [r1.build() for r1 in args]
        self._internal.args = args_resources
    
        return self
    
    def arg(self, arg: cogbuilder.Builder[promql.Expr]) -> typing.Self:    
        """
        Arguments.
        """
            
        if self._internal.args is None:
            self._internal.args = []
        
        arg_resource = arg.build()
        self._internal.args.append(arg_resource)
    
        return self
    

"""
Returns the input vector with all sample values converted to their absolute value.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#abs
"""
def abs(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("abs")
    builder.arg(v)

    return builder

"""
Returns an empty vector if the vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the vector passed to it has no elements.
This is useful for alerting on when no time series exist for a given metric name and label combination.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent
"""
def absent(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("absent")
    builder.arg(v)

    return builder

"""
Returns an empty vector if the range vector passed to it has any elements (floats or native histograms) and a 1-element vector with the value 1 if the range vector passed to it has no elements.
This is useful for alerting on when no time series exist for a given metric name and label combination for a certain amount of time.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#absent_over_time
"""
def absent_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("absent_over_time")
    builder.arg(v)

    return builder

"""
Rounds the sample values of all elements in `v` up to the nearest integer value greater than or equal to v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#ceil
"""
def ceil(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("ceil")
    builder.arg(v)

    return builder

"""
For each input time series, returns the number of times its value has changed within the provided time range as an instant vector.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#changes
"""
def changes(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("changes")
    builder.arg(v)

    return builder

"""
Clamps the sample values of all elements in `v` to have a lower limit of min and an upper limit of max.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp
"""
def clamp(v: cogbuilder.Builder[promql.Expr], min_val: float, max_val: float):
    builder = FuncCallExpr()
    builder.function("clamp")
    builder.arg(v)
    builder.arg(n(min_val))
    builder.arg(n(max_val))

    return builder

"""
Clamps the sample values of all elements in `v` to have an upper limit of `max`.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_max
"""
def clamp_max(v: cogbuilder.Builder[promql.Expr], max_val: float):
    builder = FuncCallExpr()
    builder.function("clamp_max")
    builder.arg(v)
    builder.arg(n(max_val))

    return builder

"""
Clamps the sample values of all elements in `v` to have an lower limit of `min`.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#clamp_min
"""
def clamp_min(v: cogbuilder.Builder[promql.Expr], min_val: float):
    builder = FuncCallExpr()
    builder.function("clamp_min")
    builder.arg(v)
    builder.arg(n(min_val))

    return builder

"""
Returns the day of the month in UTC. Returned values are from 1 to 31.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month
"""
def day_of_month():
    builder = FuncCallExpr()
    builder.function("day_of_month")

    return builder

"""
Returns the day of the month for each of the given times in UTC. Returned values are from 1 to 31.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_month
"""
def day_of_month_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("day_of_month")
    builder.arg(v)

    return builder

"""
Returns the day of the week in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week
"""
def day_of_week():
    builder = FuncCallExpr()
    builder.function("day_of_week")

    return builder

"""
Returns the day of the week for each of the given times in UTC. Returned values are from 0 to 6, where 0 means Sunday etc.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_week
"""
def day_of_week_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("day_of_week")
    builder.arg(v)

    return builder

"""
Returns the day of the year in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year
"""
def day_of_year():
    builder = FuncCallExpr()
    builder.function("day_of_year")

    return builder

"""
Returns the day of the year for each of the given times in UTC. Returned values are from 1 to 365 for non-leap years, and 1 to 366 in leap years.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#day_of_year
"""
def day_of_year_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("day_of_year")
    builder.arg(v)

    return builder

"""
Returns the number of days in the month. Returned values are from 28 to 31.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month
"""
def days_in_month():
    builder = FuncCallExpr()
    builder.function("days_in_month")

    return builder

"""
Returns the number of days in the month for each of the given times in UTC. Returned values are from 28 to 31.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#days_in_month
"""
def day_in_month_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("days_in_month")
    builder.arg(v)

    return builder

"""
Calculates the difference between the first and last value of each time series element in a range vector, returning an instant vector with the given deltas and equivalent labels.
The delta is extrapolated to cover the full time range as specified in the range vector selector, so that it is possible to get a non-integer result even if the sample values are all integers.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#delta
"""
def delta(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("delta")
    builder.arg(v)

    return builder

"""
Calculates the per-second derivative of the time series in a range vector using simple linear regression.
The range vector must have at least two samples in order to perform the calculation. When +Inf or -Inf are found in the range vector, the slope and offset value calculated will be NaN.
deriv should only be used with gauges.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#deriv
"""
def deriv(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("deriv")
    builder.arg(v)

    return builder

"""
Calculates the exponential function for all elements in vector
See https://prometheus.io/docs/prometheus/latest/querying/functions/#exp
"""
def exp(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("exp")
    builder.arg(v)

    return builder

"""
Rounds the sample values of all elements in v down to the nearest integer value smaller than or equal to v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#floor
"""
def floor(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("floor")
    builder.arg(v)

    return builder

"""
Returns the arithmetic average of observed values stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.
Note: This function only acts on native histograms.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_avg
"""
def histogram_avg(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_avg")
    builder.arg(v)

    return builder

"""
Returns the count of observations stored in a native histogram. Samples that are not native histograms are ignored and do not show up in the returned vector.
Note: This function only acts on native histograms.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum
"""
def histogram_count(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_count")
    builder.arg(v)

    return builder

"""
Returns the sum of observations stored in a native histogram.
Note: This function only acts on native histograms.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_count-and-histogram_sum
"""
def histogram_sum(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_sum")
    builder.arg(v)

    return builder

"""
Returns the estimated fraction of observations between the provided lower and upper values. Samples that are not native histograms are ignored and do not show up in the returned vector.
Note: This function only acts on native histograms.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_fraction
"""
def histogram_fraction(lower: float, upper: float, v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_fraction")
    builder.arg(n(lower))
    builder.arg(n(upper))
    builder.arg(v)

    return builder

"""
Calculates the φ-quantile (0 ≤ φ ≤ 1) from a classic histogram or from a native histogram.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_quantile
"""
def histogram_quantile(phi: float, v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_quantile")
    builder.arg(n(phi))
    builder.arg(v)

    return builder

"""
Returns the estimated standard deviation of observations in a native histogram, based on the geometric mean of the buckets where the observations lie.
Samples that are not native histograms are ignored and do not show up in the returned vector.
Note: This function only acts on native histograms.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stddev
"""
def histogram_stddev(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_stddev")
    builder.arg(v)

    return builder

"""
Returns the estimated standard variance of observations in a native histogram.
Samples that are not native histograms are ignored and do not show up in the returned vector.
Note: This function only acts on native histograms.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#histogram_stdvar
"""
def histogram_stdvar(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("histogram_stdvar")
    builder.arg(v)

    return builder

"""
Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour
"""
def hour():
    builder = FuncCallExpr()
    builder.function("hour")

    return builder

"""
Returns the hour of the day for each of the given times in UTC. Returned values are from 0 to 23.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#hour
"""
def hour_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("hour")
    builder.arg(v)

    return builder

"""
Calculates the difference between the last two samples in the range vector v, returning an instant vector with the given deltas and equivalent labels.
idelta should only be used with gauges.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#idelta
"""
def idelta(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("idelta")
    builder.arg(v)

    return builder

"""
Calculates the increase in the time series in the range vector
See https://prometheus.io/docs/prometheus/latest/querying/functions/#increase
"""
def increase(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("increase")
    builder.arg(v)

    return builder

"""
Calculates the per-second instant rate of increase of the time series in the range vector. This is based on the last two data points.
irate should only be used when graphing volatile, fast-moving counters. Use rate for alerts and slow-moving counters, as brief changes in the rate can reset the FOR clause and graphs consisting entirely of rare spikes are hard to read.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#irate
"""
def irate(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("irate")
    builder.arg(v)

    return builder

"""
matches the regular expression regex against the value of the label src_label. If it matches, the value of the label dst_label in the returned timeseries will be the expansion of replacement, together with the original labels in the input. Capturing groups in the regular expression can be referenced with $1, $2, etc. Named capturing groups in the regular expression can be referenced with $name (where name is the capturing group name). If the regular expression doesn't match then the timeseries is returned unchanged.
label_replace acts on float and histogram samples in the same way.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#label_replace
"""
def label_replace(v: cogbuilder.Builder[promql.Expr], dst_label: str, replacement: str, src_label: str, regex: str):
    builder = FuncCallExpr()
    builder.function("label_replace")
    builder.arg(v)
    builder.arg(s(dst_label))
    builder.arg(s(replacement))
    builder.arg(s(src_label))
    builder.arg(s(regex))

    return builder

"""
Calculates the natural logarithm for all elements in v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#ln
"""
def ln(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("ln")
    builder.arg(v)

    return builder

"""
Calculates the binary logarithm for all elements in v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#log2
"""
def log2(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("log2")
    builder.arg(v)

    return builder

"""
Calculates the decimal logarithm for all elements in v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#log10
"""
def log10(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("log10")
    builder.arg(v)

    return builder

"""
Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute
"""
def minute():
    builder = FuncCallExpr()
    builder.function("minute")

    return builder

"""
Returns the minute of the hour for each of the given times in UTC. Returned values are from 0 to 59.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#minute
"""
def minute_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("minute")
    builder.arg(v)

    return builder

"""
Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#month
"""
def month():
    builder = FuncCallExpr()
    builder.function("month")

    return builder

"""
Returns the month of the year for each of the given times in UTC. Returned values are from 1 to 12, where 1 means January etc.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#month
"""
def month_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("month")
    builder.arg(v)

    return builder

"""
Predicts the value of time series t seconds from now, based on the range vector v, using simple linear regression. The range vector must have at least two samples in order to perform the calculation.
predict_linear should only be used with gauges.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#predict_linear
"""
def predict_linear(v: cogbuilder.Builder[promql.Expr], t: float):
    builder = FuncCallExpr()
    builder.function("predict_linear")
    builder.arg(v)
    builder.arg(n(t))

    return builder

"""
Calculates the per-second average rate of increase of the time series in the range vector.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#rate
"""
def rate(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("rate")
    builder.arg(v)

    return builder

"""
For each input time series, resets(v range-vector) returns the number of counter resets within the provided time range as an instant vector. Any decrease in the value between two consecutive float samples is interpreted as a counter reset. A reset in a native histogram is detected in a more complex way: Any decrease in any bucket, including the zero bucket, or in the count of observation constitutes a counter reset, but also the disappearance of any previously populated bucket, an increase in bucket resolution, or a decrease of the zero-bucket width.
`resets` should only be used with counters and counter-like native histograms.
If the range vector contains a mix of float and histogram samples for the same series, counter resets are detected separately and their numbers added up. The change from a float to a histogram sample is not considered a counter reset. Each float sample is compared to the next float sample, and each histogram is compared to the next histogram.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#resets
"""
def resets(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("resets")
    builder.arg(v)

    return builder

"""
Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#round
"""
def round(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("round")
    builder.arg(v)

    return builder

"""
Rounds the sample values of all elements in v to the nearest integer. Ties are resolved by rounding up.
The to_nearest argument allows specifying the nearest multiple to which the sample values should be rounded. This multiple may also be a fraction.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#round
"""
def round_to(v: cogbuilder.Builder[promql.Expr], to_nearest: float):
    builder = FuncCallExpr()
    builder.function("round")
    builder.arg(v)
    builder.arg(n(to_nearest))

    return builder

"""
Given a single-element input vector, scalar() returns the sample value of that single element as a scalar.
If the input vector does not have exactly one element, scalar will return NaN.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#scalar
"""
def scalar(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("scalar")
    builder.arg(v)

    return builder

"""
Returns a vector with all sample values converted to their sign, defined as this: 1 if v is positive, -1 if v is negative and 0 if v is equal to zero.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#sgn
"""
def sgn(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sgn")
    builder.arg(v)

    return builder

"""
Returns vector elements sorted by their sample values, in ascending order. Native histograms are sorted by their sum of observations.
Note that sort only affects the results of instant queries, as range query results always have a fixed output ordering.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort
"""
def sort(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sort")
    builder.arg(v)

    return builder

"""
Same as `sort()`, but sorts in descending order.
Like sort, sort_desc only affects the results of instant queries, as range query results always have a fixed output ordering.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#sort_desc
"""
def sort_desc(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sort_desc")
    builder.arg(v)

    return builder

"""
Calculates the square root of all elements in v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#sqrt
"""
def sqrt(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sqrt")
    builder.arg(v)

    return builder

"""
Returns the number of seconds since January 1, 1970 UTC. Note that this does not actually return the current time, but the time at which the expression is to be evaluated.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#time
"""
def time():
    builder = FuncCallExpr()
    builder.function("time")

    return builder

"""
Returns the timestamp of each of the samples of the given vector as the number of seconds since January 1, 1970 UTC. It also works with histogram samples.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#timestamp
"""
def timestamp(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("timestamp")
    builder.arg(v)

    return builder

"""
Returns the scalar s as a vector with no labels.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#vector
"""
def vect(s: float):
    builder = FuncCallExpr()
    builder.function("vector")
    builder.arg(n(s))

    return builder

"""
Returns the year for each of the given times in UTC.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#year
"""
def year():
    builder = FuncCallExpr()
    builder.function("year")

    return builder

"""
Returns the year for each of the given times in UTC.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#year
"""
def year_for(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("year")
    builder.arg(v)

    return builder

"""
Calculates average value of all points in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def avg_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("avg_over_time")
    builder.arg(v)

    return builder

"""
Calculates the minimum value of all points in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def min_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("min_over_time")
    builder.arg(v)

    return builder

"""
Calculates the maximum value of all points in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def max_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("max_over_time")
    builder.arg(v)

    return builder

"""
Calculates the sum of all values in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def sum_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sum_over_time")
    builder.arg(v)

    return builder

"""
Calculates the count of all values in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def count_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("count_over_time")
    builder.arg(v)

    return builder

"""
Calculates the φ-quantile (0 ≤ φ ≤ 1) of the values in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def quantile_over_time(phi: float, v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("quantile_over_time")
    builder.arg(n(phi))
    builder.arg(v)

    return builder

"""
Calculates the population standard deviation of the values in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def stddev_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("stddev_over_time")
    builder.arg(v)

    return builder

"""
Calculates the population standard variance of the values in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def stdvar_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("stdvar_over_time")
    builder.arg(v)

    return builder

"""
Returns the most recent point value in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def last_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("last_over_time")
    builder.arg(v)

    return builder

"""
Returns the value 1 for any series in the specified interval.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#aggregation_over_time
"""
def present_over_time(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("present_over_time")
    builder.arg(v)

    return builder

"""
Calculates the arccosine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def acos(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("acos")
    builder.arg(v)

    return builder

"""
Calculates the inverse hyperbolic cosine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def acosh(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("acosh")
    builder.arg(v)

    return builder

"""
Calculates the arcsine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def asin(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("asin")
    builder.arg(v)

    return builder

"""
Calculates the inverse hyperbolic sine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def asinh(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("asinh")
    builder.arg(v)

    return builder

"""
Calculates the arctangent of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def atan(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("atan")
    builder.arg(v)

    return builder

"""
Calculates the inverse hyperbolic tangent of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def atanh(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("atanh")
    builder.arg(v)

    return builder

"""
Calculates the cosine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def cos(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("cos")
    builder.arg(v)

    return builder

"""
Calculates the hyperbolic cosine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def cosh(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("cosh")
    builder.arg(v)

    return builder

"""
Calculates the sine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def sin(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sin")
    builder.arg(v)

    return builder

"""
Calculates the hyperbolic sine of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def sinh(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("sinh")
    builder.arg(v)

    return builder

"""
Calculates the tangent of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def tan(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("tan")
    builder.arg(v)

    return builder

"""
Calculates the hyperbolic tangent of all elements in v
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def tanh(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("tanh")
    builder.arg(v)

    return builder

"""
Converts radians to degrees for all elements in v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def deg(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("deg")
    builder.arg(v)

    return builder

"""
Returns pi.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def pi():
    builder = FuncCallExpr()
    builder.function("pi")

    return builder

"""
Converts degrees to radians for all elements in v.
See https://prometheus.io/docs/prometheus/latest/querying/functions/#trigonometric-functions
"""
def rad(v: cogbuilder.Builder[promql.Expr]):
    builder = FuncCallExpr()
    builder.function("rad")
    builder.arg(v)

    return builder
