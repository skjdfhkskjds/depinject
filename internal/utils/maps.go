package utils

type OrderedMap[K comparable, V any] struct {
	keys []K
	m    map[K]V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		m: make(map[K]V),
	}
}

func (om *OrderedMap[K, V]) Set(k K, v V) {
	// If the key doesn't exist, add it to the keys slice.
	if _, ok := om.m[k]; !ok {
		om.keys = append(om.keys, k)
	}

	om.m[k] = v
}

func (om *OrderedMap[K, V]) Get(k K) (V, bool) {
	v, ok := om.m[k]
	return v, ok
}

func (om *OrderedMap[K, V]) Keys() []K {
	return om.keys
}

func (om *OrderedMap[K, V]) Values() []V {
	values := make([]V, len(om.keys))
	for i, k := range om.keys {
		values[i] = om.m[k]
	}
	return values
}

func (om *OrderedMap[K, V]) Len() int {
	return len(om.keys)
}

// FilterMap returns a new map with all the keys that satisfy the predicate.
func (om *OrderedMap[K, V]) Filter(predicate func(K) bool) *OrderedMap[K, V] {
	result := NewOrderedMap[K, V]()
	for _, k := range om.keys {
		if predicate(k) {
			result.Set(k, om.m[k])
		}
	}
	return result
}
