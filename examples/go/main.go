package main

import (
	"fmt"

	"github.com/grafana/promql-builder/go/promql"
)

// time() - demo_batch_last_success_timestamp_seconds
func ageOfLastSuccessfulBatchJobRun() *promql.BinaryExprBuilder {
	return promql.Sub(
		promql.Time(),
		promql.Vector("demo_batch_last_success_timestamp_seconds"),
	)
}

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

// process_resident_memory_bytes offset 1d
func offsetData() *promql.VectorExprBuilder {
	return promql.Vector("process_resident_memory_bytes").Offset("1d")
}

// node_cpu_seconds_total{cpu!="0",mode=~"user|system"}
func labelMatchers() *promql.VectorExprBuilder {
	return promql.Vector("node_cpu_seconds_total").
		Label("cpu", "0").
		LabelMatchRegexp("mode", "user|system")
}

// node_cpu_seconds_total[5m]
func simpleRange() *promql.VectorExprBuilder {
	return promql.Vector("node_cpu_seconds_total").Range("5m")
}

// rate(demo_api_request_duration_seconds_count[5m])
func simpleRate() *promql.FuncCallExprBuilder {
	return promql.Rate(
		promql.Vector("demo_api_request_duration_seconds_count").
			Range("5m"),
	)
}

// method_code:http_errors:rate5m{code="500"} / ignoring(code) method:http_requests:rate5m
func errorRatioPerHTTPMethod() *promql.BinaryExprBuilder {
	return promql.Div(
		promql.Vector("method_code:http_errors:rate5m").Label("code", "500"),
		promql.Vector("method:http_requests:rate5m"),
	).Ignoring([]string{"code"})
}

// absent(up{job="some-job"})
func absent() *promql.FuncCallExprBuilder {
	return promql.Absent(
		promql.Vector("up").Label("job", "some-job"),
	)
}

// sum by(device) (node_filesystem_free_bytes)
func freeDiskSpacePerDevice() *promql.AggregationExprBuilder {
	return promql.Sum(
		promql.Vector("free_disk_space_per_device"),
	).By([]string{"device"})
}

// 90th percentile request latency over last 5 minutes per path and method
// histogram_quantile(0.9, sum by(le, path, method) (
//
//	rate(demo_api_request_duration_seconds_bucket[5m])
//
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

// min_over_time(rate(http_requests_total[5m])[30m:1m])
func simpleSubquery() *promql.FuncCallExprBuilder {
	httpRequestsRate := promql.Rate(
		promql.Vector("http_requests_total").Range("5m"),
	)

	return promql.MinOverTime(
		promql.Subquery(httpRequestsRate).
			Range("30m").
			Resolution("1m"),
	)
}

func main() {
	fmt.Println(ageOfLastSuccessfulBatchJobRun())
	fmt.Println(batchJobsWithNoSuccessInLastHour())
	fmt.Println(offsetData())
	fmt.Println(labelMatchers())
	fmt.Println(simpleRange())
	fmt.Println(simpleRate())
	fmt.Println(errorRatioPerHTTPMethod())
	fmt.Println(absent())
	fmt.Println(freeDiskSpacePerDevice())
	fmt.Println(requestLatency90thPercentilePerPathAndMethod())
	fmt.Println(simpleSubquery())
}
