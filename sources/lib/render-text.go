

package zscratchpad




func DocumentRenderToText (_document *Document) (string, *Error) {
	
	if _document.RenderText != "" {
		return _document.RenderText, nil
	}
	
	_format := _document.Format
	if _format == "" {
		_format = "text"
		// return "", errorf (0xb50eb076, "format empty")
	}
	
	_render := ""
	_error := (*Error) (nil)
	
	switch _format {
		
		case "text" :
			_render, _error = documentRenderTextToText (_document.Body)
		
		case "snippets" :
			_render, _error = documentRenderSnippetsToText (_document.Body)
		
		case "commonmark" :
			_render, _error = documentRenderCommonMarkToText (_document.Body)
		
		default :
			return "", errorf (0x215b1603, "format invalid `%s`", _document.Format)
	}
	
	if _error != nil {
		return "", _error
	}
	
	_document.RenderText = _render
	
	return _document.RenderText, nil
}




func documentRenderCommonMarkToText (_source string) (string, *Error) {
	// FIXME:  Implement rendering to plain text!
	return _source, nil
}

func documentRenderSnippetsToText (_source string) (string, *Error) {
	// FIXME:  Implement rendering to plain text!
	return _source, nil
}

func documentRenderTextToText (_source string) (string, *Error) {
	// FIXME:  Implement rendering to plain text!
	return _source, nil
}

