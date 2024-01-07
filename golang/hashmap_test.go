package hashmap

import (
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
}

func TestHash(t *testing.T) {
	v, e := Hash("")
	assert.Error(t, e)
	assert.Equal(t, uint32(0), uint32(0))

	v, e = Hash("a")
	assert.Nil(t, e)
	assert.Equal(t, uint32(3826002220), v)
}

func TestInsertNewItem(t *testing.T) {
	hm := NewHashMap()
	err := hm.Insert("a", "a")

	assert.Nil(t, err)

	v, err := hm.Value("a")
	assert.Nil(t, err)

	assert.Equal(t, "a", v)
}

func TestInsertReplace(t *testing.T) {
	hm := NewHashMap()
	err := hm.Insert("a", "a")
	assert.Nil(t, err)

	err = hm.Insert("a", "b")
	v, err := hm.Value("a")
	assert.Nil(t, err)

	assert.Equal(t, "b", v)

	err = hm.Insert("a", "c")
	v, err = hm.Value("a")
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
		err := hm.Insert(v, val)
		if err != nil {
			return
		}
	}

	t.Log(hm)
}

func TestRandom(t *testing.T) {
	m := make(map[string]string)
	hm := NewHashMap()

	for i := 0; i < 10000; i++ {
		key := randstr.String(10)
		val := randstr.String(10)

		_, err := getIndex(key)
		if err != nil {
			continue
		}

		err = hm.Insert(key, val)
		if err != nil {
			t.Error(err)
		}
		m[key] = val
	}

	for k, v := range m {
		vv, err := hm.Value(k)
		assert.Nil(t, err)
		assert.Equal(t, vv, v)
	}

}
