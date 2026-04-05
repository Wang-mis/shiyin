package ui

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"shiyin/internal/data"
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E8C87A")).
			Bold(true)

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3D3D3D"))

	metaStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	bodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#DADADA"))

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))

	tooSmallStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555"))
)

const minWidth = 20
const minHeight = 8

type ViewerModel struct {
	poems      []data.Poem
	index      int
	width      int
	height     int
	showHelp   bool
	collection string // display name
}

func NewViewerModel(poems []data.Poem, collection string) ViewerModel {
	w, h, _ := term.GetSize(os.Stdout.Fd())
	if w == 0 {
		w = 80
	}
	if h == 0 {
		h = 24
	}
	return ViewerModel{
		poems:      poems,
		index:      0,
		width:      w,
		height:     h,
		collection: collection,
	}
}

func (m ViewerModel) Init() tea.Cmd {
	return nil
}

func (m ViewerModel) Update(msg tea.Msg) (ViewerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "right", "l", " ", "n":
			m.index = (m.index + 1) % len(m.poems)
		case "left", "p":
			m.index = (m.index - 1 + len(m.poems)) % len(m.poems)
		case "r":
			m.index = rand.Intn(len(m.poems))
		case "h", "?":
			m.showHelp = !m.showHelp
		}
	}
	return m, nil
}

func (m ViewerModel) View() string {
	width := m.width
	height := m.height

	// Window too small
	if width < minWidth || height < minHeight {
		center := lipgloss.NewStyle().Width(width).Align(lipgloss.Center)
		msg := center.Render(tooSmallStyle.Render("窗口太小"))
		topPad := (height - 1) / 2
		var sb strings.Builder
		for i := 0; i < topPad; i++ {
			sb.WriteByte('\n')
		}
		sb.WriteString(msg)
		return sb.String()
	}

	poem := m.poems[m.index]
	center := lipgloss.NewStyle().Width(width).Align(lipgloss.Center)

	// Title
	titleLine := center.Render(titleStyle.Render(poem.Title))

	// Divider — at most one terminal line wide, with side margins
	maxDiv := width - 8
	if maxDiv < 6 {
		maxDiv = 6
	}
	divLen := longestLineWidth(poem.Paragraphs)
	if divLen < 6 {
		divLen = 6
	}
	if divLen > maxDiv {
		divLen = maxDiv
	}
	divider := center.Render(dividerStyle.Render(strings.Repeat("─", divLen)))

	// Author + dynasty
	meta := fmt.Sprintf("%s  %s", poem.Author, poem.Dynasty)
	metaLine := center.Render(metaStyle.Render(meta))

	// Poem body — each paragraph on its own line
	var bodyLines []string
	for _, p := range poem.Paragraphs {
		if p == "" {
			bodyLines = append(bodyLines, "")
			continue
		}
		bodyLines = append(bodyLines, center.Render(bodyStyle.Render(p)))
	}
	body := strings.Join(bodyLines, "\n")

	content := titleLine + "\n" + divider + "\n" + metaLine + "\n\n" + body

	// Vertical centering
	contentHeight := strings.Count(content, "\n") + 1
	topPad := (height - contentHeight) / 2
	if topPad < 0 {
		topPad = 0
	}

	var sb strings.Builder
	for i := 0; i < topPad; i++ {
		sb.WriteByte('\n')
	}
	sb.WriteString(content)

	// Help bar (shown only when h/? is pressed)
	if m.showHelp {
		hint := fmt.Sprintf("← 上首   → 下首   r 随机   Esc 重选集合   h 关闭帮助   q 退出   %d / %d",
			m.index+1, len(m.poems))
		// fill rest of height
		usedLines := topPad + contentHeight
		remaining := height - usedLines
		if remaining < 0 {
			remaining = 0
		}
		for i := 0; i < remaining; i++ {
			sb.WriteByte('\n')
		}
		sb.WriteString(hintStyle.Render(hint))
	}

	return sb.String()
}

// longestLineWidth returns the display width of the widest paragraph line.
func longestLineWidth(paragraphs []string) int {
	max := 0
	for _, p := range paragraphs {
		w := lipgloss.Width(p)
		if w > max {
			max = w
		}
	}
	return max
}
