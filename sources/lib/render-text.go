

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
			_render, _error = documentRenderTextToText (_document.BodyLines)
		
		case "snippets" :
			_render, _error = documentRenderSnippetsToText (_document.BodyLines)
		
		case "commonmark" :
			_render, _error = documentRenderCommonMarkToText (_document.BodyLines)
		
		default :
			return "", errorf (0x215b1603, "format invalid `%s`", _document.Format)
	}
	
	if _error != nil {
		return "", _error
	}
	
	_document.RenderText = _render
	
	return _document.RenderText, nil
}




func documentRenderCommonMarkToText (_source []string) (string, *Error) {
	// FIXME:  Implement rendering to plain text!
	return documentRenderAnyToText (_source)
}

func documentRenderSnippetsToText (_source []string) (string, *Error) {
	// FIXME:  Implement rendering to plain text!
	return documentRenderAnyToText (_source)
}

func documentRenderTextToText (_source []string) (string, *Error) {
	// FIXME:  Implement rendering to plain text!
	return documentRenderAnyToText (_source)
}

func documentRenderAnyToText (_source []string) (string, *Error) {
	_buffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_buffer)
	for _, _line := range _source {
		_buffer.WriteString (_line)
		_buffer.WriteByte ('\n')
	}
	_render := string (_buffer.Bytes ())
	return _render, nil
}

