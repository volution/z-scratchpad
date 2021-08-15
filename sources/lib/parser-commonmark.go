

package zscratchpad


import "bytes"
import "unicode/utf8"
import "strings"

import goldmark "github.com/yuin/goldmark"
import goldmark_text "github.com/yuin/goldmark/text"
import goldmark_html "github.com/yuin/goldmark/renderer/html"




func parseAndRenderCommonMarkToHtml (_sourceLines []string) (string, *Error) {
	
	_sourceBuffer := bytes.NewBuffer (nil)
	for _, _line := range _sourceLines {
		_sourceBuffer.WriteString (_line)
		_sourceBuffer.WriteByte ('\n')
	}
	_sourceBytes := _sourceBuffer.Bytes ()
	
	_parser := goldmark.DefaultParser ()
	
	_renderer := goldmark.DefaultRenderer ()
	_renderer.AddOptions (
			goldmark_html.WithXHTML (),
			goldmark_html.WithUnsafe (),
		)
	
	_reader := goldmark_text.NewReader (_sourceBytes)
	_writer := bytes.NewBuffer (nil)
	
	_ast := _parser.Parse (_reader)
	if _error := _renderer.Render (_writer, _sourceBytes, _ast); _error != nil {
		return "", errorw (0xfc82f523, _error)
	}
	
	_outputBytes := _writer.Bytes ()
	if ! utf8.Valid (_outputBytes) {
		return "", errorw (0xbc65423b, nil)
	}
	
	_output := string (_outputBytes)
	
	// FIXME:  This is generated for custom HTML!
	_output = strings.ReplaceAll (_output, "<p><!-- raw HTML omitted --><!-- raw HTML omitted --></p>\n", "")
	
	return _output, nil
}

