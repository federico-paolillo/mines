package id_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/server/id"
)

func TestGeneratesTwoDifferentIdsWhenCalledTwice(t *testing.T) {
	id1, err := id.Generate()

	if err != nil {
		t.Fatalf(
			"could not generate id. %v",
			err,
		)
	}

	id2, err := id.Generate()

	if err != nil {
		t.Fatalf(
			"could not generate id. %v",
			err,
		)
	}

	if id1 == id2 {
		t.Fatalf(
			"generated the same id twice. first id is '%s' and second id is '%s'",
			id1,
			id2,
		)
	}
}
