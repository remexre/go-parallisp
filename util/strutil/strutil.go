package strutil

import (
	"bytes"
	"strings"
)

// MakeHeader appends a header composed of the # and the given string to the
// given buffer.
func MakeHeader(buf *bytes.Buffer, text string) {
	buf.WriteString(strings.Repeat("#", 80))
	buf.WriteRune('\n')
	if len(text) < 76 {
		numSpacesTotal := 76 - len(text)
		numSpacesLeft := numSpacesTotal / 2
		numSpacesRight := numSpacesTotal - numSpacesLeft

		buf.WriteString("# ")
		buf.WriteString(strings.Repeat(" ", numSpacesLeft))
		buf.WriteString(text)
		buf.WriteString(strings.Repeat(" ", numSpacesRight))
		buf.WriteString(" #")
	} else {
		buf.WriteString(text)
	}
	buf.WriteRune('\n')
	buf.WriteString(strings.Repeat("#", 80))
	buf.WriteRune('\n')
}

// Comment sanitizes a string before placing it into a comment.
func Comment(str string) string {
	return strings.Replace(str, "\n", "\n# ", -1)
}

// Indent indents a string.
func Indent(str string) string {
	oldLines := strings.Split(str, "\n")
	lines := make([]string, len(oldLines))
	for i, line := range oldLines {
		lines[i] = "\t" + line
	}
	return strings.Join(lines, "\n")
}
