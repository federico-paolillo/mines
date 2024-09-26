package tui

import "maps"

type Inputs map[string]string

var NoInputs Inputs

type Step struct {
	Prompt   []string
	Name     string
	Validate func(value string) bool
}

type Dialog struct {
	Console               *Console
	Steps                 []Step
	OnCompleteInteraction func(inputs Inputs)
}

func (d *Dialog) Interact(prevInputs Inputs) {
	inputs := make(Inputs)

	maps.Copy(inputs, prevInputs)

	for _, step := range d.Steps {
		d.renderPrompt(step)
		d.gatherInput(step, inputs)
	}

	d.OnCompleteInteraction(inputs)
}

func (d *Dialog) renderPrompt(step Step) {
	for _, prompt := range step.Prompt {
		d.Console.Printline(prompt)
	}
}

func (d *Dialog) gatherInput(step Step, destination Inputs) {
	for {
		in := d.Console.Scanline()

		ok := true
		inputHasName := step.Name != ""

		if step.Validate != nil {
			ok = step.Validate(in)
		}

		if ok {
			if inputHasName {
				destination[step.Name] = in
			}

			break
		}

		d.Console.Printline("huh?")
	}
}
