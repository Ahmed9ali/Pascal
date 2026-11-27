package model

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// shimmerFrame is the redraw interval of the "Processing" shimmer (~30 fps).
const shimmerFrame = 33 * time.Millisecond

// shimmerTickMsg is one frame of the shimmer's own clock.
type shimmerTickMsg struct{}

// shimmerWanted reports whether the busy placeholder is on screen.
func (m *UI) shimmerWanted() bool {
	return m.isAgentBusy() && m.textarea.Value() == ""
}

// ensureShimmer arms the shimmer clock when it is wanted and not running.
func (m *UI) ensureShimmer() tea.Cmd {
	if !m.shimmerWanted() {
		m.shimmerRunning = false
		return nil
	}
	if m.shimmerRunning {
		return nil
	}
	m.shimmerRunning = true
	return tea.Tick(shimmerFrame, func(time.Time) tea.Msg { return shimmerTickMsg{} })
}

// shimmerText renders s in grey with a soft light band sweeping left to
// right across the word. Trailing dots stay a steady grey, outside the sweep.
func shimmerText(s string, now time.Time) string {
	head := strings.TrimRight(s, ".")
	dots := s[len(head):]
	runes := []rune(head)
	n := len(runes)
	if n == 0 {
		return s
	}
	const band = 3.0
	const msPerChar = 70.0
	cycle := float64(n) + 2*band
	ms := float64(now.UnixNano()/int64(time.Millisecond)) / msPerChar
	pos := math.Mod(ms, cycle) - band
	var b strings.Builder
	for i, r := range runes {
		t := 0.0
		if d := math.Abs(float64(i) - pos); d < band {
			t = 0.5 * (1 + math.Cos(math.Pi*d/band))
		}
		c := 0x80 + int(t*float64(0xF0-0x80))
		col := lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", c, c, c))
		b.WriteString(lipgloss.NewStyle().Foreground(col).Render(string(r)))
	}
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Render(dots))
	return b.String()
}
