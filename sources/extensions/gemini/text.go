
// NOTE:  Imported from <https://git.sr.ht/~adnano/go-gemini>, see `LICENSE` and `LICENSE-GO` at the following links:
//        * https://git.sr.ht/~adnano/go-gemini/commit/964c17b99f434899363c8514b3dfb9509e6a7945#text.go
//        * https://git.sr.ht/~adnano/go-gemini/tree/master/item/LICENSE
//        * https://git.sr.ht/~adnano/go-gemini/tree/master/item/LICENSE-GO


package gemini

import (
//	"bufio"
	"fmt"
//	"io"
	"strings"
)

// Line represents a line of a Gemini text response.
type Line interface {
	// String formats the line for use in a Gemini text response.
	String() string
	Html(*HtmlState)
	line() // private function to prevent other packages from implementing Line
}

// LineLink is a link line.
type LineLink struct {
	URL  string
	Name string
}

// LinePreformattingToggle is a preformatting toggle line.
type LinePreformattingToggle string

// LinePreformattedText is a preformatted text line.
type LinePreformattedText string

// LineHeading1 is a first-level heading line.
type LineHeading1 string

// LineHeading2 is a second-level heading line.
type LineHeading2 string

// LineHeading3 is a third-level heading line.
type LineHeading3 string

// LineListItem is an unordered list item line.
type LineListItem string

// LineQuote is a quote line.
type LineQuote string

// LineText is a text line.
type LineText string

func (l LineLink) String() string {
	if l.Name != "" {
		return fmt.Sprintf("=> %s %s", l.URL, l.Name)
	}
	return fmt.Sprintf("=> %s", l.URL)
}
func (l LinePreformattingToggle) String() string {
	return fmt.Sprintf("```%s", string(l))
}
func (l LinePreformattedText) String() string {
	return string(l)
}
func (l LineHeading1) String() string {
	return fmt.Sprintf("# %s", string(l))
}
func (l LineHeading2) String() string {
	return fmt.Sprintf("## %s", string(l))
}
func (l LineHeading3) String() string {
	return fmt.Sprintf("### %s", string(l))
}
func (l LineListItem) String() string {
	return fmt.Sprintf("* %s", string(l))
}
func (l LineQuote) String() string {
	return fmt.Sprintf("> %s", string(l))
}
func (l LineText) String() string {
	return string(l)
}

func (l LineLink) line()                {}
func (l LinePreformattingToggle) line() {}
func (l LinePreformattedText) line()    {}
func (l LineHeading1) line()            {}
func (l LineHeading2) line()            {}
func (l LineHeading3) line()            {}
func (l LineListItem) line()            {}
func (l LineQuote) line()               {}
func (l LineText) line()                {}

// Text represents a Gemini text response.
type Text []Line

// ParseText parses Gemini text from the provided io.Reader.
/*
func ParseText(r io.Reader) (Text, error) {
	var t Text
	err := ParseLines(r, func(line Line) {
		t = append(t, line)
	})
	return t, err
}
*/

// ParseLines parses Gemini text from the provided io.Reader.
// It calls handler with each line that it parses.
/*
func ParseLines(r io.Reader, handler func(Line)) error {
	const spacetab = " \t"
	var pre bool
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		...
	}
	return scanner.Err()
}
*/
func ParseLines(textLines []string, handler func(Line)) error {
	const spacetab = " \t"
	var pre bool
	for _, text := range textLines {
		var line Line
		if strings.HasPrefix(text, "```") {
			pre = !pre
			text = text[3:]
			text = strings.Trim(text, spacetab)
			line = LinePreformattingToggle(text)
		} else if pre {
			line = LinePreformattedText(text)
		} else if strings.HasPrefix(text, "=>") {
			text = text[2:]
			text = strings.Trim(text, spacetab)
			split := strings.IndexAny(text, spacetab)
			if split == -1 {
				// text is a URL
				line = LineLink{URL: text}
			} else {
				url := text[:split]
				name := text[split:]
				name = strings.Trim(name, spacetab)
				line = LineLink{url, name}
			}
		} else if strings.HasPrefix(text, "*") {
			text = text[1:]
			text = strings.Trim(text, spacetab)
			line = LineListItem(text)
		} else if strings.HasPrefix(text, "###") {
			text = text[3:]
			text = strings.Trim(text, spacetab)
			line = LineHeading3(text)
		} else if strings.HasPrefix(text, "##") {
			text = text[2:]
			text = strings.Trim(text, spacetab)
			line = LineHeading2(text)
		} else if strings.HasPrefix(text, "#") {
			text = text[1:]
			text = strings.Trim(text, spacetab)
			line = LineHeading1(text)
		} else if strings.HasPrefix(text, ">") {
			text = text[1:]
			text = strings.Trim(text, spacetab)
			line = LineQuote(text)
		} else {
			line = LineText(text)
		}
		if text != "" || pre {
			handler(line)
		}
	}
	return nil
}

// String writes the Gemini text response to a string and returns it.
func (t Text) String() string {
	var b strings.Builder
	for _, l := range t {
		b.WriteString(l.String())
		b.WriteByte('\n')
	}
	return b.String()
}
