

package zscratchpad


import "bytes"
import "html"




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
			_render, _error = documentRenderTextToHtml (_document.Body)
		
		case "snippets" :
			_render, _error = documentRenderSnippetsToHtml (_document.Body)
		
		case "commonmark" :
			_render, _error = documentRenderCommonMarkToHtml (_document.Body)
		
		default :
			return "", errorf (0xaf60ea6d, "format invalid `%s`", _document.Format)
	}
	
	if _error != nil {
		return "", _error
	}
	
	_document.RenderHtml = _render
	
	return _document.RenderHtml, nil
}




func documentRenderCommonMarkToHtml (_source string) (string, *Error) {
	return parseAndRenderCommonMarkToHtml (_source)
}

func documentRenderSnippetsToHtml (_source string) (string, *Error) {
	return parseAndRenderSnippetsToHtml (_source)
}




func documentRenderTextToHtml (_source string) (string, *Error) {
	
	_lines, _ := stringSplitLines (_source)
	
	_buffer := bytes.NewBuffer (nil)
	
	_buffer.WriteString ("<pre>\n")
	for _, _line := range _lines {
		_line = html.EscapeString (_line)
		_buffer.WriteString (_line)
		_buffer.WriteString ("\n")
	}
	_buffer.WriteString ("</pre>\n")
	
	_output := string (_buffer.Bytes ())
	
	return _output, nil
}

