package model

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/lipgloss/v2"
)

// lineNumberGutter renders the dim ">" used on lines the cursor is not on.
func lineNumberGutter(_ textarea.PromptInfo, base lipgloss.Style) string {
	return lipgloss.NewStyle().Foreground(base.GetForeground()).Render("  > ")
}

// cursorLinePrompt wraps a prompt function so the line holding the cursor
// gets the active (green) prompt and every other line gets the dim one.
func (m *UI) cursorLinePrompt(inner func(textarea.PromptInfo) string) func(textarea.PromptInfo) string {
	return func(info textarea.PromptInfo) string {
		if info.LineNumber == m.textarea.Line() {
			return inner(textarea.PromptInfo{LineNumber: 0, Focused: info.Focused})
		}
		return inner(textarea.PromptInfo{LineNumber: 1, Focused: info.Focused})
	}
}
