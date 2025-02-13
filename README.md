# PromQL builder

A set of libraries for writing and composing PromQL queries as-code in Go,
Python or Typescript.

## Go

```go
package main

import (
    "fmt"

    "github.com/grafana/promql-builder/go/promql"
)

// time() - demo_batch_last_success_timestamp_seconds > 3600
func batchJobsWithNoSuccessInLastHour() *promql.BinaryExprBuilder {
    return promql.Gt(
        promql.Sub(
            promql.Time(),
            promql.Vector("demo_batch_last_success_timestamp_seconds"),
        ),
        promql.N(3600),
    )
}

// method_code:http_errors:rate5m{code="500"} / ignoring(code) method:http_requests:rate5m
func errorRatioPerHTTPMethod() *promql.BinaryExprBuilder {
    return promql.Div(
        promql.Vector("method_code:http_errors:rate5m").Label("code", "500"),
        promql.Vector("method:http_requests:rate5m"),
    ).Ignoring([]string{"code"})
}

// sum by(device) (node_filesystem_free_bytes)
func freeDiskSpacePerDevice() *promql.AggregationExprBuilder {
    return promql.Sum(
        promql.Vector("free_disk_space_per_device"),
    ).By([]string{"device"})
}

// 90th percentile request latency over last 5 minutes per path and method
// histogram_quantile(0.9, sum by(le, path, method) (
//	rate(demo_api_request_duration_seconds_bucket[5m])
// ))
func requestLatency90thPercentilePerPathAndMethod() *promql.FuncCallExprBuilder {
    return promql.HistogramQuantile(0.9,
        promql.Sum(
            promql.Rate(
                promql.Vector("demo_api_request_duration_seconds_bucket").Range("5m"),
            ),
        ).By([]string{"le", "path", "method"}),
    )
}

func main() {
    fmt.Println(batchJobsWithNoSuccessInLastHour().String())
    fmt.Println(errorRatioPerHTTPMethod().String())
    fmt.Println(freeDiskSpacePerDevice().String())
    fmt.Println(requestLatency90thPercentilePerPathAndMethod().String())
}
```

## Python

```python
from promql_builder.builders.promql import (
    div,
    gt,
    histogram_quantile,
    n,
    rate,
    sub,
    sum,
    time,
    vector,
)

# time() - demo_batch_last_success_timestamp_seconds > 3600
def batch_jobs_with_no_success_in_last_hour():
    return gt(
        sub(
            time(),
            vector("demo_batch_last_success_timestamp_seconds"),
        ),
        n(3600),
    )


# method_code:http_errors:rate5m{code="500"} / ignoring(code) method:http_requests:rate5m
def error_ratio_per_http_method():
    return div(
        vector("method_code:http_errors:rate5m").label("code", "500"),
        vector("method:http_requests:rate5m"),
    ).ignoring(['code'])


# sum by(device) (node_filesystem_free_bytes)
def free_disk_space_per_device():
    return sum(
        vector("free_disk_space_per_device"),
    ).by(["device"])


# 90th percentile request latency over last 5 minutes per path and method
# histogram_quantile(0.9, sum by(le, path, method) (
#	rate(demo_api_request_duration_seconds_bucket[5m])
# ))
def request_latency_90th_percentile_per_path_and_method():
    return histogram_quantile(0.9,
        sum(
            rate(vector("demo_api_request_duration_seconds_bucket").range("5m")),
        ).by(["le", "path", "method"]),
    )


if __name__ == '__main__':
    print(str(batch_jobs_with_no_success_in_last_hour()))
    print(str(error_ratio_per_http_method()))
    print(str(free_disk_space_per_device()))
    print(str(request_latency_90th_percentile_per_path_and_method()))
```

## Maturity

This project should be considered as "public preview". While it is used by
Grafana Labs, it is still under active development.

Additional information can be found in
[Release life cycle for Grafana Labs](https://grafana.com/docs/release-life-cycle/).

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support
> or SLAs are not available.

## License

[Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0)
