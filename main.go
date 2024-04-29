package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycruz/makeit/internal/target"
	"github.com/nycruz/makeit/internal/tui"
)

func main() {
	var makefileName = flag.String("f", "Makefile", "Name of Makefile. Default: Makefile")
	var isVerbose = flag.Bool("v", false, "Prints the command")
	flag.Parse()

	t := target.New()
	targets, err := t.GetMakefileTargets(*makefileName)
	if err != nil {
		log.Fatalf("could not get targets: %v", err)
	}

	if len(targets) > 0 {
		var items []list.Item
		for _, v := range targets {
			desc := v.Description
			if *isVerbose {
				desc = fmt.Sprintf("%s: `%s`", v.Description, v.Command)
			}

			items = append(items, tui.Item{
				Name: v.Name,
				Desc: desc,
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
