package id_test

import (
	. "github.com/nasermirzaei89/id"
	"testing"
)

func TestNew(t *testing.T) {
	id := New()
	if l := len(id); l != 40 {
		t.Errorf("generated string length is '%d' not '40'", l)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New()
	}
}
