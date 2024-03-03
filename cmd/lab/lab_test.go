package main

import (
	"net/http"
	"testing"
)

func TestHeaderToString(t *testing.T) {
	t.Run("CanoncicalHeaderForm", func(t *testing.T) {
		header := http.Header{
			"Header1": {"Value1, Value2"},
			"Header2": {"Another Value"},
		}

		got := HeaderToString(&header)
		want := "Header1: Value1, Value2\nHeader2: Another Value\n"

		assertHeader(t, got, want, header)
	})

	// http.Header is case insensitive and puts same header's values in a slice
	t.Run("formated headers", func(t *testing.T) {
		header := http.Header{
			"Header1": {"ValueA", "ValueB"},
		}

		got := HeaderToString(&header)
		want := "Header1: ValueA, ValueB\n"

		assertHeader(t, got, want, header)
	})
}

func assertHeader(t testing.TB, got, want string, header http.Header) {
	t.Helper()

	if got != want {
		t.Errorf("got:\n%s but want:\n%s given:\n%v", got, want, header)
	}
}
