package oteltracer

import "errors"

var ErrOTLPEndpointRequired = errors.New("endpoint must be provided")
var ErrUnsupportedEndpoint = errors.New("unsupported: Endpoint must contain port 4318 (HTTP) or 4317 (gRPC)")
