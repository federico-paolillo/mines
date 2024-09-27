package menus

import (
	"fmt"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
)

type entriesMap = map[string]*Entry

type Entry struct {
	Prompt string
	Dialog *dialog.Dialog
}

func NewMenu(
	console *console.Console,
	dispatcher *dispatcher.Dispatcher,
	entries []Entry,
) *dialog.Dialog {
	entriesMap := makeEntriesMap(entries)
	steps := makeMenuSteps(entriesMap)

	return &dialog.Dialog{
		Console: console,
		Steps:   steps,
		OnCompleteInteraction: func(inputs dialog.Inputs) {
			selection := inputs["selection"]
			entry := entriesMap[selection]
			entry.Dialog.Interact(inputs)
		},
	}
}

func makeEntriesMap(entries []Entry) map[string]*Entry {
	entriesMap := make(entriesMap, 0)

	for i, entry := range entries {
		index := fmt.Sprintf("%d", i+1)

		entriesMap[index] = &entry
	}
	return entriesMap
}

func makeMenuSteps(
	entries entriesMap,
) []dialog.Step {
	prompt := make([]string, 0, len(entries))

	for i, entry := range entries {
		entryPrompt := fmt.Sprintf("%s %s", i, entry.Prompt)
		prompt = append(prompt, entryPrompt)
	}

	return []dialog.Step{
		{
			Prompt: prompt,
			Name:   "selection",
			Validate: func(value string) bool {
				_, ok := entries[value]
				return ok
			},
		},
	}
}
