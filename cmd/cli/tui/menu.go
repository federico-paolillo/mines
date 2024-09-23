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
	Console *Console
	Steps   []Step
	Execute func(inputs Inputs)
}

func (d *Dialog) Interact(prevInputs Inputs) {
	inputs := make(Inputs)

	maps.Copy(inputs, prevInputs)

	for _, step := range d.Steps {
		d.RenderPrompt(step)
		d.GatherInput(step, inputs)
	}

	d.Execute(inputs)
}

func (d *Dialog) RenderPrompt(step Step) {
	for _, prompt := range step.Prompt {
		d.Console.Printline(prompt)
	}
}

func (d *Dialog) GatherInput(step Step, destination Inputs) {
	for {
		in := d.Console.Scanline()

		ok := step.Validate(in)

		if ok {
			destination[step.Name] = in
			break
		}

		d.Console.Printline("huh?")
	}
}
