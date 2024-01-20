package hashmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bondhan/hashmap/randstr"
)

func TestNewHashMap(t *testing.T) {
	n := NewHashMap()

	assert.Equal(t, len(n.Data), SIZE)
	for _, v := range n.Data {
		assert.Nil(t, v)
	}

	n = nil

	n = NewHashMap(10)
	assert.Equal(t, len(n.Data), 10)
	for _, v := range n.Data {
		assert.Nil(t, v)
	}

	n = nil

	n = NewHashMap(0)
	assert.Equal(t, len(n.Data), SIZE)
	for _, v := range n.Data {
		assert.Nil(t, v)
	}

}

func TestHash(t *testing.T) {
	v := Hash("a")
	assert.Equal(t, uint32(0x61), v)

	v = Hash("A")
	assert.Equal(t, uint32(0x41), v)
}

func TestInsertNewItem(t *testing.T) {
	hm := NewHashMap()
	err := hm.Put("a", "a")

	assert.Nil(t, err)

	v, err := hm.Get("a")
	assert.Nil(t, err)

	assert.Equal(t, "a", v)
}

func TestInsertReplace(t *testing.T) {
	hm := NewHashMap()
	err := hm.Put("a", "a")
	assert.Nil(t, err)

	err = hm.Put("a", "b")
	v, err := hm.Get("a")
	assert.Nil(t, err)

	assert.Equal(t, "b", v)

	err = hm.Put("a", "c")
	v, err = hm.Get("a")
	assert.Nil(t, err)

	assert.Equal(t, "c", v)
}

func TestPermutableKey(t *testing.T) {
	i, err := getIndex("aab")
	assert.Nil(t, err)

	t.Log("i=", i)

	i, err = getIndex("aba")
	assert.Nil(t, err)

	t.Log("i=", i)
}

func TestSameIndex(t *testing.T) {
	key := []string{"LNiMxseVng", "dg6YXtgutS", "nMoLYAacSx"}

	hm := NewHashMap()
	for _, v := range key {
		val := randstr.String(10)
		err := hm.Put(v, val)
		if err != nil {
			return
		}
	}

	t.Log(hm)
}

func TestRandomPut(t *testing.T) {
	m := make(map[string]string)
	hm := NewHashMap()

	for i := 0; i < 10000; i++ {
		key := randstr.String(10)
		val := randstr.String(10)

		_, err := getIndex(key)
		if err != nil {
			continue
		}

		err = hm.Put(key, val)
		if err != nil {
			t.Error(err)
		}
		m[key] = val
	}

	for k, v := range m {
		vv, err := hm.Get(k)
		assert.Nil(t, err)
		assert.Equal(t, vv, v)
	}
}

func TestRemove(t *testing.T) {
	key := []string{"a", "b", "c"}

	hm := NewHashMap()
	for _, v := range key {
		err := hm.Put(v, v)
		if err != nil {
			return
		}
	}

	g := hm.Remove("a")

	assert.NotNil(t, g)
	assert.Equal(t, "a", *g)

	g = hm.Remove("b")
	assert.NotNil(t, g)
	assert.Equal(t, "b", *g)

	g = hm.Remove("c")
	assert.NotNil(t, g)
	assert.Equal(t, "c", *g)
}

func TestRemoveLinkedlist(t *testing.T) {
	key := []string{"LNiMxseVng", "dg6YXtgutS", "nMoLYAacSx"}

	hm := NewHashMap()
	for k, v := range key {
		err := hm.Put(v, fmt.Sprint(k))
		if err != nil {
			return
		}
	}

	isE := hm.IsEmpty()
	assert.Equal(t, false, isE)

	g := hm.Remove("LNiMxseVng")

	assert.NotNil(t, g)
	assert.Equal(t, "0", *g)

	g = hm.Remove("nMoLYAacSx")

	assert.NotNil(t, g)
	assert.Equal(t, "2", *g)

	g = hm.Remove("dg6YXtgutS")

	assert.NotNil(t, g)
	assert.Equal(t, "1", *g)

	isE = hm.IsEmpty()
	assert.Equal(t, true, isE)
}

func TestRandomRemove(t *testing.T) {
	m := make(map[string]string)
	hm := NewHashMap()

	for i := 0; i < 10000; i++ {
		key := randstr.String(10)
		val := randstr.String(10)

		_, err := getIndex(key)
		if err != nil {
			continue
		}

		err = hm.Put(key, val)
		if err != nil {
			t.Error(err)
		}
		m[key] = val
	}

	for k, _ := range m {
		vv := hm.Remove(k)
		assert.NotNil(t, vv)
	}

	isE := hm.IsEmpty()
	assert.Equal(t, true, isE)
}
