package goalexa

import (
	"fmt"
	"strings"
)

func EscapeSSMLText(text string) string {
	text = strings.ReplaceAll(text, "&", "&amp;")
	text = strings.ReplaceAll(text, "\"", "&quot;")
	text = strings.ReplaceAll(text, "'", "&apos;")
	text = strings.ReplaceAll(text, "<", "&lt;")
	text = strings.ReplaceAll(text, ">", "&gt;")
	text = strings.ReplaceAll(text, "\\", "\\\\")
	for _, r := range text {
		if r > 127 {
			text = strings.ReplaceAll(text, string(r), fmt.Sprintf("&#%d;", r))
		}
	}
	return text
}