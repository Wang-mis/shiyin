package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"shiyin/internal/data"
)

// collectionItem implements list.Item for bubbles/list.
type collectionItem struct {
	key   string
	title string
}

func (i collectionItem) Title() string       { return i.title }
func (i collectionItem) Description() string { return "" }
func (i collectionItem) FilterValue() string { return i.title }

// SelectorModel is the collection selection screen.
type SelectorModel struct {
	list   list.Model
	chosen string // non-empty once user selects
	width  int
	height int
}

func NewSelectorModel() SelectorModel {
	items := []list.Item{
		collectionItem{key: "tang300", title: "唐诗三百首"},
		collectionItem{key: "ci300", title: "宋词三百首"},
		collectionItem{key: "all", title: "全部 (唐诗 + 宋词)"},
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DADADA")).
		Padding(0, 0, 0, 2)
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8C87A")).
		Padding(0, 0, 0, 1).
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color("#E8C87A"))

	l := list.New(items, delegate, 30, 10)
	l.Title = "选择诗词集合"
	l.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#AAAAAA")).
		Padding(1, 0, 1, 2)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return SelectorModel{list: l}
}

func (m SelectorModel) Init() tea.Cmd {
	return nil
}

func (m SelectorModel) Update(msg tea.Msg) (SelectorModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width, msg.Height-4)

	case tea.KeyMsg:
		if msg.String() == "enter" {
			if item, ok := m.list.SelectedItem().(collectionItem); ok {
				m.chosen = item.key
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SelectorModel) View() string {
	return m.list.View()
}

// ChosenKey returns the selected collection key, empty if none yet.
func (m SelectorModel) ChosenKey() string {
	return m.chosen
}

// ChosenName returns the display name of the selected collection.
func (m SelectorModel) ChosenName() string {
	switch m.chosen {
	case "tang300":
		return data.AvailableCollections[0].Name
	case "ci300":
		return data.AvailableCollections[1].Name
	case "all":
		return "全部"
	}
	return ""
}
