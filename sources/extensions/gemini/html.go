
package gemini


import "bytes"
import "fmt"
import "html"




type HtmlState struct {
	buffer *bytes.Buffer
	state string
}


func NewHtmlState (_buffer *bytes.Buffer) (*HtmlState) {
	return & HtmlState { buffer : _buffer }
}


func (_state *HtmlState) Flush () {
	switch _state.state {
		case "blockquote" :
			fmt.Fprintf (_state.buffer, "</blockquote>\n")
		case "list" :
			fmt.Fprintf (_state.buffer, "</ul>\n")
		case "links" :
			fmt.Fprintf (_state.buffer, "</ul>\n")
		case "code" :
			fmt.Fprintf (_state.buffer, "</code></pre>\n")
	}
	_state.state = ""
}


func (_state *HtmlState) push (_new string) () {
	if _state.state == _new {
		return
	}
	_state.Flush ()
	switch _new {
		case "blockquote" :
			fmt.Fprintf (_state.buffer, "<blockquote>\n")
		case "list" :
			fmt.Fprintf (_state.buffer, "<ul>\n")
		case "links" :
			fmt.Fprintf (_state.buffer, "<ul class=\"links\">\n")
		case "code" :
			fmt.Fprintf (_state.buffer, "<pre><code>")
	}
	_state.state = _new
}




func (l LineLink) Html(_state *HtmlState) {
	_state.push ("links")
	if l.Name != "" {
		fmt.Fprintf(_state.buffer, "<li><a href=\"%s\">%s</a></li>\n", html.EscapeString(l.URL), html.EscapeString(l.Name))
	} else {
		fmt.Fprintf(_state.buffer, "<li><a href=\"%s\">%s</a></li>\n", html.EscapeString(l.URL), html.EscapeString(l.URL))
	}
}

func (l LinePreformattingToggle) Html(_state *HtmlState) {
	_state.Flush ()
	_state.push ("code")
}

func (l LinePreformattedText) Html(_state *HtmlState) {
	_state.push ("code")
	fmt.Fprintf(_state.buffer, "%s\n", html.EscapeString(string(l)))
}

func (l LineHeading1) Html(_state *HtmlState) {
	_state.push ("header")
	fmt.Fprintf(_state.buffer, "<h1>%s</h1>\n", html.EscapeString(string(l)))
}
func (l LineHeading2) Html(_state *HtmlState) {
	_state.push ("header")
	fmt.Fprintf(_state.buffer, "<h2>%s</h2>\n", html.EscapeString(string(l)))
}
func (l LineHeading3) Html(_state *HtmlState) {
	_state.push ("header")
	fmt.Fprintf(_state.buffer, "<h3>%s</h3>\n", html.EscapeString(string(l)))
}

func (l LineListItem) Html(_state *HtmlState) {
	_state.push ("list")
	fmt.Fprintf(_state.buffer, "<li>%s</li>\n", html.EscapeString(string(l)))
}

func (l LineQuote) Html(_state *HtmlState) {
	_state.push ("blockquote")
	fmt.Fprintf(_state.buffer, "<p>%s</p>\n", html.EscapeString(string(l)))
}

func (l LineText) Html(_state *HtmlState) {
	_state.push ("paragraph")
	fmt.Fprintf(_state.buffer, "<p>%s</p>\n", html.EscapeString(string(l)))
}

