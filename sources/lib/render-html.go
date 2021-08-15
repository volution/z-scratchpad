

package zscratchpad


import "bytes"
import "html"

import "github.com/microcosm-cc/bluemonday"




func DocumentRenderToHtml (_document *Document) (string, *Error) {
	
	if _document.RenderHtml != "" {
		return _document.RenderHtml, nil
	}
	
	_format := _document.Format
	if _format == "" {
		_format = "text"
		// return "", errorf (0xaff80238, "format empty")
	}
	
	_render := ""
	_error := (*Error) (nil)
	
	switch _format {
		
		case "text" :
			_render, _error = documentRenderTextToHtml (_document.BodyLines)
		
		case "snippets" :
			_render, _error = documentRenderSnippetsToHtml (_document.BodyLines)
		
		case "commonmark" :
			_render, _error = documentRenderCommonMarkToHtml (_document.BodyLines)
		
		default :
			return "", errorf (0xaf60ea6d, "format invalid `%s`", _document.Format)
	}
	
	if _error != nil {
		return "", _error
	}
	
	if true {
		
		_parser := bluemonday.UGCPolicy()
		
		_parser.RequireParseableURLs (true)
		_parser.RequireNoFollowOnLinks (true)
		_parser.RequireNoReferrerOnLinks (true)
		_parser.RequireCrossOriginAnonymous (true)
		
		_sanitized := _parser.Sanitize (_render)
		if _sanitized != _render {
			logf ('d', 0xe4eb5c90, "rendered document was sanitized!")
			_render = _sanitized
		}
	}
	
	_document.RenderHtml = _render
	
	return _document.RenderHtml, nil
}




func documentRenderCommonMarkToHtml (_source []string) (string, *Error) {
	return parseAndRenderCommonMarkToHtml (_source)
}

func documentRenderSnippetsToHtml (_source []string) (string, *Error) {
	return parseAndRenderSnippetsToHtml (_source)
}




func documentRenderTextToHtml (_source []string) (string, *Error) {
	
	_buffer := bytes.NewBuffer (nil)
	
	_buffer.WriteString ("<pre>\n")
	for _, _line := range _source {
		_line = html.EscapeString (_line)
		_buffer.WriteString (_line)
		_buffer.WriteString ("\n")
	}
	_buffer.WriteString ("</pre>\n")
	
	_output := string (_buffer.Bytes ())
	
	return _output, nil
}

