package hashmap

import (
	"errors"
	"hash/fnv"
)

const SIZE = 100

type Node struct {
	key  string
	val  string
	next *Node
}

type HashMap struct {
	Data []*Node
}

func NewHashMap() *HashMap {
	return &HashMap{Data: make([]*Node, SIZE)}
}

func Hash(key string) (uint32, error) {
	hash := fnv.New32a()
	write, err := hash.Write([]byte(key))
	if err != nil {
		return 0, err
	}

	if write <= 0 {
		return 0, errors.New("invalid hash")
	}

	return hash.Sum32(), nil
}

func getIndex(key string) (int, error) {
	hash, err := Hash(key)
	if err != nil {
		return 0, err
	}

	index := hash % SIZE

	if index < 0 || index >= SIZE {
		return 0, errors.New("out of boundary")
	}

	return int(index), nil
}

func (h *HashMap) Insert(key, value string) error {

	index, err := getIndex(key)
	if err != nil {
		return err
	}

	curr := h.Data[index]
	n := &Node{
		key:  key,
		val:  value,
		next: nil,
	}

	// if new data
	if curr == nil {
		h.Data[index] = n
		return nil
	}

	// if array not empty
	for {

		if curr.key == key {
			h.Data[index].val = value //replace
			return nil
		}

		if curr.next == nil {
			curr.next = n
			return nil
		}

		curr = curr.next
	}
}

func (h *HashMap) Value(key string) (string, error) {
	index, err := getIndex(key)
	if err != nil {
		return "", err
	}

	curr := h.Data[index]

	if curr == nil {
		return "", nil
	}

	for {
		if curr.key == key {
			return curr.val, nil
		}

		if curr.next == nil {
			return "", nil
		}
		curr = curr.next
	}
}
