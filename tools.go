//go:build tools

package tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
//
// This module is just a hack for 'go mod tidy' command
//
// Problem is - 'go mod tidy' removes from go.mod file lines if you don't use them in source code
//
// But code generators we use as binaries - not by including them into source code
// To provide reproducible builds - go.mod file must be source of truth about code generators version
// 'go mod tidy' will not remove binary deps from go.mod file if you add them here
//
// use `make devtools` - does install all binary deps of right version

// build tag 'trick_go_mod_tidy' - is used to hide warnings of IDEA (because we can't import `main` packages in go)

import (
	_ "github.com/matryer/moq"
	_ "github.com/tenderly/erigon-interfaces"
	_ "github.com/tenderly/erigon-interfaces/zkevmdownloader"
	_ "github.com/tenderly/erigon-interfaces/zkevmexecution"
	_ "github.com/tenderly/erigon-interfaces/zkevmp2psentinel"
	_ "github.com/tenderly/erigon-interfaces/zkevmp2psentry"
	_ "github.com/tenderly/erigon-interfaces/zkevmremote"
	_ "github.com/tenderly/erigon-interfaces/zkevmtxpool"
	_ "github.com/tenderly/erigon-interfaces/zkevmtypes"
	_ "github.com/tenderly/erigon-interfaces/zkevmweb3"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
