package zkevmremote

import (
	"strings"

	types "github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevmtypes"
)

func NodeInfoReplyLess(i, j *types.NodeInfoReply) bool {
	if cmp := strings.Compare(i.Name, j.Name); cmp != 0 {
		return cmp == -1
	}
	return strings.Compare(i.Enode, j.Enode) == -1
}
