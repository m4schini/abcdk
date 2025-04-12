package consts

const (
	DefaultTelemetryEndpoint = DefaultOtelEndpointGrpc
	DefaultOtelEndpointGrpc  = "localhost:4317"
)

const (
	SchemeDocstore = "mongodb"
	SchemeGraph    = "neo4j"
	SchemeKv       = "valkey"
)
