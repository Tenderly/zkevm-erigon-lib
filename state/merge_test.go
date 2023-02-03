package state

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	btree2 "github.com/tidwall/btree"

	"github.com/ledgerwatch/erigon-lib/recsplit/eliasfano32"
)

func TestFindMergeRangeCornerCases(t *testing.T) {
	t.Run("not equal amount of files", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		ii.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		h := &History{InvertedIndex: ii, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		h.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})

		r := h.findMergeRange(4, 32)
		assert.True(t, r.history)
		assert.Equal(t, r.historyEndTxNum, uint64(2))
		assert.Equal(t, r.indexEndTxNum, uint64(2))
	})
	t.Run("idx merged, history not yet", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		ii.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		h := &History{InvertedIndex: ii, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		h.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})

		r := h.findMergeRange(4, 32)
		assert.True(t, r.history)
		assert.False(t, r.index)
		assert.Equal(t, uint64(2), r.historyEndTxNum)
	})
	t.Run("idx merged, history not yet, 2", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 4})

		h := &History{InvertedIndex: ii, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		h.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		h.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		h.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		r := h.findMergeRange(4, 32)
		assert.False(t, r.index)
		assert.True(t, r.history)
		assert.Equal(t, 2, int(r.historyEndTxNum))
	})
	t.Run("history merged, but index not and history garbage left", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})

		// `kill -9` may leave small garbage files, but if big one already exists we assume it's good(fsynced) and no reason to merge again
		h := &History{InvertedIndex: ii, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		h.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})

		r := h.findMergeRange(4, 32)
		assert.True(t, r.index)
		assert.False(t, r.history)
		assert.Equal(t, uint64(2), r.indexEndTxNum)
	})
	t.Run("history merge progress ahead of idx", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		ii.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		// `kill -9` may leave small garbage files, but if big one already exists we assume it's good(fsynced) and no reason to merge again
		h := &History{InvertedIndex: ii, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		h.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})
		h.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		h.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		r := h.findMergeRange(4, 32)
		assert.True(t, r.index)
		assert.False(t, r.history)
		assert.Equal(t, 2, int(r.indexEndTxNum))
	})
	t.Run("idx merge progress ahead of history", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		ii.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		// `kill -9` may leave small garbage files, but if big one already exists we assume it's good(fsynced) and no reason to merge again
		h := &History{InvertedIndex: ii, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		h.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		h.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		h.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		h.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})

		r := h.findMergeRange(4, 32)
		assert.False(t, r.index)
		assert.True(t, r.history)
		assert.Equal(t, 2, int(r.historyEndTxNum))
	})
	t.Run("idx merged, but garbage left", func(t *testing.T) {
		ii := &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})
		needMerge, _, _ := ii.findMergeRange(4, 32)
		assert.False(t, needMerge)

		ii = &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		needMerge, _, _ = ii.findMergeRange(4, 32)
		assert.False(t, needMerge)

		ii = &InvertedIndex{aggregationStep: 1, files: btree2.NewBTreeG[*filesItem](filesItemLess)}
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 1})
		ii.files.Set(&filesItem{startTxNum: 1, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 0, endTxNum: 2})
		ii.files.Set(&filesItem{startTxNum: 2, endTxNum: 3})
		ii.files.Set(&filesItem{startTxNum: 3, endTxNum: 4})
		needMerge, from, to := ii.findMergeRange(4, 32)
		assert.True(t, needMerge)
		require.Equal(t, 0, int(from))
		require.Equal(t, 4, int(to))
	})

}

func Test_mergeEliasFano(t *testing.T) {
	t.Skip()

	firstList := []int{1, 298164, 298163, 13, 298160, 298159}
	sort.Ints(firstList)
	uniq := make(map[int]struct{})

	first := eliasfano32.NewEliasFano(uint64(len(firstList)), uint64(firstList[len(firstList)-1]))
	for _, v := range firstList {
		uniq[v] = struct{}{}
		first.AddOffset(uint64(v))
	}
	first.Build()
	firstBytes := first.AppendBytes(nil)

	fit := first.Iterator()
	for fit.HasNext() {
		v, _ := fit.Next()
		require.Contains(t, firstList, int(v))
	}

	secondList := []int{
		1, 644951, 644995, 682653, 13,
		644988, 644987, 644946, 644994,
		644942, 644945, 644941, 644940,
		644939, 644938, 644792, 644787}
	sort.Ints(secondList)
	second := eliasfano32.NewEliasFano(uint64(len(secondList)), uint64(secondList[len(secondList)-1]))

	for _, v := range secondList {
		second.AddOffset(uint64(v))
		uniq[v] = struct{}{}
	}
	second.Build()
	secondBytes := second.AppendBytes(nil)

	sit := second.Iterator()
	for sit.HasNext() {
		v, _ := sit.Next()
		require.Contains(t, secondList, int(v))
	}

	menc, err := mergeEfs(firstBytes, secondBytes, nil)
	require.NoError(t, err)

	merged, _ := eliasfano32.ReadEliasFano(menc)
	require.NoError(t, err)
	require.EqualValues(t, len(uniq), merged.Count())
	require.EqualValues(t, merged.Count(), eliasfano32.Count(menc))
	mergedLists := append(firstList, secondList...)
	sort.Ints(mergedLists)
	require.EqualValues(t, mergedLists[len(mergedLists)-1], merged.Max())
	require.EqualValues(t, merged.Max(), eliasfano32.Max(menc))

	mit := merged.Iterator()
	for mit.HasNext() {
		v, _ := mit.Next()
		require.Contains(t, mergedLists, int(v))
	}
}
