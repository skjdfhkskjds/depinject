package testutils

import (
	"errors"
	"reflect"
	"testing"
)

func RequireNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func RequireErrorIs(t *testing.T, actual, expected error) {
	if !errors.Is(actual, expected) {
		t.Fatalf("expected error to be %v, got %v", expected, actual)
	}
}

func RequireError(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func RequireNotNil(t *testing.T, value any) {
	if value == nil {
		t.Fatalf("value is nil")
	}
}

func RequireTrue(t *testing.T, value bool) {
	if !value {
		t.Fatalf("expected value to be true")
	}
}

func RequireEquals(t *testing.T, actual, expected any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func RequireLen(t *testing.T, collection any, expectedLen int) {
	if collection == nil {
		t.Fatalf("expected collection to be non-nil")
	}

	length := reflect.ValueOf(collection).Len()
	if length != expectedLen {
		t.Fatalf("expected length %d, got %d", expectedLen, length)
	}
}

func RequireEmpty(t *testing.T, collection any) {
	if reflect.ValueOf(collection).Len() != 0 {
		t.Fatalf("expected collection to be empty")
	}
}
