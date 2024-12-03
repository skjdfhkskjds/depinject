package reflect_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
)

type testStruct struct {
	Field1 string
	Field2 int
	Field3 bool
}

func (t *testStruct) testMethod() string {
	return "test"
}

func TestNewStruct(t *testing.T) {
	t.Run("valid struct", func(t *testing.T) {
		testStructType := reflect.TypeOf(testStruct{})
		s, err := reflect.NewStruct(testStructType)
		require.NoError(t, err)
		require.Equal(t, "testStruct", s.Name)
		require.Len(t, s.Fields, 3)
		require.Equal(t, reflect.TypeOf(""), s.Fields[0])
		require.Equal(t, reflect.TypeOf(0), s.Fields[1])
		require.Equal(t, reflect.TypeOf(false), s.Fields[2])
	})

	t.Run("not a struct", func(t *testing.T) {
		_, err := reflect.NewStruct(42)
		require.ErrorIs(t, err, reflect.ErrNotAStruct)
	})
}

func TestStruct_Constructor(t *testing.T) {
	testStructType := reflect.TypeOf(testStruct{})
	s, err := reflect.NewStruct(testStructType)
	require.NoError(t, err)

	constructor := s.Constructor()
	err = constructor.Call(true, "test", 42, true)
	require.NoError(t, err)

	result := constructor.Ret
	require.Len(t, result, 1)
	constructedStruct, ok := result[testStructType].Interface().(testStruct)
	require.True(t, ok)

	require.Equal(t, "test", constructedStruct.Field1)
	require.Equal(t, 42, constructedStruct.Field2)
	require.Equal(t, true, constructedStruct.Field3)

	// Call TestMethod and check the result
	testMethodResult := constructedStruct.testMethod()
	require.Equal(t, "test", testMethodResult)
}
