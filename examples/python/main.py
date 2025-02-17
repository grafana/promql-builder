import promql_builder.builders.promql as promql


# time() - demo_batch_last_success_timestamp_seconds
def age_of_last_successful_batch_job():
    return promql.sub(
        promql.time(),
        promql.vector("demo_batch_last_success_timestamp_seconds"),
    )


# time() - demo_batch_last_success_timestamp_seconds > 3600
def batch_jobs_with_no_success_in_last_hour():
    return promql.gt(
        promql.sub(
            promql.time(),
            promql.vector("demo_batch_last_success_timestamp_seconds"),
        ),
        promql.n(3600),
    )


# process_resident_memory_bytes offset 1d
def offset_data():
    return promql.vector("process_resident_memory_bytes").offset("1d")


# node_cpu_seconds_total{cpu!="0",mode=~"user|system"}
def label_matchers():
    return (promql.vector("node_cpu_seconds_total")
        .label("cpu", "0")
        .label_match_regexp("mode", "user|system"))


# node_cpu_seconds_total[5m]
def simple_range():
    return promql.vector("node_cpu_seconds_total").range("5m")


# rate(demo_api_request_duration_seconds_count[5m])
def simple_rate():
    return promql.rate(
        promql.vector("demo_api_request_duration_seconds_count").range("5m"),
    )


# method_code:http_errors:rate5m{code="500"} / ignoring(code) method:http_requests:rate5m
def error_ration_per_http_method():
    return promql.div(
        promql.vector("method_code:http_errors:rate5m").label("code", "500"),
        promql.vector("method:http_requests:rate5m").label("code", "500"),
    ).ignoring(["code"])


# absent(up{job="some-job"})
def absent_metric():
    return promql.absent(
        promql.vector("up").label("job", "some-job"),
    )


# Free disk space per device
# sum by(device) (node_filesystem_free_bytes)
def free_disk_space_per_device():
    return promql.sum(
        promql.vector("free_disk_space_per_device"),
    ).by(["device"])


# 90th percentile request latency over last 5 minutes per path and method
# histogram_quantile(0.9, sum by(le, path, method) (
#	rate(demo_api_request_duration_seconds_bucket[5m])
# ))
def request_latency_90th_percentile_per_path_and_method():
    return promql.histogram_quantile(0.9,
        promql.sum(
            promql.rate(promql.vector("demo_api_request_duration_seconds_bucket").range("5m")),
        ).by(["le", "path", "method"]),
    )


# min_over_time(rate(http_requests_total[5m])[30m:1m])
def simple_subquery():
    httpRequestsRate = promql.rate(promql.vector("http_requests_total").range("5m"))

    return promql.min_over_time(
        promql.subquery(httpRequestsRate).range("30m").resolution("1m"),
    )


if __name__ == '__main__':
   queries = [
       age_of_last_successful_batch_job(),
       batch_jobs_with_no_success_in_last_hour(),
       offset_data(),
       label_matchers(),
       simple_range(),
       simple_rate(),
       error_ration_per_http_method(),
       absent_metric(),
       free_disk_space_per_device(),
       request_latency_90th_percentile_per_path_and_method(),
       simple_subquery(),
   ]

   for query in queries:
       print(query)
