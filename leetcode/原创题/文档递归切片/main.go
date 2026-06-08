package main

import (
	"fmt"
	"strings"
)

// RecursiveSplit splits `content` into chunks no larger than chunkSize (rune count).
// It tries separators by priority: paragraph ("\n\n"), newline ("\n"), period ("."),
// question ("?"), exclamation ("!"), and finally space (" "). If a segment is still
// too large it falls back to character-based chunking. Overlap controls how many
// runes from the end of each chunk should be prepended to the next chunk.
func RecursiveSplit(content string, chunkSize, overlap int) []string {
	if chunkSize <= 0 {
		return nil
	}
	if overlap < 0 {
		overlap = 0
	}

	content = strings.TrimSpace(content)
	if content == "" {
		return nil
	}

	// ensure overlap smaller than chunkSize
	if overlap >= chunkSize {
		overlap = chunkSize - 1
		if overlap < 0 {
			overlap = 0
		}
	}

	separators := []string{"\n\n", "\n", ".", "?", "!", " "}

	// First produce atomic units by recursively splitting by separators
	var units []string
	var splitUnit func(text string, sepIndex int)
	splitUnit = func(text string, sepIndex int) {
		text = strings.TrimSpace(text)
		if text == "" {
			return
		}
		if runeLen(text) <= chunkSize {
			units = append(units, text)
			return
		}
		if sepIndex >= len(separators) {
			// fallback: raw character splitting
			runes := []rune(text)
			for i := 0; i < len(runes); i += chunkSize {
				end := i + chunkSize
				if end > len(runes) {
					end = len(runes)
				}
				units = append(units, string(runes[i:end]))
			}
			return
		}

		sep := separators[sepIndex]
		parts := strings.Split(text, sep)
		if len(parts) == 1 {
			// try next separator
			splitUnit(text, sepIndex+1)
			return
		}
		for i, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			// If adding back the separator is meaningful (e.g., sentences), reappend it
			if sep == "." || sep == "?" || sep == "!" {
				// preserve punctuation removed by Split
				p = p + sep
			} else if sep == "\n" || sep == "\n\n" {
				// keep structural separation as a single newline
				if i != len(parts)-1 {
					p = p + "\n"
				}
			}

			if runeLen(p) <= chunkSize {
				units = append(units, p)
			} else {
				// needs further splitting with next separator
				splitUnit(p, sepIndex+1)
			}
		}
	}

	splitUnit(content, 0)

	// Now join units into chunks up to chunkSize, applying overlap between chunks.
	var chunks []string
	var current []rune
	for _, u := range units {
		ur := []rune(strings.TrimSpace(u))
		if len(current) == 0 {
			current = append(current, ur...)
			continue
		}
		// plus one rune for a separating space if needed
		needSep := true
		if len(current) > 0 && (isPunctRune(current[len(current)-1]) || isPunctRune(ur[0]) || current[len(current)-1] == '\n') {
			needSep = false
		}

		sepLen := 0
		if needSep {
			sepLen = 1
		}

		if len(current)+sepLen+len(ur) <= chunkSize {
			if needSep {
				current = append(current, ' ')
			}
			current = append(current, ur...)
			continue
		}

		// finalize current chunk
		chunkStr := strings.TrimSpace(string(current))
		chunks = append(chunks, chunkStr)

		// prepare next chunk: include overlap runes from end of chunk
		ov := overlap
		if ov > runeLen(chunkStr) {
			ov = runeLen(chunkStr)
		}
		prefix := lastNRunes(chunkStr, ov)
		// start new current with prefix (but ensure doesn't exceed chunkSize)
		current = []rune(prefix)
		// if adding ur would exceed chunkSize, try to fit a trimmed prefix
		if len(current)+1+len(ur) > chunkSize {
			// reduce prefix to fit
			maxPrefix := chunkSize - 1 - len(ur)
			if maxPrefix < 0 {
				maxPrefix = 0
			}
			current = []rune(firstNRunes(prefix, maxPrefix))
		}
		if needSep && len(current) > 0 {
			current = append(current, ' ')
		}
		current = append(current, ur...)
	}
	if len(current) > 0 {
		chunks = append(chunks, strings.TrimSpace(string(current)))
	}

	return chunks
}

func runeLen(s string) int { return len([]rune(s)) }

func lastNRunes(s string, n int) string {
	r := []rune(s)
	if n <= 0 || len(r) == 0 {
		return ""
	}
	if n >= len(r) {
		return string(r)
	}
	return string(r[len(r)-n:])
}

func firstNRunes(s string, n int) string {
	r := []rune(s)
	if n <= 0 || len(r) == 0 {
		return ""
	}
	if n >= len(r) {
		return string(r)
	}
	return string(r[:n])
}

func isPunctRune(r rune) bool {
	return r == '.' || r == '?' || r == '!' || r == ',' || r == ';' || r == ':'
}

func main() {
	text := `第一段：这是一个示例段落，用于展示段落切分。

第二段：下一行是一个很长的句子，需要被进一步拆分。这个句子可能很长，很长，需要分成多个 chunk。第三段：短句。`

	chunks := RecursiveSplit(text, 40, 10)
	fmt.Printf("Got %d chunks:\n", len(chunks))
	for i, c := range chunks {
		fmt.Printf("--- chunk %d (%d runes) ---\n%v\n", i+1, runeLen(c), c)
	}
}
