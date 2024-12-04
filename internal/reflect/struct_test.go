package reflect_test

import (
	"testing"

	"github.com/skjdfhkskjds/depinject/internal/reflect"
	"github.com/skjdfhkskjds/depinject/internal/testutils"
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
		testutils.RequireNoError(t, err)
		testutils.RequireEquals(t, "testStruct", s.Name)
		testutils.RequireLen(t, s.Fields, 3)
		testutils.RequireEquals(t, reflect.TypeOf(""), s.Fields[0])
		testutils.RequireEquals(t, reflect.TypeOf(0), s.Fields[1])
		testutils.RequireEquals(t, reflect.TypeOf(false), s.Fields[2])
	})

	t.Run("not a struct", func(t *testing.T) {
		_, err := reflect.NewStruct(42)
		testutils.RequireErrorIs(t, err, reflect.ErrNotAStruct)
	})
}

func TestStruct_Constructor(t *testing.T) {
	testStructType := reflect.TypeOf(testStruct{})
	s, err := reflect.NewStruct(testStructType)
	testutils.RequireNoError(t, err)

	constructor := s.Constructor()
	err = constructor.Call(false, "test", 42, true)
	testutils.RequireNoError(t, err)

	result := constructor.Ret
	testutils.RequireLen(t, result, 1)
	constructedStruct, ok := result[testStructType].Interface().(testStruct)
	testutils.RequireTrue(t, ok)

	testutils.RequireEquals(t, "test", constructedStruct.Field1)
	testutils.RequireEquals(t, 42, constructedStruct.Field2)
	testutils.RequireEquals(t, true, constructedStruct.Field3)

	// Call TestMethod and check the result
	testMethodResult := constructedStruct.testMethod()
	testutils.RequireEquals(t, "test", testMethodResult)
}

func TestStruct_Constructor_Multiple(t *testing.T) {
	for i := 0; i < 100; i++ {
		TestStruct_Constructor(t)
	}
}

func TestStruct_Provider(t *testing.T) {
	testStructType := reflect.TypeOf(testStruct{})
	s, err := reflect.NewStruct(testStructType)
	testutils.RequireNoError(t, err)

	provider := s.Provider()
	err = provider.Call(false,
		testStruct{
			Field1: "test",
			Field2: 42,
			Field3: true,
		},
	)
	testutils.RequireNoError(t, err)

	result := provider.Ret
	testutils.RequireLen(t, result, 3)
	testutils.RequireEquals(t, "test", result[reflect.TypeOf("")].Interface())
	testutils.RequireEquals(t, 42, result[reflect.TypeOf(0)].Interface())
	testutils.RequireEquals(t, true, result[reflect.TypeOf(false)].Interface())
}

func TestStruct_Provider_Multiple(t *testing.T) {
	for i := 0; i < 100; i++ {
		TestStruct_Provider(t)
	}
}
