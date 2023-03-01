package main

import (
	"flag"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycruz/makeit/internal/target"
	"github.com/nycruz/makeit/internal/ui"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("fatal %v", err)
	}
	defer f.Close()

	var makefileName = flag.String("f", "Makefile", "Name of Makefile. Default: Makefile")

	t := target.NewTarget()
	targets, err := t.GetMakefileTargets(*makefileName)
	if err != nil {
		log.Fatalf("could not get targets: %v", err)
	}

	var items []list.Item
	for _, v := range targets {
		items = append(items, ui.Item{
			Name: v.Name,
			Desc: v.Description,
		})
	}

	m := ui.NewModel(items, f)
	m.List.Title = "Select Target"
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("error running makeit: %v", err)
	}

}
