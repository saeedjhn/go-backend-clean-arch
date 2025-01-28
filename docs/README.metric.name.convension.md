# Metric Naming Conventions

When writing metric names, there are commonly followed rules that make them easier to understand and implement in
different systems. Some of these rules are:

### 1. Use of `_total` suffix for Counters

Counters typically have names with a `_total` suffix to indicate that they keep track of the total count.

Example: `http_requests_total`, `grpc_errors_total`

### 2. Use of `_seconds` or `_milliseconds` suffix for time

For measuring time, it is common to use suffixes like `_seconds` (for seconds) or `_milliseconds` (for milliseconds).

Example: `http_request_duration_seconds`, `grpc_latency_seconds`

### 3. Use of descriptive and simple names

Metric names should be descriptive and simple, so their meaning can be easily understood.

Example: `cpu_usage`, `memory_usage_bytes`, `request_size_bytes`

### 4. Use of underscores to separate words

In naming metrics, underscores are commonly used to separate words for better readability.

Example: `disk_usage_percent`, `active_connections`

### 5. Use of related namespace for the metric names

To avoid name clashes, especially when metrics are used in different projects or systems, a namespace is used.
Typically, the name of the program or system starts the metric name.

Example: `myapp_http_requests_total`, `service_name_grpc_latency_seconds`

### 6. Indicating the type of metric in the name

The type of metric (e.g., Counter, Gauge, Histogram) is usually indicated in the name, or it is inferred by the
monitoring system by default.

Example: `http_requests_total` (for counters), `http_request_duration_seconds` (for time-based metrics)

### 7. Use of labels (attributes)

For metrics that need additional characteristics, labels are used. These labels are appended to the metric name to
provide more context about each instance of the metric.

Example: `http_requests_total{status="200", method="GET", endpoint="/api/v1/resource"}`

### 8. Use of meaningful and readable names

Names should be chosen in a way that is easy to understand and meaningful. Avoid using technical jargon or complicated
abbreviations whenever possible.

Example: Instead of `http_2xx_count`, use `http_requests_successful_total`

---

### Example Metric Names:

- `http_requests_total` – Counter for the total number of HTTP requests.
- `http_request_duration_seconds` – Total time taken for HTTP requests (in seconds).
- `http_response_size_bytes` – Size of HTTP responses in bytes.
- `disk_usage_percent` – Percentage of disk usage.
- `grpc_latency_seconds` – Latency for gRPC requests (in seconds).
- `grpc_errors_total` – Counter for the total number of gRPC errors.
- `memory_usage_bytes` – Memory usage in bytes.

---

### Summary:

When naming metrics, it is important to focus on readability, descriptiveness, and adherence to general monitoring
standards (e.g., Prometheus). Using prefixes, standard suffixes, and separating words with underscores helps make
metrics easily identifiable and interpretable.