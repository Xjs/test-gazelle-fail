// +build tools

// the comment above deactivates this file unless a buildflag "tools" is used.
// this is best practice for tool support with gomodules according to
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/envoyproxy/protoc-gen-validate"
//	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
//	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
//	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
