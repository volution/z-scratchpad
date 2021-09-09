

package zscratchpad


import "fmt"




func DocumentRenderToSource (_document *Document) (string, *Error) {
	
	_buffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_buffer)
	
	if _document.TitleOriginal != "" {
		fmt.Fprintf (_buffer, "## -- title:       %s\n", _document.TitleOriginal)
	}
	for _, _title := range _document.TitleOriginalAlternatives {
		if _title == _document.TitleOriginal {
			continue
		}
		fmt.Fprintf (_buffer, "## -- title:       %s\n", _title)
	}
	
	if _document.Library != "" {
		fmt.Fprintf (_buffer, "## -- library:     %s\n", _document.Library)
	}
	if _document.Identifier != "" {
		fmt.Fprintf (_buffer, "## -- identifier:  %s\n", _document.Identifier)
	}
	if _document.Format != "" {
		fmt.Fprintf (_buffer, "## -- format:      %s\n", _document.Format)
	}
	if ! _document.Timestamp.IsZero () {
		fmt.Fprintf (_buffer, "## -- timestamp:   %s\n", _document.Timestamp.Format ("2006-01-02 15:04:05"))
	}
	
	if !_document.BodyEmpty {
		_buffer.WriteByte ('\n')
		_buffer.WriteByte ('\n')
		for _, _line := range _document.BodyLines {
			_buffer.WriteString (_line)
			_buffer.WriteByte ('\n')
		}
	}
	
	_source := _buffer.String ()
	
	return _source, nil
}

