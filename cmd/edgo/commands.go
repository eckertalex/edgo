package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

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

func (ed *Editor) cmdEdit(filename string) error {
	ed.filename = filepath.ToSlash(filename)

	stat, err := os.Stat(ed.filename)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", ed.filename)
		return nil
	}

	if err != nil {
		return fmt.Errorf("error stating file: %w", err)
	}

	ed.buffer.Clear()
	fmt.Fprintf(os.Stdout, "%d\n", stat.Size())

	file, err := os.Open(ed.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ed.buffer.Append(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
