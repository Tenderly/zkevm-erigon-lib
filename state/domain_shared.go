package state

import (
	"bytes"
	"container/heap"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	btree2 "github.com/tidwall/btree"

	"github.com/ledgerwatch/erigon-lib/commitment"
	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/common/length"
	"github.com/ledgerwatch/erigon-lib/kv"
)

// KvList sort.Interface to sort write list by keys
type KvList struct {
	Keys []string
	Vals [][]byte
}

func (l *KvList) Push(key string, val []byte) {
	l.Keys = append(l.Keys, key)
	l.Vals = append(l.Vals, val)
}

func (l *KvList) Len() int {
	return len(l.Keys)
}

func (l *KvList) Less(i, j int) bool {
	return l.Keys[i] < l.Keys[j]
}

func (l *KvList) Swap(i, j int) {
	l.Keys[i], l.Keys[j] = l.Keys[j], l.Keys[i]
	l.Vals[i], l.Vals[j] = l.Vals[j], l.Vals[i]
}

func splitKey(key []byte) (k1, k2 []byte) {
	switch {
	case len(key) <= length.Addr:
		return key, nil
	case len(key) >= length.Addr+length.Hash:
		return key[:length.Addr], key[length.Addr:]
	default:
		panic(fmt.Sprintf("invalid key length %d", len(key)))
	}
	return
}

type SharedDomains struct {
	aggCtx *AggregatorV3Context
	roTx   kv.Tx

	txNum    atomic.Uint64
	blockNum atomic.Uint64
	estSize  atomic.Uint64

	muMaps     sync.RWMutex
	account    *btree2.Map[string, []byte]
	code       *btree2.Map[string, []byte]
	storage    *btree2.Map[string, []byte]
	commitment *btree2.Map[string, []byte]
	Account    *Domain
	Storage    *Domain
	Code       *Domain
	Commitment *DomainCommitted
}

func (sd *SharedDomains) Unwind(rwtx kv.RwTx) {
	sd.muMaps.Lock()
	defer sd.muMaps.Unlock()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	//defer cancel()
	//
	//logEvery := time.NewTicker(time.Second * 30)
	//if err := sd.flushBtree(ctx, rwtx, sd.Account.valsTable, sd.account, "sd_unwind", logEvery); err != nil {
	//	panic(err)
	//}
	//if err := sd.flushBtree(ctx, rwtx, sd.Storage.valsTable, sd.storage, "sd_unwind", logEvery); err != nil {
	//	panic(err)
	//}
	//if err := sd.flushBtree(ctx, rwtx, sd.Code.valsTable, sd.code, "sd_unwind", logEvery); err != nil {
	//	panic(err)
	//}
	//if err := sd.flushBtree(ctx, rwtx, sd.Commitment.valsTable, sd.commitment, "sd_unwind", logEvery); err != nil {
	//	panic(err)
	//}
	sd.account.Clear()
	sd.code.Clear()
	sd.commitment.Clear()
	sd.Commitment.patriciaTrie.Reset()
	sd.storage.Clear()
	sd.estSize.Store(0)
}

func (sd *SharedDomains) clear() {
	sd.muMaps.Lock()
	defer sd.muMaps.Unlock()
	sd.account.Clear()
	sd.code.Clear()
	sd.commitment.Clear()
	sd.Commitment.patriciaTrie.Reset()
	sd.storage.Clear()
	sd.estSize.Store(0)
}

func NewSharedDomains(a, c, s *Domain, comm *DomainCommitted) *SharedDomains {
	sd := &SharedDomains{
		Account:    a,
		account:    btree2.NewMap[string, []byte](128),
		Code:       c,
		code:       btree2.NewMap[string, []byte](128),
		Storage:    s,
		storage:    btree2.NewMap[string, []byte](128),
		Commitment: comm,
		commitment: btree2.NewMap[string, []byte](128),
	}

	sd.Commitment.ResetFns(sd.BranchFn, sd.AccountFn, sd.StorageFn)
	return sd
}

func (sd *SharedDomains) put(table kv.Domain, key, val []byte) {
	sd.muMaps.Lock()
	defer sd.muMaps.Unlock()
	sd.puts(table, hex.EncodeToString(key), val)
}

func (sd *SharedDomains) puts(table kv.Domain, key string, val []byte) {
	switch table {
	case kv.AccountsDomain:
		if old, ok := sd.account.Set(key, val); ok {
			sd.estSize.Add(uint64(len(val) - len(old)))
		} else {
			sd.estSize.Add(uint64(len(key) + len(val)))
		}
	case kv.CodeDomain:
		if old, ok := sd.code.Set(key, val); ok {
			sd.estSize.Add(uint64(len(val) - len(old)))
		} else {
			sd.estSize.Add(uint64(len(key) + len(val)))
		}
	case kv.StorageDomain:
		if old, ok := sd.storage.Set(key, val); ok {
			sd.estSize.Add(uint64(len(val) - len(old)))
		} else {
			sd.estSize.Add(uint64(len(key) + len(val)))
		}
	case kv.CommitmentDomain:
		if old, ok := sd.commitment.Set(key, val); ok {
			sd.estSize.Add(uint64(len(val) - len(old)))
		} else {
			sd.estSize.Add(uint64(len(key) + len(val)))
		}
	default:
		panic(fmt.Errorf("sharedDomains put to invalid table %s", table))
	}
}

func (sd *SharedDomains) Get(table kv.Domain, key []byte) (v []byte, ok bool) {
	sd.muMaps.RLock()
	v, ok = sd.get(table, key)
	sd.muMaps.RUnlock()
	return v, ok
}

func (sd *SharedDomains) get(table kv.Domain, key []byte) (v []byte, ok bool) {
	//keyS := *(*string)(unsafe.Pointer(&key))
	keyS := hex.EncodeToString(key)
	switch table {
	case kv.AccountsDomain:
		v, ok = sd.account.Get(keyS)
	case kv.CodeDomain:
		v, ok = sd.code.Get(keyS)
	case kv.StorageDomain:
		v, ok = sd.storage.Get(keyS)
	case kv.CommitmentDomain:
		v, ok = sd.commitment.Get(keyS)
	default:
		panic(table)
	}
	return v, ok
}

func (sd *SharedDomains) SizeEstimate() uint64 {
	return sd.estSize.Load()
}

func (sd *SharedDomains) LatestCommitment(prefix []byte) ([]byte, error) {
	v0, ok := sd.Get(kv.CommitmentDomain, prefix)
	if ok {
		return v0, nil
	}
	v, _, err := sd.aggCtx.CommitmentLatest(prefix, sd.roTx)
	if err != nil {
		return nil, fmt.Errorf("commitment prefix %x read error: %w", prefix, err)
	}
	return v, nil
}

func (sd *SharedDomains) LatestCode(addr []byte) ([]byte, error) {
	v0, ok := sd.Get(kv.CodeDomain, addr)
	if ok {
		return v0, nil
	}
	v, _, err := sd.aggCtx.CodeLatest(addr, sd.roTx)
	if err != nil {
		return nil, fmt.Errorf("code %x read error: %w", addr, err)
	}
	return v, nil
}

func (sd *SharedDomains) LatestAccount(addr []byte) ([]byte, error) {
	v0, ok := sd.Get(kv.AccountsDomain, addr)
	if ok {
		return v0, nil
	}
	v, _, err := sd.aggCtx.AccountLatest(addr, sd.roTx)
	if err != nil {
		return nil, fmt.Errorf("account %x read error: %w", addr, err)
	}
	return v, nil
}

func (sd *SharedDomains) ReadsValidBtree(table kv.Domain, list *KvList) bool {
	sd.muMaps.RLock()
	defer sd.muMaps.RUnlock()

	var m *btree2.Map[string, []byte]
	switch table {
	case kv.AccountsDomain:
		m = sd.account
	case kv.CodeDomain:
		m = sd.code
	case kv.StorageDomain:
		m = sd.storage
	default:
		panic(table)
	}

	for i, key := range list.Keys {
		if val, ok := m.Get(key); ok {
			if !bytes.Equal(list.Vals[i], val) {
				return false
			}
		}
	}
	return true
}

func (sd *SharedDomains) LatestStorage(addr, loc []byte) ([]byte, error) {
	v0, ok := sd.Get(kv.StorageDomain, common.Append(addr, loc))
	if ok {
		return v0, nil
	}
	v, _, err := sd.aggCtx.StorageLatest(addr, loc, sd.roTx)
	if err != nil {
		return nil, fmt.Errorf("storage %x|%x read error: %w", addr, loc, err)
	}
	return v, nil
}

func (sd *SharedDomains) BranchFn(pref []byte) ([]byte, error) {
	v, err := sd.LatestCommitment(pref)
	if err != nil {
		return nil, fmt.Errorf("branchFn failed: %w", err)
	}
	//fmt.Printf("branchFn[sd]: %x: %x\n", pref, v)
	if len(v) == 0 {
		return nil, nil
	}
	// skip touchmap
	return v[2:], nil
}

func (sd *SharedDomains) AccountFn(plainKey []byte, cell *commitment.Cell) error {
	encAccount, err := sd.LatestAccount(plainKey)
	if err != nil {
		return fmt.Errorf("accountFn failed: %w", err)
	}
	cell.Nonce = 0
	cell.Balance.Clear()
	if len(encAccount) > 0 {
		nonce, balance, chash := DecodeAccountBytes(encAccount)
		cell.Nonce = nonce
		cell.Balance.Set(balance)
		if len(chash) > 0 {
			copy(cell.CodeHash[:], chash)
		}
		//fmt.Printf("accountFn[sd]: %x: n=%d b=%d ch=%x\n", plainKey, nonce, balance, chash)
	}

	code, err := sd.LatestCode(plainKey)
	if err != nil {
		return fmt.Errorf("accountFn[sd]: failed to read latest code: %w", err)
	}
	if len(code) > 0 {
		//fmt.Printf("accountFn[sd]: code %x - %x\n", plainKey, code)
		sd.Commitment.updates.keccak.Reset()
		sd.Commitment.updates.keccak.Write(code)
		copy(cell.CodeHash[:], sd.Commitment.updates.keccak.Sum(nil))
	} else {
		copy(cell.CodeHash[:], commitment.EmptyCodeHash)
	}
	cell.Delete = len(encAccount) == 0 && len(code) == 0
	return nil
}

func (sd *SharedDomains) StorageFn(plainKey []byte, cell *commitment.Cell) error {
	// Look in the summary table first
	addr, loc := splitKey(plainKey)
	enc, err := sd.LatestStorage(addr, loc)
	if err != nil {
		return err
	}
	//fmt.Printf("storageFn[sd]: %x|%x - %x\n", addr, loc, enc)
	cell.StorageLen = len(enc)
	copy(cell.Storage[:], enc)
	cell.Delete = cell.StorageLen == 0
	return nil
}

func (sd *SharedDomains) UpdateAccountData(addr []byte, account, prevAccount []byte) error {
	sd.Commitment.TouchPlainKey(addr, account, sd.Commitment.TouchAccount)
	sd.put(kv.AccountsDomain, addr, account)
	return sd.Account.PutWithPrev(addr, nil, account, prevAccount)
}

func (sd *SharedDomains) UpdateAccountCode(addr []byte, code, codeHash []byte) error {
	sd.Commitment.TouchPlainKey(addr, code, sd.Commitment.TouchCode)
	prevCode, _ := sd.LatestCode(addr)
	if bytes.Equal(prevCode, code) {
		return nil
	}
	sd.put(kv.CodeDomain, addr, code)
	if len(code) == 0 {
		return sd.Code.DeleteWithPrev(addr, nil, prevCode)
	}
	return sd.Code.PutWithPrev(addr, nil, code, prevCode)
}

func (sd *SharedDomains) UpdateCommitmentData(prefix []byte, data, prev []byte) error {
	sd.put(kv.CommitmentDomain, prefix, data)
	return sd.Commitment.PutWithPrev(prefix, nil, data, prev)
}

func (sd *SharedDomains) DeleteAccount(addr, prev []byte) error {
	sd.Commitment.TouchPlainKey(addr, nil, sd.Commitment.TouchAccount)

	sd.put(kv.AccountsDomain, addr, nil)
	if err := sd.Account.DeleteWithPrev(addr, nil, prev); err != nil {
		return err
	}

	sd.put(kv.CodeDomain, addr, nil)
	// commitment delete already has been applied via account
	if err := sd.Code.Delete(addr, nil); err != nil {
		return err
	}

	var err error
	type pair struct{ k, v []byte }
	tombs := make([]pair, 0, 8)
	err = sd.IterateStoragePrefix(sd.roTx, addr, func(k, v []byte) {
		if !bytes.HasPrefix(k, addr) {
			return
		}
		sd.put(kv.StorageDomain, k, nil)
		sd.Commitment.TouchPlainKey(k, nil, sd.Commitment.TouchStorage)
		err = sd.Storage.DeleteWithPrev(k, nil, v)

		tombs = append(tombs, pair{k, v})
	})

	for _, tomb := range tombs {
		sd.Commitment.TouchPlainKey(tomb.k, nil, sd.Commitment.TouchStorage)
		err = sd.Storage.DeleteWithPrev(tomb.k, nil, tomb.v)
	}
	return err
}

func (sd *SharedDomains) WriteAccountStorage(addr, loc []byte, value, preVal []byte) error {
	composite := common.Append(addr, loc)

	sd.Commitment.TouchPlainKey(composite, value, sd.Commitment.TouchStorage)
	sd.put(kv.StorageDomain, composite, value)
	if len(value) == 0 {
		return sd.Storage.DeleteWithPrev(addr, loc, preVal)
	}
	return sd.Storage.PutWithPrev(addr, loc, value, preVal)
}

func (sd *SharedDomains) SetContext(ctx *AggregatorV3Context) {
	sd.aggCtx = ctx
}

func (sd *SharedDomains) SetTx(tx kv.RwTx) {
	sd.roTx = tx
	sd.Commitment.SetTx(tx)
	sd.Code.SetTx(tx)
	sd.Account.SetTx(tx)
	sd.Storage.SetTx(tx)
}

func (sd *SharedDomains) SetTxNum(txNum uint64) {
	sd.txNum.Store(txNum)
	sd.Account.SetTxNum(txNum)
	sd.Code.SetTxNum(txNum)
	sd.Storage.SetTxNum(txNum)
	sd.Commitment.SetTxNum(txNum)
}

func (sd *SharedDomains) SetBlockNum(blockNum uint64) {
	sd.blockNum.Store(blockNum)
}

func (sd *SharedDomains) Commit(saveStateAfter, trace bool) (rootHash []byte, err error) {
	// if commitment mode is Disabled, there will be nothing to compute on.
	rootHash, branchNodeUpdates, err := sd.Commitment.ComputeCommitment(trace)
	if err != nil {
		return nil, err
	}

	defer func(t time.Time) { mxCommitmentWriteTook.UpdateDuration(t) }(time.Now())

	for pref, update := range branchNodeUpdates {
		prefix := []byte(pref)

		stateValue, err := sd.LatestCommitment(prefix)
		if err != nil {
			return nil, err
		}
		stated := commitment.BranchData(stateValue)
		merged, err := sd.Commitment.branchMerger.Merge(stated, update)
		if err != nil {
			return nil, err
		}
		if bytes.Equal(stated, merged) {
			continue
		}
		if trace {
			fmt.Printf("sd computeCommitment merge [%x] [%x]+[%x]=>[%x]\n", prefix, stated, update, merged)
		}

		if err = sd.UpdateCommitmentData(prefix, merged, stated); err != nil {
			return nil, err
		}
		mxCommitmentUpdatesApplied.Inc()
	}

	if saveStateAfter {
		if err := sd.Commitment.storeCommitmentState(sd.blockNum.Load(), rootHash); err != nil {
			return nil, err
		}
	}
	return rootHash, nil
}

// IterateStoragePrefix iterates over key-value pairs of the storage domain that start with given prefix
// Such iteration is not intended to be used in public API, therefore it uses read-write transaction
// inside the domain. Another version of this for public API use needs to be created, that uses
// roTx instead and supports ending the iterations before it reaches the end.
func (sd *SharedDomains) IterateStoragePrefix(roTx kv.Tx, prefix []byte, it func(k, v []byte)) error {
	sd.Storage.stats.FilesQueries.Add(1)

	var cp CursorHeap
	heap.Init(&cp)
	var k, v []byte
	var err error

	iter := sd.storage.Iter()
	cnt := 0
	if iter.Seek(string(prefix)) {
		kx := iter.Key()
		v = iter.Value()
		cnt++
		//fmt.Printf("c %d kx: %s, k: %x\n", cnt, kx, v)
		k, _ = hex.DecodeString(kx)

		if len(kx) > 0 && bytes.HasPrefix(k, prefix) {
			heap.Push(&cp, &CursorItem{t: RAM_CURSOR, key: common.Copy(k), val: common.Copy(v), iter: iter, endTxNum: sd.txNum.Load(), reverse: true})
		}
	}

	keysCursor, err := roTx.CursorDupSort(sd.Storage.keysTable)
	if err != nil {
		return err
	}
	defer keysCursor.Close()
	if k, v, err = keysCursor.Seek(prefix); err != nil {
		return err
	}
	if k != nil && bytes.HasPrefix(k, prefix) {
		keySuffix := make([]byte, len(k)+8)
		copy(keySuffix, k)
		copy(keySuffix[len(k):], v)
		step := ^binary.BigEndian.Uint64(v)
		txNum := step * sd.Storage.aggregationStep
		if v, err = roTx.GetOne(sd.Storage.valsTable, keySuffix); err != nil {
			return err
		}
		heap.Push(&cp, &CursorItem{t: DB_CURSOR, key: common.Copy(k), val: common.Copy(v), c: keysCursor, endTxNum: txNum, reverse: true})
	}

	sctx := sd.aggCtx.storage
	for i, item := range sctx.files {
		bg := sctx.statelessBtree(i)
		if bg.Empty() {
			continue
		}

		cursor, err := bg.Seek(prefix)
		if err != nil {
			continue
		}

		g := sctx.statelessGetter(i)
		key := cursor.Key()
		if key != nil && bytes.HasPrefix(key, prefix) {
			val := cursor.Value()
			heap.Push(&cp, &CursorItem{t: FILE_CURSOR, key: key, val: val, dg: g, btCursor: cursor, endTxNum: item.endTxNum, reverse: true})
		}
	}

	for cp.Len() > 0 {
		lastKey := common.Copy(cp[0].key)
		lastVal := common.Copy(cp[0].val)
		// Advance all the items that have this key (including the top)
		for cp.Len() > 0 && bytes.Equal(cp[0].key, lastKey) {
			ci1 := cp[0]
			switch ci1.t {
			case RAM_CURSOR:
				if ci1.iter.Next() {
					k, _ = hex.DecodeString(ci1.iter.Key())
					if k != nil && bytes.HasPrefix(k, prefix) {
						ci1.key = common.Copy(k)
						ci1.val = common.Copy(ci1.iter.Value())
					}
					heap.Fix(&cp, 0)
				} else {
					heap.Pop(&cp)
				}
			case FILE_CURSOR:
				if ci1.btCursor.Next() {
					ci1.key = ci1.btCursor.Key()
					if ci1.key != nil && bytes.HasPrefix(ci1.key, prefix) {
						ci1.val = ci1.btCursor.Value()
						heap.Fix(&cp, 0)
					} else {
						heap.Pop(&cp)
					}
				} else {
					heap.Pop(&cp)
				}

				//if ci1.dg.HasNext() {
				//	ci1.key, _ = ci1.dg.Next(ci1.key[:0])
				//	if ci1.key != nil && bytes.HasPrefix(ci1.key, prefix) {
				//		ci1.val, _ = ci1.dg.Next(ci1.val[:0])
				//		heap.Fix(&cp, 0)
				//	} else {
				//		heap.Pop(&cp)
				//	}
				//} else {
				//	heap.Pop(&cp)
				//}
			case DB_CURSOR:
				k, v, err = ci1.c.NextNoDup()
				if err != nil {
					return err
				}
				if k != nil && bytes.HasPrefix(k, prefix) {
					ci1.key = common.Copy(k)
					keySuffix := make([]byte, len(k)+8)
					copy(keySuffix, k)
					copy(keySuffix[len(k):], v)
					if v, err = roTx.GetOne(sd.Storage.valsTable, keySuffix); err != nil {
						return err
					}
					ci1.val = common.Copy(v)
					heap.Fix(&cp, 0)
				} else {
					heap.Pop(&cp)
				}
			}
		}
		if len(lastVal) > 0 {
			it(lastKey, lastVal)
		}
	}
	return nil
}

func (sd *SharedDomains) Clean() {
	sd.account.Clear()
	sd.code.Clear()
	sd.storage.Clear()
	sd.commitment.Clear()
	sd.estSize.Store(0)
}

func (sd *SharedDomains) Close() {
	sd.aggCtx.Close()
	sd.account.Clear()
	sd.code.Clear()
	sd.storage.Clear()
	sd.commitment.Clear()
	sd.Account.Close()
	sd.Storage.Close()
	sd.Code.Close()
	sd.Commitment.Close()
}
