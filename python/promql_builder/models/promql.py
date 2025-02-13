# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
import enum


# Represents a PromQL expression.
Expr: typing.TypeAlias = typing.Union['NumberLiteralExpr', 'StringLiteralExpr', 'SubqueryExpr', 'AggregationExpr', 'VectorExpr', 'BinaryExpr', 'UnaryExpr', 'FuncCallExpr']


class NumberLiteralExpr:
    """
    Represents a number literal expression.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#float-literals-and-time-durations
    """

    type_val: typing.Literal["numberLiteralExpr"]
    value: float

    def __init__(self, value: float = 0):
        self.type_val = "numberLiteralExpr"
        self.value = value

    
    def __str__(self):
        return str(self.value)


class StringLiteralExpr:
    """
    Represents a string literal expression.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#string-literals
    """

    type_val: typing.Literal["stringLiteralExpr"]
    value: str

    def __init__(self, value: str = ""):
        self.type_val = "stringLiteralExpr"
        self.value = value

    
    def __str__(self):
        return self.value


class SubqueryExpr:
    """
    Represents a subquery.
    See https://prometheus.io/docs/prometheus/latest/querying/basics/#subquery
    """

    type_val: typing.Literal["subqueryExpr"]
    expr: 'Expr'
    # The offset modifier allows changing the time offset for individual instant and range vectors in a query.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
    offset: str
    # The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
    # The time supplied to the @ modifier is a unix timestamp.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
    at: str
    # Range of samples back from the current instant.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
    range_val: str
    # Empty string for default resolution.
    resolution: typing.Optional[str]

    def __init__(self, expr: typing.Optional['Expr'] = None, offset: str = "", at: str = "", range_val: str = "", resolution: typing.Optional[str] = None):
        self.type_val = "subqueryExpr"
        self.expr = expr if expr is not None else NumberLiteralExpr()
        self.offset = offset
        self.at = at
        self.range_val = range_val
        self.resolution = resolution

    
    def __str__(self):
        buffer = f'({self.expr!s})'
    
        if self.range_val != "":
            buffer += "[" + self.range_val
    
            if self.resolution is not None:
                buffer += ":"+self.resolution
    
        if self.offset != "":
            buffer += " offset " + self.offset
    
        if self.at != "":
            buffer += " @ " + self.at
    
        return buffer


class AggregationExpr:
    """
    Represents an aggregation.
    See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
    """

    type_val: typing.Literal["aggregationExpr"]
    op: 'AggregationOp'
    expr: 'Expr'
    param: typing.Optional['Expr']
    # By drops labels that are not listed in the by clause.
    by: list[str]
    # List of labels to remove from the result vector, while all other labels are preserved in the output.
    without: list[str]

    def __init__(self, op: typing.Optional['AggregationOp'] = None, expr: typing.Optional['Expr'] = None, param: typing.Optional['Expr'] = None, by: typing.Optional[list[str]] = None, without: typing.Optional[list[str]] = None):
        self.type_val = "aggregationExpr"
        self.op = op if op is not None else AggregationOp.SUM
        self.expr = expr if expr is not None else NumberLiteralExpr()
        self.param = param
        self.by = by if by is not None else []
        self.without = without if without is not None else []

    
    def __str__(self):
        buffer = self.op
    
        if len(self.without) != 0:
            buffer += ' without(' + ', '.join(self.without) + ') '
    
        if len(self.by) != 0:
            buffer += ' by(' + ', '.join(self.by) + ') '
    
        buffer += '('
    
        if self.param is not None:
            buffer += str(self.param) + ', '
    
        buffer += str(self.expr)
        buffer += ')'
    
        return buffer


class AggregationOp(enum.StrEnum):
    """
    Possible aggregation operators.
    """

    SUM = "sum"
    MIN = "min"
    MAX = "max"
    AVG = "avg"
    STDDEV = "stddev"
    STDVAR = "stdvar"
    COUNT = "count"
    GROUP = "group"
    COUNT_VALUES = "count_values"
    BOTTOMK = "bottomk"
    TOPK = "topk"
    QUANTILE = "quantile"
    LIMITK = "limitk"
    LIMIT_RATIO = "limit_ratio"


class VectorExpr:
    """
    Represents both instant and range vectors
    """

    type_val: typing.Literal["vectorExpr"]
    # Metric name.
    metric: str
    # Label selectors used to filter the timeseries.
    labels: list['LabelSelector']
    # The offset modifier allows changing the time offset for individual instant and range vectors in a query.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#offset-modifier
    offset: str
    # The `at` (or `@`) modifier allows changing the evaluation time for individual instant and range vectors in a query.
    # The time supplied to the @ modifier is a unix timestamp.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier
    at: str
    # Range of samples back from the current instant.
    # https://prometheus.io/docs/prometheus/latest/querying/basics/#range-vector-selectors
    range_val: str

    def __init__(self, metric: str = "", labels: typing.Optional[list['LabelSelector']] = None, offset: str = "", at: str = "", range_val: str = ""):
        self.type_val = "vectorExpr"
        self.metric = metric
        self.labels = labels if labels is not None else []
        self.offset = offset
        self.at = at
        self.range_val = range_val

    
    def __str__(self):
        buffer = self.metric
    
        if len(self.labels) != 0:
            buffer += "{" + ", ".join([str(label) for label in self.labels]) + "}"
    
        if self.range_val != "":
            buffer += "[" + self.range_val + "]"
    
        if self.offset != "":
            buffer += " offset " + self.offset
    
        if self.at != "":
            buffer += " @ " + self.at
    
        return buffer


class LabelSelector:
    # Name of the label to select.
    name: str
    # Value to match against.
    value: str
    # Operator used to perform the selection.
    operator: 'LabelMatchingOperator'

    def __init__(self, name: str = "", value: str = "", operator: typing.Optional['LabelMatchingOperator'] = None):
        self.name = name
        self.value = value
        self.operator = operator if operator is not None else LabelMatchingOperator.EQUAL

    
    def __str__(self):
        return f'{self.name}{self.operator}"{self.value}"'


class LabelMatchingOperator(enum.StrEnum):
    """
    Possible label matching operators.
    """

    EQUAL = "="
    NOT_EQUAL = "!="
    MATCH_REGEXP = "=~"
    NOT_MATCH_REGEXP = "!~"


class BinaryExpr:
    """
    Represents a binary operation expression.
    """

    type_val: typing.Literal["binaryExpr"]
    op: 'BinaryOp'
    left: 'Expr'
    right: 'Expr'
    # https://prometheus.io/docs/prometheus/latest/querying/operators/#vector-matching-keywords
    match_type: typing.Optional[typing.Literal["on", "ignore"]]
    match_labels: typing.Optional[list[str]]
    group_modifier: typing.Optional[typing.Literal["left", "right"]]
    group_labels: typing.Optional[list[str]]

    def __init__(self, op: typing.Optional['BinaryOp'] = None, left: typing.Optional['Expr'] = None, right: typing.Optional['Expr'] = None, match_type: typing.Optional[typing.Literal["on", "ignore"]] = None, match_labels: typing.Optional[list[str]] = None, group_modifier: typing.Optional[typing.Literal["left", "right"]] = None, group_labels: typing.Optional[list[str]] = None):
        self.type_val = "binaryExpr"
        self.op = op if op is not None else BinaryOp.ADD
        self.left = left if left is not None else NumberLiteralExpr()
        self.right = right if right is not None else NumberLiteralExpr()
        self.match_type = match_type
        self.match_labels = match_labels
        self.group_modifier = group_modifier
        self.group_labels = group_labels

    
    def __str__(self):
        buffer = ""
    
        buffer += "(" + str(self.left) + ")"
        buffer += " " + self.op + " "
    
        if self.match_type is not None:
            buffer += "on" if self.match_type == "on" else "ignoring"
    
            if self.match_labels is not None and len(self.match_labels) != 0:
                buffer += "(" + ", ".join(self.match_labels) + ") "
    
        if self.group_modifier is not None:
            buffer += "group_left" if self.group_modifier == "left" else "group_right"
    
            if self.group_labels is not None and len(self.group_labels) != 0:
                buffer += "(" + ", ".join(self.group_labels) + ") "
    
        buffer += "(" + str(self.right) + ")"
    
        return buffer


class BinaryOp(enum.StrEnum):
    """
    Possible binary operators.
    """

    ADD = "+"
    SUB = "-"
    MUL = "*"
    DIV = "/"
    MOD = "%"
    POW = "^"
    EQL = "=="
    NEQ = "!="
    GTR = ">"
    LSS = "<"
    GTE = ">="
    LTE = "<="
    AND = "and"
    OR = "or"
    UNLESS = "unless"
    ATAN2 = "atan2"


class UnaryExpr:
    """
    Represents an unary operation expression.
    """

    type_val: typing.Literal["unaryExpr"]
    op: 'UnaryOp'
    expr: 'Expr'

    def __init__(self, op: typing.Optional['UnaryOp'] = None, expr: typing.Optional['Expr'] = None):
        self.type_val = "unaryExpr"
        self.op = op if op is not None else UnaryOp.PLUS
        self.expr = expr if expr is not None else NumberLiteralExpr()

    
    def __str__(self):
        return f'{self.op}{self.expr!s}'


class UnaryOp(enum.StrEnum):
    """
    Possible unary operators.
    """

    PLUS = "+"
    MINUS = "-"


class FuncCallExpr:
    """
    Represents a function call expression.
    """

    type_val: typing.Literal["funcCallExpr"]
    # Name of the function.
    function: str
    # Arguments.
    args: list['Expr']

    def __init__(self, function: str = "", args: typing.Optional[list['Expr']] = None):
        self.type_val = "funcCallExpr"
        self.function = function
        self.args = args if args is not None else []

    
    def __str__(self):
        buffer = str(self.function) + "("
        buffer += ", ".join([str(arg) for arg in self.args])
        buffer += ")"
    
        return buffer



