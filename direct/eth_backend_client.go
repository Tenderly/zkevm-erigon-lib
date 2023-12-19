/*
   Copyright 2021 Erigon contributors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package direct

import (
	"context"
	"io"

	"github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevm_remote"
	"github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevm_types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EthBackendClientDirect struct {
	server zkevm_remote.ETHBACKENDServer
}

func NewEthBackendClientDirect(server zkevm_remote.ETHBACKENDServer) *EthBackendClientDirect {
	return &EthBackendClientDirect{server: server}
}

func (s *EthBackendClientDirect) EngineGetBlobsBundleV1(ctx context.Context, in *zkevm_remote.EngineGetBlobsBundleRequest, opts ...grpc.CallOption) (*zkevm_types.BlobsBundleV1, error) {
	return s.server.EngineGetBlobsBundleV1(ctx, in)
}

func (s *EthBackendClientDirect) Etherbase(ctx context.Context, in *zkevm_remote.EtherbaseRequest, opts ...grpc.CallOption) (*zkevm_remote.EtherbaseReply, error) {
	return s.server.Etherbase(ctx, in)
}

func (s *EthBackendClientDirect) NetVersion(ctx context.Context, in *zkevm_remote.NetVersionRequest, opts ...grpc.CallOption) (*zkevm_remote.NetVersionReply, error) {
	return s.server.NetVersion(ctx, in)
}

func (s *EthBackendClientDirect) NetPeerCount(ctx context.Context, in *zkevm_remote.NetPeerCountRequest, opts ...grpc.CallOption) (*zkevm_remote.NetPeerCountReply, error) {
	return s.server.NetPeerCount(ctx, in)
}

func (s *EthBackendClientDirect) EngineNewPayload(ctx context.Context, in *zkevm_types.ExecutionPayload, opts ...grpc.CallOption) (*zkevm_remote.EnginePayloadStatus, error) {
	return s.server.EngineNewPayload(ctx, in)
}

func (s *EthBackendClientDirect) EngineForkChoiceUpdated(ctx context.Context, in *zkevm_remote.EngineForkChoiceUpdatedRequest, opts ...grpc.CallOption) (*zkevm_remote.EngineForkChoiceUpdatedResponse, error) {
	return s.server.EngineForkChoiceUpdated(ctx, in)
}

func (s *EthBackendClientDirect) EngineGetPayload(ctx context.Context, in *zkevm_remote.EngineGetPayloadRequest, opts ...grpc.CallOption) (*zkevm_remote.EngineGetPayloadResponse, error) {
	return s.server.EngineGetPayload(ctx, in)
}

func (s *EthBackendClientDirect) EngineGetPayloadBodiesByHashV1(ctx context.Context, in *zkevm_remote.EngineGetPayloadBodiesByHashV1Request, opts ...grpc.CallOption) (*zkevm_remote.EngineGetPayloadBodiesV1Response, error) {
	return s.server.EngineGetPayloadBodiesByHashV1(ctx, in)
}

func (s *EthBackendClientDirect) EngineGetPayloadBodiesByRangeV1(ctx context.Context, in *zkevm_remote.EngineGetPayloadBodiesByRangeV1Request, opts ...grpc.CallOption) (*zkevm_remote.EngineGetPayloadBodiesV1Response, error) {
	return s.server.EngineGetPayloadBodiesByRangeV1(ctx, in)
}

func (s *EthBackendClientDirect) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*zkevm_types.VersionReply, error) {
	return s.server.Version(ctx, in)
}

func (s *EthBackendClientDirect) ProtocolVersion(ctx context.Context, in *zkevm_remote.ProtocolVersionRequest, opts ...grpc.CallOption) (*zkevm_remote.ProtocolVersionReply, error) {
	return s.server.ProtocolVersion(ctx, in)
}

func (s *EthBackendClientDirect) ClientVersion(ctx context.Context, in *zkevm_remote.ClientVersionRequest, opts ...grpc.CallOption) (*zkevm_remote.ClientVersionReply, error) {
	return s.server.ClientVersion(ctx, in)
}

// -- start Subscribe

func (s *EthBackendClientDirect) Subscribe(ctx context.Context, in *zkevm_remote.SubscribeRequest, opts ...grpc.CallOption) (zkevm_remote.ETHBACKEND_SubscribeClient, error) {
	ch := make(chan *subscribeReply, 16384)
	streamServer := &SubscribeStreamS{ch: ch, ctx: ctx}
	go func() {
		defer close(ch)
		streamServer.Err(s.server.Subscribe(in, streamServer))
	}()
	return &SubscribeStreamC{ch: ch, ctx: ctx}, nil
}

type subscribeReply struct {
	r   *zkevm_remote.SubscribeReply
	err error
}
type SubscribeStreamS struct {
	ch  chan *subscribeReply
	ctx context.Context
	grpc.ServerStream
}

func (s *SubscribeStreamS) Send(m *zkevm_remote.SubscribeReply) error {
	s.ch <- &subscribeReply{r: m}
	return nil
}
func (s *SubscribeStreamS) Context() context.Context { return s.ctx }
func (s *SubscribeStreamS) Err(err error) {
	if err == nil {
		return
	}
	s.ch <- &subscribeReply{err: err}
}

type SubscribeStreamC struct {
	ch  chan *subscribeReply
	ctx context.Context
	grpc.ClientStream
}

func (c *SubscribeStreamC) Recv() (*zkevm_remote.SubscribeReply, error) {
	m, ok := <-c.ch
	if !ok || m == nil {
		return nil, io.EOF
	}
	return m.r, m.err
}
func (c *SubscribeStreamC) Context() context.Context { return c.ctx }

// -- end Subscribe

// -- SubscribeLogs

func (s *EthBackendClientDirect) SubscribeLogs(ctx context.Context, opts ...grpc.CallOption) (zkevm_remote.ETHBACKEND_SubscribeLogsClient, error) {
	subscribeLogsRequestChan := make(chan *subscribeLogsRequest, 16384)
	subscribeLogsReplyChan := make(chan *subscribeLogsReply, 16384)
	srv := &SubscribeLogsStreamS{
		chSend: subscribeLogsReplyChan,
		chRecv: subscribeLogsRequestChan,
		ctx:    ctx,
	}
	go func() {
		defer close(subscribeLogsRequestChan)
		defer close(subscribeLogsReplyChan)
		srv.Err(s.server.SubscribeLogs(srv))
	}()
	cli := &SubscribeLogsStreamC{
		chSend: subscribeLogsRequestChan,
		chRecv: subscribeLogsReplyChan,
		ctx:    ctx,
	}
	return cli, nil
}

type SubscribeLogsStreamS struct {
	chSend chan *subscribeLogsReply
	chRecv chan *subscribeLogsRequest
	ctx    context.Context
	grpc.ServerStream
}

type subscribeLogsReply struct {
	r   *zkevm_remote.SubscribeLogsReply
	err error
}

type subscribeLogsRequest struct {
	r   *zkevm_remote.LogsFilterRequest
	err error
}

func (s *SubscribeLogsStreamS) Send(m *zkevm_remote.SubscribeLogsReply) error {
	s.chSend <- &subscribeLogsReply{r: m}
	return nil
}

func (s *SubscribeLogsStreamS) Recv() (*zkevm_remote.LogsFilterRequest, error) {
	m, ok := <-s.chRecv
	if !ok || m == nil {
		return nil, io.EOF
	}
	return m.r, m.err
}

func (s *SubscribeLogsStreamS) Err(err error) {
	if err == nil {
		return
	}
	s.chSend <- &subscribeLogsReply{err: err}
}

type SubscribeLogsStreamC struct {
	chSend chan *subscribeLogsRequest
	chRecv chan *subscribeLogsReply
	ctx    context.Context
	grpc.ClientStream
}

func (c *SubscribeLogsStreamC) Send(m *zkevm_remote.LogsFilterRequest) error {
	c.chSend <- &subscribeLogsRequest{r: m}
	return nil
}

func (c *SubscribeLogsStreamC) Recv() (*zkevm_remote.SubscribeLogsReply, error) {
	m, ok := <-c.chRecv
	if !ok || m == nil {
		return nil, io.EOF
	}
	return m.r, m.err
}

// -- end SubscribeLogs

func (s *EthBackendClientDirect) Block(ctx context.Context, in *zkevm_remote.BlockRequest, opts ...grpc.CallOption) (*zkevm_remote.BlockReply, error) {
	return s.server.Block(ctx, in)
}

func (s *EthBackendClientDirect) TxnLookup(ctx context.Context, in *zkevm_remote.TxnLookupRequest, opts ...grpc.CallOption) (*zkevm_remote.TxnLookupReply, error) {
	return s.server.TxnLookup(ctx, in)
}

func (s *EthBackendClientDirect) NodeInfo(ctx context.Context, in *zkevm_remote.NodesInfoRequest, opts ...grpc.CallOption) (*zkevm_remote.NodesInfoReply, error) {
	return s.server.NodeInfo(ctx, in)
}

func (s *EthBackendClientDirect) Peers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*zkevm_remote.PeersReply, error) {
	return s.server.Peers(ctx, in)
}

func (s *EthBackendClientDirect) PendingBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*zkevm_remote.PendingBlockReply, error) {
	return s.server.PendingBlock(ctx, in)
}
