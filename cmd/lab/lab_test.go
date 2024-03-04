package main

import (
	"net/http"
	"testing"
)

func TestHeaderToString(t *testing.T) {
	t.Run("normal headers", func(t *testing.T) {
		header := http.Header{
			"Header1": {"Value1, Value2"},
		}

		got := headerToString(&header)
		want := "Header1: Value1, Value2\n"

		assertHeader(t, got, want, header)
	})

	// http.Header is case insensitive and puts same header's values in a slice
	// e.g. Header1: ValueA
	//      header1: ValueB
	// becomes
	//      map{"Header1": ["ValueA", "ValueB"]}
	// so we need to be able to also iterate over a slice with many strings and format accordingly
	t.Run("formatted headers", func(t *testing.T) {
		header := http.Header{
			"Header1": {"ValueA", "ValueB"},
		}

		got := headerToString(&header)
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
