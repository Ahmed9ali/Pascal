package model

import (
	"fmt"
	"math"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
)

// shimmerText renders s in grey with a soft light band sweeping left to
// right. Position and brightness are continuous, so the motion is smooth.
func shimmerText(s string, now time.Time) string {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}
	const band = 5.0
	const msPerChar = 75.0
	cycle := float64(n) + 2*band
	ms := float64(now.UnixNano()/int64(time.Millisecond)) / msPerChar
	pos := math.Mod(ms, cycle) - band
	var b strings.Builder
	for i, r := range runes {
		t := 0.0
		if d := math.Abs(float64(i) - pos); d < band {
			t = 0.5 * (1 + math.Cos(math.Pi*d/band))
		}
		c := 0x88 + int(t*float64(0xEE-0x88))
		col := lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", c, c, c))
		b.WriteString(lipgloss.NewStyle().Foreground(col).Render(string(r)))
	}
	return b.String()
}
