package tui_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
)

func TestDialogRendersSteps(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"input\n",
	)

	testConsole := tui.NewConsole(
		stdin,
		&stdout,
	)

	d := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				[]string{
					"prompt_1",
					"prompt_2",
				},
				"input_a",
				func(_ string) bool {
					return true
				},
			},
		},
		func(_ tui.Inputs) {

		},
	}

	d.Interact(tui.NoInputs)

	screen := stdout.String()

	expectation := "prompt_1\nprompt_2\n"

	if screen != expectation {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screen, expectation)
	}
}

func TestDialogRendersHuhWhenInvalidInput(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"wrong\nok\n",
	)

	testConsole := tui.NewConsole(
		stdin,
		&stdout,
	)

	d := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				[]string{
					"prompt_1",
					"prompt_2",
				},
				"input_a",
				func(in string) bool {
					return in == "ok"
				},
			},
		},
		func(_ tui.Inputs) {

		},
	}

	d.Interact(tui.NoInputs)

	screen := stdout.String()

	expectation := "prompt_1\nprompt_2\nhuh?\n"

	if screen != expectation {
		t.Errorf("dialog did not render expected output. wanted '%sq got '%q'", screen, expectation)
	}
}

func TestDialogCollectsInputCorrectly(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"wrong\nabc\n",
	)

	testConsole := tui.NewConsole(
		stdin,
		&stdout,
	)

	var inputs tui.Inputs

	d := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				[]string{
					"prompt_1",
					"prompt_2",
				},
				"input_a",
				func(in string) bool {
					return in == "abc"
				},
			},
		},
		func(i tui.Inputs) {
			inputs = i
		},
	}

	d.Interact(tui.NoInputs)

	screen := stdout.String()

	screenExpected := "prompt_1\nprompt_2\nhuh?\n"

	if screen != screenExpected {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screen, screenExpected)
	}

	inputsExpected := tui.Inputs{
		"input_a": "abc",
	}

	if !reflect.DeepEqual(inputs, inputsExpected) {
		t.Errorf("wrong inputs collected. wanted '%q' got '%q'", inputsExpected, inputs)
	}
}

func TestDialogCanActAsMenu(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"1\n",
	)

	testConsole := tui.NewConsole(
		stdin,
		&stdout,
	)

	chosenEntry := ""

	d := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				[]string{
					"1. do x",
					"2. do y",
				},
				"choice",
				func(in string) bool {
					return in == "1" ||
						in == "2"
				},
			},
		},
		func(i tui.Inputs) {
			chosenEntry = i["choice"]
		},
	}

	d.Interact(tui.NoInputs)

	selectionExpected := "1"

	if selectionExpected != chosenEntry {
		t.Errorf("wrong choice selected. wanted '%q' got '%q'", selectionExpected, chosenEntry)
	}
}

func TestDialogCanConstructComplexFlow(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"1\n1234\ny\n",
	)

	testConsole := tui.NewConsole(
		stdin,
		&stdout,
	)

	didConfirm := false

	confirmMenu := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				Prompt: []string{
					"confirm? (y/n)",
				},
				Name: "confirm",
				Validate: func(value string) bool {
					return value == "y" || value == "n"
				},
			},
		},
		func(inputs tui.Inputs) {
			didConfirm = inputs["confirm"] == "y"
		},
	}

	inputMenu := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				Prompt: []string{
					"enter x",
				},
				Name: "x",
				Validate: func(_ string) bool {
					return true
				},
			},
		},
		func(inputs tui.Inputs) {
			confirmMenu.Interact(inputs)
		},
	}

	mainMenu := tui.Dialog{
		testConsole,
		[]tui.Step{
			{
				Prompt: []string{
					"1. quit",
				},
				Name: "",
				Validate: func(value string) bool {
					return true
				},
			},
		},
		func(inputs tui.Inputs) {
			inputMenu.Interact(inputs)
		},
	}

	mainMenu.Interact(tui.NoInputs)

	if !didConfirm {
		t.Errorf("did not confirm")
	}
}
