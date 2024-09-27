package menus

import (
	"fmt"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
)

type entriesMap = map[string]*Entry

type Entry struct {
	Prompt string
	Dialog *tui.Dialog
}

func NewMenu(
	console *tui.Console,
	dispatcher *tui.Dispatcher,
	entries []Entry,
) *tui.Dialog {
	entriesMap := makeEntriesMap(entries)
	steps := makeMenuSteps(entriesMap)

	return &tui.Dialog{
		Console: console,
		Steps:   steps,
		OnCompleteInteraction: func(inputs tui.Inputs) {
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
) []tui.Step {
	prompt := make([]string, 0, len(entries))

	for i, entry := range entries {
		entryPrompt := fmt.Sprintf("%s %s", i, entry.Prompt)
		prompt = append(prompt, entryPrompt)
	}

	return []tui.Step{
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
