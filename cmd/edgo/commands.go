package main

import "fmt"

func (ed *Editor) cmdQuit() error {
	ed.running = false
	return nil
}

func (ed *Editor) cmdTogglePrompt() error {
	ed.isPromptShown = !ed.isPromptShown
	return nil
}

func (ed *Editor) cmdToggleShowFullError() error {
	ed.isFullErrorShown = !ed.isFullErrorShown

	if ed.isFullErrorShown {
		ed.cmdShowLastError()
	}

	return nil
}

func (ed *Editor) cmdShowLastError() error {
	if ed.lastError != nil {
		fmt.Fprintln(ed.writer, ed.lastError)
	}
	return nil
}
