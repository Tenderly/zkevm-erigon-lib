package zkevm_remote

import (
	"strings"

	"github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevm_types"
)

func NodeInfoReplyLess(i, j *zkevm_types.NodeInfoReply) bool {
	if cmp := strings.Compare(i.Name, j.Name); cmp != 0 {
		return cmp == -1
	}
	return strings.Compare(i.Enode, j.Enode) == -1
}
