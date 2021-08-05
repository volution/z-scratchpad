

package zscratchpad


import "bytes"
import "fmt"




func DocumentRenderToSource (_document *Document) (string, *Error) {
	
	_buffer := bytes.NewBuffer (nil)
	
	if _document.Title != "" {
		fmt.Fprintf (_buffer, "## %s\n", _document.Title)
	}
	for _, _title := range _document.TitleAlternatives {
		if _title == _document.Title {
			continue
		}
		fmt.Fprintf (_buffer, "## %s\n", _title)
	}
	
	if _document.Library != "" {
		fmt.Fprintf (_buffer, "## -- library: %s\n", _document.Library)
	}
	if _document.Identifier != "" {
		fmt.Fprintf (_buffer, "## -- identifier: %s\n", _document.Identifier)
	}
	if _document.Format != "" {
		fmt.Fprintf (_buffer, "## -- format: %s\n", _document.Format)
	}
	
	if _document.Body != "" {
		fmt.Fprintf (_buffer, "\n\n")
		_lines, _ := stringSplitLines (_document.Body)
		for _, _line := range _lines {
			fmt.Fprintf (_buffer, "%s\n", _line)
		}
	}
	
	_source := _buffer.String ()
	
	return _source, nil
}

