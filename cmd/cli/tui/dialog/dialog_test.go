package dialog_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
)

func TestDialogRendersSteps(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"input\n",
	)

	testConsole := console.NewConsole(
		stdin,
		&stdout,
	)

	d := dialog.Dialog{
		testConsole,
		[]dialog.Step{
			{
				Prompt: []string{
					"prompt_1",
					"prompt_2",
				},
				Name: "input_a",
			},
		},
		func(_ dialog.Inputs) {

		},
	}

	d.Interact(dialog.NoInputs)

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

	testConsole := console.NewConsole(
		stdin,
		&stdout,
	)

	d := dialog.Dialog{
		testConsole,
		[]dialog.Step{
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
		func(_ dialog.Inputs) {

		},
	}

	d.Interact(dialog.NoInputs)

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

	testConsole := console.NewConsole(
		stdin,
		&stdout,
	)

	var inputs dialog.Inputs

	d := dialog.Dialog{
		testConsole,
		[]dialog.Step{
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
		func(i dialog.Inputs) {
			inputs = i
		},
	}

	d.Interact(dialog.NoInputs)

	screen := stdout.String()

	screenExpected := "prompt_1\nprompt_2\nhuh?\n"

	if screen != screenExpected {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screen, screenExpected)
	}

	inputsExpected := dialog.Inputs{
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

	testConsole := console.NewConsole(
		stdin,
		&stdout,
	)

	chosenEntry := ""

	d := dialog.Dialog{
		testConsole,
		[]dialog.Step{
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
		func(i dialog.Inputs) {
			chosenEntry = i["choice"]
		},
	}

	d.Interact(dialog.NoInputs)

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

	testConsole := console.NewConsole(
		stdin,
		&stdout,
	)

	didConfirm := false

	confirmMenu := dialog.Dialog{
		testConsole,
		[]dialog.Step{
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
		func(inputs dialog.Inputs) {
			didConfirm = inputs["confirm"] == "y"
		},
	}

	inputMenu := dialog.Dialog{
		testConsole,
		[]dialog.Step{
			{
				Prompt: []string{
					"enter x",
				},
				Name: "x",
			},
		},
		func(inputs dialog.Inputs) {
			confirmMenu.Interact(inputs)
		},
	}

	mainMenu := dialog.Dialog{
		testConsole,
		[]dialog.Step{
			{
				Prompt: []string{
					"1. quit",
				},
			},
		},
		func(inputs dialog.Inputs) {
			inputMenu.Interact(inputs)
		},
	}

	mainMenu.Interact(dialog.NoInputs)

	if !didConfirm {
		t.Errorf("did not confirm")
	}
}
