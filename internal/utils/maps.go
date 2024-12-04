package utils

// FilterMap returns a new map with all the keys that satisfy the predicate.
func FilterMap[K comparable, V any](m map[K]V, predicate func(K) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k) {
			result[k] = v
		}
	}
	return result
}
