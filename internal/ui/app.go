package ui

import (
	"github.com/charmbracelet/bubbletea"
	"shiyin/internal/data"
)

type appState int

const (
	stateSelector appState = iota
	stateViewer
)

// AppModel is the top-level Bubble Tea model.
type AppModel struct {
	state    appState
	selector SelectorModel
	viewer   ViewerModel
}

// NewAppModel creates the app starting at the collection selector.
func NewAppModel() AppModel {
	return AppModel{
		state:    stateSelector,
		selector: NewSelectorModel(),
	}
}

// NewAppModelWithCollection creates the app starting directly in viewer mode.
func NewAppModelWithCollection(key, name string, poems []data.Poem) AppModel {
	return AppModel{
		state:    stateViewer,
		selector: NewSelectorModel(),
		viewer:   NewViewerModel(poems, name),
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.WindowSize()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Global quit
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.state == stateViewer {
				// Return to selector
				m.state = stateSelector
				m.selector = NewSelectorModel()
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		// Propagate to both sub-models
		m.selector, _ = m.selector.Update(msg)
		m.viewer, _ = m.viewer.Update(msg)
		return m, nil
	}

	switch m.state {
	case stateSelector:
		var cmd tea.Cmd
		m.selector, cmd = m.selector.Update(msg)

		// Check if user made a selection
		if key := m.selector.ChosenKey(); key != "" {
			poems, err := data.Load(key)
			if err == nil && len(poems) > 0 {
				m.viewer = NewViewerModel(poems, m.selector.ChosenName())
				m.state = stateViewer
			}
		}
		return m, cmd

	case stateViewer:
		var cmd tea.Cmd
		m.viewer, cmd = m.viewer.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m AppModel) View() string {
	switch m.state {
	case stateSelector:
		return m.selector.View()
	case stateViewer:
		return m.viewer.View()
	}
	return ""
}
