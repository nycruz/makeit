package ui

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Margin(1, 2).
	BorderStyle(lipgloss.NormalBorder())

type Item struct {
	Name string
	Desc string
}

func (i Item) Title() string {
	return i.Name
}

func (i Item) Description() string {
	return i.Desc
}

func (i Item) FilterValue() string {
	return i.Name
}

type Model struct {
	List   list.Model
	logger *os.File
}

func NewModel(items []list.Item, logger *os.File) *Model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)

	return &Model{
		List:   l,
		logger: logger,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			cmd := exec.Command("make", m.List.SelectedItem().FilterValue())
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				log.Fatalf("error running make command: %v", err)
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		horizontal, _ := style.GetFrameSize()
		listTotalItems := len(m.List.Items())
		m.logger.Write([]byte(fmt.Sprintf("listTotalItems: %d", listTotalItems)))
		m.List.SetSize(msg.Width-horizontal, 4*listTotalItems)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return style.Render(m.List.View())
}
