package id_test

import (
	. "github.com/nasermirzaei89/id"
	"testing"
)

func TestNew(t *testing.T) {
	id := New()
	if l := len(id); l != 40 {
		t.Errorf("generated strign length is '%d' not '40'", l)
	}
}

func BenchmarkNew(b *testing.B) {
	ids := make([]string, 0)
	for {
		l := len(ids)
		ids = append(ids, New())
		for i := range ids {
			if i < l && ids[i] == ids[l] {
				b.Errorf("generated duplicate id in round %d and %d", i, l)
			}
		}

		if len(ids) > b.N {
			break
		}
	}
}
