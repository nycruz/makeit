package main

import (
	"flag"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycruz/makeit/internal/target"
	"github.com/nycruz/makeit/internal/tui"
)

func main() {
	var makefileName = flag.String("f", "Makefile", "Name of Makefile. Default: Makefile")
	// var isDebug = flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	t := target.New()
	targets, err := t.GetMakefileTargets(*makefileName)
	if err != nil {
		log.Fatalf("could not get targets: %v", err)
	}

	if len(targets) > 0 {
		var items []list.Item
		for _, v := range targets {
			items = append(items, tui.Item{
				Name: v.Name,
				Desc: v.Description,
			})
		}

		m := tui.New(items)
		m.List.Title = "Select Target"
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			log.Fatalf("error running makeit: %v", err)
		}
	} else {
		log.Fatalln("no targets found")
	}

}
