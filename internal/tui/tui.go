package tui

import (
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("8")).
	Margin(1, 2)

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
	List list.Model
}

func New(items []list.Item) *Model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetFilteringEnabled(true)
	l.SetShowFilter(true)

	return &Model{
		List: l,
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
				log.Fatalf("tui: error running make command: %s", err)
			}
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		horizontal, vertical := style.GetFrameSize()
		listTotalItems := len(m.List.Items())
		m.List.SetWidth(msg.Width - horizontal)
		m.List.SetHeight(4 * listTotalItems)

		style.Width(msg.Width - horizontal)
		style.Height(4*listTotalItems - vertical)

	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return style.Render(m.List.View())
}
