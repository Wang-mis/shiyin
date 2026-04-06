package ui

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"runtime"
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
			Foreground(lipgloss.Color("#3A3A3A"))

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
		case "o":
			poem := m.poems[m.index]
			u := "https://www.guwendao.net/search.aspx?value=" + url.QueryEscape(poem.Title)
			openInBrowser(u)
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
		return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center,
			tooSmallStyle.Render("窗口太小"))
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

	// Fill remaining lines so Bubble Tea clears any leftover content on resize/toggle
	usedLines := topPad + contentHeight
	remaining := height - usedLines - 1
	if remaining < 0 {
		remaining = 0
	}

	if m.showHelp {
		keys := "← →  翻页    r  随机    o  详情    Esc  返回    q  退出"
		page := fmt.Sprintf("%d / %d", m.index+1, len(m.poems))
		pageW := lipgloss.Width(page)
		keysW := lipgloss.Width(keys)
		gap := width - keysW - pageW
		if gap < 2 {
			gap = 2
		}
		hint := keys + strings.Repeat(" ", gap) + page
		if lipgloss.Width(hint) > width {
			hint = truncateToWidth(hint, width)
		}
		for i := 0; i < remaining; i++ {
			sb.WriteByte('\n')
		}
		sb.WriteString("\n" + hintStyle.Render(hint))
	} else {
		for i := 0; i < remaining; i++ {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

// truncateToWidth truncates s so its display width does not exceed maxW.
func openInBrowser(u string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", u)
	case "darwin":
		cmd = exec.Command("open", u)
	default:
		cmd = exec.Command("xdg-open", u)
	}
	_ = cmd.Start()
}

// truncateToWidth truncates s so its display width does not exceed maxW.
func truncateToWidth(s string, maxW int) string {
	w := 0
	for i, r := range s {
		rw := lipgloss.Width(string(r))
		if w+rw > maxW {
			return s[:i]
		}
		w += rw
	}
	return s
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
