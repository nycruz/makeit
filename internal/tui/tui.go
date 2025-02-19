package tui

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var style = lipgloss.NewStyle().
	Margin(0, 0, 1, 0).
	Padding(1, 1)

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
	err  error
}

func New(items []list.Item) *Model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetFilteringEnabled(true)
	l.SetShowFilter(true)

	return &Model{
		List: l,
		err:  nil,
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
			if err := cmd.Run(); err != nil {
				m.err = fmt.Errorf("'make' command failed: %v", err)
				return m, tea.Quit
			}

			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		horizontal, vertical := style.GetFrameSize()

		style.Width(msg.Width - horizontal)
		m.List.SetWidth(msg.Width - horizontal)

		// Calculate the height of the list based on the number of items
		itemCount := len(m.List.Items())
		itemHeight := 4 // Height per item
		totalHeight := itemCount * itemHeight

		if totalHeight > msg.Height-vertical {
			totalHeight = msg.Height - vertical
		}

		style.Height(totalHeight)
		m.List.SetHeight(totalHeight)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	return style.Render(m.List.View())
}
