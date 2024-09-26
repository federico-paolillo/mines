package tui_test

import (
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
)

func TestDispatcherCallsAllSubscribers(t *testing.T) {
	counter := 0

	l1 := func(_ any) {
		counter++
	}

	l2 := func(_ any) {
		counter++
	}

	d := tui.NewDispatcher()

	_ = d.Subscribe(l1)
	_ = d.Subscribe(l2)

	expectedCalls := 2

	d.Dispatch(0)

	if counter != expectedCalls {
		t.Fatalf(
			"expected %d subscribers calls. got %d",
			expectedCalls,
			counter,
		)
	}
}

func TestDispatcherDoesNotCallUnsubscribedSubscribers(t *testing.T) {
	counter := 0

	l1 := func(_ any) {
		counter++
	}

	l2 := func(_ any) {
		counter++
	}

	d := tui.NewDispatcher()

	_ = d.Subscribe(l1)
	unsub := d.Subscribe(l2)

	d.Dispatch(0)

	unsub()

	d.Dispatch(0)

	expectedCalls := 3

	if counter != expectedCalls {
		t.Fatalf(
			"expected %d subscribers calls. got %d",
			expectedCalls,
			counter,
		)
	}
}
