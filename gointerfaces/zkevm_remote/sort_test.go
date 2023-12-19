package zkevm_remote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevm_remote"
	"github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevm_types"
	"golang.org/x/exp/slices"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		got  *zkevm_remote.NodesInfoReply
		want *zkevm_remote.NodesInfoReply
	}{
		{
			name: "sort by name",
			got: &zkevm_remote.NodesInfoReply{
				NodesInfo: []*zkevm_types.NodeInfoReply{
					{Name: "b", Enode: "c"},
					{Name: "a", Enode: "d"},
				},
			},
			want: &zkevm_remote.NodesInfoReply{
				NodesInfo: []*zkevm_types.NodeInfoReply{
					{Name: "a", Enode: "d"},
					{Name: "b", Enode: "c"},
				},
			},
		},
		{
			name: "sort by enode",
			got: &zkevm_remote.NodesInfoReply{
				NodesInfo: []*zkevm_types.NodeInfoReply{
					{Name: "a", Enode: "d"},
					{Name: "a", Enode: "c"},
				},
			},
			want: &zkevm_remote.NodesInfoReply{
				NodesInfo: []*zkevm_types.NodeInfoReply{
					{Name: "a", Enode: "c"},
					{Name: "a", Enode: "d"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			slices.SortFunc(tt.got.NodesInfo, zkevm_remote.NodeInfoReplyLess)
			assert.Equal(t, tt.want, tt.got)
		})
	}
}
