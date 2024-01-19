package hashmap

import (
	"errors"
)

const SIZE = 64

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

func Hash(key string) uint32 {
	total := uint32(0)
	for _, c := range key {
		total += uint32(c)
	}

	return total
}

func getIndex(key string) (int, error) {
	hash := Hash(key)

	index := hash % SIZE

	if index < 0 || index >= SIZE {
		return 0, errors.New("out of boundary")
	}

	return int(index), nil
}

func (h *HashMap) Put(key, value string) error {

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

func (h *HashMap) Get(key string) (string, error) {
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

func (h *HashMap) Remove(key string) *string {
	index, err := getIndex(key)
	if err != nil {
		return nil
	}

	curr := h.Data[index]

	if curr == nil {
		return nil
	}

	prev := curr

	for {
		if curr.key == key {
			if prev == curr && curr != nil && curr.next != nil {
				h.Data[index] = curr.next
				return &curr.val
			}

			if prev == curr && curr != nil && curr.next == nil {
				h.Data[index] = nil
				return &curr.val
			}

			if prev != curr && curr.next != nil {
				prev.next = curr.next
				return &curr.val
			}

			if prev != curr && curr.next == nil {
				prev.next = nil
				return &curr.val
			}
		}

		if curr.next == nil {
			return nil
		}

		prev = curr
		curr = curr.next
	}
}

func (h *HashMap) IsEmpty() bool {
	for _, v := range h.Data {
		if v != nil {
			return false
		}
	}

	return true
}
