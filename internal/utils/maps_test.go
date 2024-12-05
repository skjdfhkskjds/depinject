package utils_test

import (
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/testutils"
	"github.com/skjdfhkskjds/depinject/internal/utils"
)

func TestOrderedMap(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		om := utils.NewOrderedMap[string, int]()
		testutils.RequireEquals(t, 0, om.Len())
		testutils.RequireLen(t, om.Keys(), 0)
		testutils.RequireLen(t, om.Values(), 0)
	})

	t.Run("set and get", func(t *testing.T) {
		om := utils.NewOrderedMap[string, int]()
		om.Set("a", 1)
		om.Set("b", 2)
		om.Set("c", 3)

		val, ok := om.Get("b")
		testutils.RequireTrue(t, ok)
		testutils.RequireEquals(t, 2, val)

		val, ok = om.Get("d")
		testutils.RequireFalse(t, ok)
		testutils.RequireEquals(t, 0, val)
	})

	t.Run("maintains order", func(t *testing.T) {
		om := utils.NewOrderedMap[string, int]()
		om.Set("c", 3)
		om.Set("a", 1)
		om.Set("b", 2)

		keys := om.Keys()
		testutils.RequireLen(t, keys, 3)
		testutils.RequireEquals(t, "c", keys[0])
		testutils.RequireEquals(t, "a", keys[1])
		testutils.RequireEquals(t, "b", keys[2])

		values := om.Values()
		testutils.RequireLen(t, values, 3)
		testutils.RequireEquals(t, 3, values[0])
		testutils.RequireEquals(t, 1, values[1])
		testutils.RequireEquals(t, 2, values[2])
	})

	t.Run("filter", func(t *testing.T) {
		om := utils.NewOrderedMap[string, int]()
		om.Set("a", 1)
		om.Set("b", 2)
		om.Set("c", 3)

		filtered := om.Filter(func(k string) bool {
			return k != "b"
		})

		testutils.RequireEquals(t, 2, filtered.Len())
		keys := filtered.Keys()
		testutils.RequireEquals(t, "a", keys[0])
		testutils.RequireEquals(t, "c", keys[1])
	})
}
