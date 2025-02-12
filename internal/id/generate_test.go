package id_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/id"
)

func TestGeneratesTwoDifferentIdsWhenCalledTwice(t *testing.T) {
	id1 := id.Generate()
	id2 := id.Generate()

	if id1 == id2 {
		t.Fatalf(
			"generated the same id twice. first id is '%s' and second id is '%s'",
			id1,
			id2,
		)
	}
}
