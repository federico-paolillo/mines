package menus

import (
	"fmt"
	"strconv"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
)

type entriesMap = map[string]*Entry

type Entry struct {
	Prompt string
	Dialog *dialog.Dialog
}

func NewMenu(
	console *console.Console,
	entries []Entry,
) *dialog.Dialog {
	loneStep, entriesMap := makeMenuStep(entries)

	return &dialog.Dialog{
		Console: console,
		Steps: []dialog.Step{
			loneStep,
		},
		OnCompleteInteraction: func(inputs dialog.Inputs) {
			selection := inputs["selection"]
			entry := entriesMap[selection]
			entry.Dialog.Interact(inputs)
		},
	}
}

func makeMenuStep(
	entries []Entry,
) (dialog.Step, entriesMap) {
	entriesMap := make(entriesMap, 0)
	prompt := make([]string, 0, len(entries))

	for i, entry := range entries {
		entryIndex := strconv.FormatInt(int64(i+1), 10)
		entryPrompt := fmt.Sprintf("%s %s", entryIndex, entry.Prompt)

		entriesMap[entryIndex] = &entry

		prompt = append(prompt, entryPrompt)
	}

	return dialog.Step{
		Prompt: prompt,
		Name:   "selection",
		Validate: func(value string) bool {
			_, ok := entriesMap[value]
			return ok
		},
	}, entriesMap
}
