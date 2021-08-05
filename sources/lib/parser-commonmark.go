

package zscratchpad


import "bytes"
import "unicode/utf8"

import goldmark "github.com/yuin/goldmark"
import goldmark_text "github.com/yuin/goldmark/text"




func parseAndRenderCommonMarkToHtml (_source string) (string, *Error) {
	
	_sourceBytes := []byte (_source)
	
	_parser := goldmark.DefaultParser ()
	_renderer := goldmark.DefaultRenderer ()
	
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
	
	return _output, nil
}

