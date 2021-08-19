

package zscratchpad


import "unicode/utf8"
import "strings"

import goldmark "github.com/yuin/goldmark"
import goldmark_extensions "github.com/yuin/goldmark/extension"
import goldmark_parser "github.com/yuin/goldmark/parser"
import goldmark_html "github.com/yuin/goldmark/renderer/html"
import goldmark_text "github.com/yuin/goldmark/text"




func parseAndRenderCommonMarkToHtml (_sourceLines []string) (string, *Error) {
	
	_sourceBuffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_sourceBuffer)
	for _, _line := range _sourceLines {
		_sourceBuffer.WriteString (_line)
		_sourceBuffer.WriteByte ('\n')
	}
	_sourceBytes := _sourceBuffer.Bytes ()
	
	_parser := goldmark.DefaultParser ()
	_parser.AddOptions (goldmark_parser.WithAutoHeadingID ())
	
	_renderer := goldmark.DefaultRenderer ()
	_renderer.AddOptions (goldmark_html.WithXHTML ())
	_renderer.AddOptions (goldmark_html.WithUnsafe ())
	
	goldmark.New (
			goldmark.WithParser (_parser),
			goldmark.WithRenderer (_renderer),
			goldmark.WithExtensions (
					// NOTE:  Part of the GFM (GitHub Flavoured Markdown).
					goldmark_extensions.Linkify,
					goldmark_extensions.Strikethrough,
					goldmark_extensions.TaskList,
					goldmark_extensions.Table,
				),
			goldmark.WithExtensions (
					// NOTE:  Other useful extensions.
					goldmark_extensions.DefinitionList,
					goldmark_extensions.Footnote,
				),
		)
	
	_reader := goldmark_text.NewReader (_sourceBytes)
	_outputBuffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_outputBuffer)
	
	_ast := _parser.Parse (_reader)
	if _error := _renderer.Render (_outputBuffer, _sourceBytes, _ast); _error != nil {
		return "", errorw (0xfc82f523, _error)
	}
	
	_outputBytes := _outputBuffer.Bytes ()
	if ! utf8.Valid (_outputBytes) {
		return "", errorw (0xbc65423b, nil)
	}
	
	_output := string (_outputBytes)
	
	// FIXME:  This is generated for custom HTML!
	_output = strings.ReplaceAll (_output, "<p><!-- raw HTML omitted --><!-- raw HTML omitted --></p>\n", "")
	
	return _output, nil
}

