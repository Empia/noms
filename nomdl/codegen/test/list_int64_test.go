package test

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestListInt64Def(t *testing.T) {
	assert := assert.New(t)

	def := ListOfInt64Def{}
	l := def.New()

	def2 := l.Def()
	l2 := def.New()

	assert.Equal(def, def2)
	assert.True(l.Equals(l2))

	l3 := NewListOfInt64()
	assert.True(l.Equals(l3))

	def3 := ListOfInt64Def{0, 1, 2, 3, 4}
	l4 := def3.New()
	assert.Equal(uint64(5), l4.Len())
	assert.Equal(int64(0), l4.Get(0))
	assert.Equal(int64(2), l4.Get(2))
	assert.Equal(int64(4), l4.Get(4))

	l4 = l4.Set(4, 44).Slice(3, 5)
	assert.Equal(ListOfInt64Def{3, 44}, l4.Def())
}

func TestListIter(t *testing.T) {
	l := ListOfInt64Def{0, 1, 2, 3, 4}.New()
	acc := ListOfInt64Def{}
	l.Iter(func(v int64) (stop bool) {
		stop = v == 2
		acc = append(acc, v)
		return
	})
	assert.Equal(t, ListOfInt64Def{0, 1, 2}, acc)
}

func TestListFilter(t *testing.T) {
	l := ListOfInt64Def{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}.New()
	l2 := l.Filter(func(v int64) bool {
		return v%2 == 0
	})
	assert.Equal(t, ListOfInt64Def{0, 2, 4, 6, 8}, l2.Def())
}