package install

import (
	"fmt"
	"strings"
)

// The markers around forge's block. Docket's kit has its own, which forge
// never touches.
const (
	blockStart = "<!-- forge:start -->"
	blockEnd   = "<!-- forge:end -->"
)

// findBlock finds forge's block in text: from the start of its first marker's
// line to the end of its last marker's line, newline included. It reports
// false when there's none, and an error when the markers don't pair up.
func findBlock(file, text string) (start, end int, found bool, err error) {
	starts, ends := 0, 0
	start, end = -1, -1
	offset := 0
	for line := range strings.SplitAfterSeq(text, "\n") {
		switch strings.TrimRight(line, "\r\n") {
		case blockStart:
			starts++
			if start < 0 {
				start = offset
			}
		case blockEnd:
			ends++
			if end < 0 {
				end = offset + len(line)
			}
		}
		offset += len(line)
	}
	switch {
	case starts == 0 && ends == 0:
		return 0, 0, false, nil
	case starts != 1 || ends != 1 || end < start:
		return 0, 0, false, fmt.Errorf("%s: forge's markers are unbalanced (%d start, %d end); fix them by hand", file, starts, ends)
	}
	return start, end, true, nil
}

// wrap puts a block between forge's markers.
func wrap(block string) string {
	return blockStart + "\n" + strings.TrimRight(block, "\n") + "\n" + blockEnd + "\n"
}

// setBlock puts block in text: in place of forge's block, or at the end after
// a blank line.
func setBlock(text, block string) string {
	start, end, found, _ := findBlock("", text)
	if found {
		return text[:start] + wrap(block) + text[end:]
	}
	if text == "" {
		return wrap(block)
	}
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	return text + "\n" + wrap(block)
}

// dropBlock takes forge's block out of text, with the blank line setBlock put
// before it.
func dropBlock(text string) string {
	start, end, found, _ := findBlock("", text)
	if !found {
		return text
	}
	before := strings.TrimRight(text[:start], "\n")
	after := strings.TrimLeft(text[end:], "\n")
	switch {
	case before == "":
		return after
	case after == "":
		return before + "\n"
	}
	return before + "\n\n" + after
}

// hasLine reports whether text has line on a line of its own.
func hasLine(text, line string) bool {
	for l := range strings.SplitSeq(text, "\n") {
		if strings.TrimSpace(l) == line {
			return true
		}
	}
	return false
}

// addLine adds line at the end of text, after a blank line.
func addLine(text, line string) string {
	if text == "" {
		return line + "\n"
	}
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	return text + "\n" + line + "\n"
}

// dropLine takes line out of text, with the blank line addLine put before it.
func dropLine(text, line string) string {
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		if strings.TrimSpace(l) != line {
			continue
		}
		lines = append(lines[:i], lines[i+1:]...)
		if i > 0 && strings.TrimSpace(lines[i-1]) == "" && (i == len(lines) || strings.TrimSpace(lines[i]) == "") {
			lines = append(lines[:i-1], lines[i:]...)
		}
		break
	}
	return strings.Join(lines, "\n")
}
