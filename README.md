# protoc-gen-fastergrpc
Package implements a protobuf generator that generates a memory optimized version of the grpc service. It does this by creating memory pools (sync.Pool) for each input and output type of a unary method.

See [/examples](/examples) for more.