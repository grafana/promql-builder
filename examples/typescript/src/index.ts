import * as promql from "@grafana/promql-builder";

// Free disk space per device
// sum by(device) (node_filesystem_free_bytes{machine="foo"})
function freeDiskSpacePerDevice() {
    return promql.sum(promql.vector("free_disk_space_per_device").label("machine", "foo")).by(["device"]);
}

// time() - demo_batch_last_success_timestamp_seconds
function ageOfLastSuccessfulBatchJobRun() {
    return promql.sub(
        promql.time(),
        promql.vector("demo_batch_last_success_timestamp_seconds"),
    );
}

// time() - demo_batch_last_success_timestamp_seconds > 3600
function batchJobsWithNoSuccessInLastHour() {
    return promql.gt(
        promql.sub(
            promql.time(),
            promql.vector('demo_batch_last_success_timestamp_seconds'),
        ),
        promql.n(3600)
    );
}

// process_resident_memory_bytes offset 1d
function offsetData() {
    return promql.vector("process_resident_memory_bytes").offset('1d');
}

// node_cpu_seconds_total{cpu!="0",mode=~"user|system"}
function labelMatchers() {
    return promql.vector("node_cpu_seconds_total")
        .labelNeq("cpu", "0")
        .labelMatchRegexp("mode", "user|system");
}

// node_cpu_seconds_total[5m]
function rangeVector() {
    return promql.vector("node_cpu_seconds_total").range("5m");
}

// rate(demo_api_request_duration_seconds_count[5m])
function simpleRate() {
    return promql.rate(
        promql.vector("demo_api_request_duration_seconds_count").range("5m"),
    );
}

// method_code:http_errors:rate5m{code="500"} / ignoring(code) method:http_requests:rate5m
function errorRatioPerHTTPMethod() {
    return promql.div(
        promql.vector("method_code:http_errors:rate5m").label("code", "500"),
        promql.vector("method:http_requests:rate5m"),
    ).ignoring(["code"]);
}

// absent(up{job="some-job"})
function absent() {
    return promql.absent(promql.vector("up").label("job", "some-job"));
}

// min_over_time(rate(http_requests_total[5m])[30m:1m])
function simpleSubquery() {
    const httpRequestsRate = promql.rate(
        promql.vector("http_requests_total").range("5m"),
    );

    return promql.minOverTime(
        promql.subquery(httpRequestsRate).range('30m').resolution('1m')
    );
}

// 90th percentile request latency over last 5 minutes per path and method
// histogram_quantile(0.9, sum by(le, path, method) (
//	rate(demo_api_request_duration_seconds_bucket[5m])
// ))
function requestLatency90thPercentilePerPathAndMethod() {
    return promql.histogramQuantile(
        0.9,
        promql.sum(
            promql.rate(
                promql.vector("demo_api_request_duration_seconds_bucket").range("5m"),
            ),
        ).by(["le", "path", "method"]),
    );
}
const expressions = [
    freeDiskSpacePerDevice(),
    ageOfLastSuccessfulBatchJobRun(),
    batchJobsWithNoSuccessInLastHour(),
    offsetData(),
    labelMatchers(),
    rangeVector(),
    simpleRate(),
    errorRatioPerHTTPMethod(),
    absent(),
    simpleSubquery(),
    requestLatency90thPercentilePerPathAndMethod(),
];

expressions.forEach(expression => console.log(expression.toString()));
