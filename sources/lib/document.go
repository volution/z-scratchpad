

package zscratchpad


import "bytes"
import "fmt"
import "io"
import "os"
import "strings"
import "unicode/utf8"




type Document struct {
	
	Identifier string
	
	Library string
	PathInLibrary string
	PathOriginal string
	
	Title string
	TitleAlternatives []string
	
	SourceFingerprint string
	
	Format string
	
	Body string
	BodyFingerprint string
	
	RenderHtml string
	RenderText string
}




func DocumentResolveIdentifier (_document *Document, _perhapsUseFileName bool) (*Error) {
	
	if _document.Identifier != "" {
		return nil
	}
	
	if (_document.Library != "") && (_document.PathInLibrary != "") && _perhapsUseFileName {
		if _name, _, _error := pathSplitFileNameAndExtension (_document.PathOriginal); _error == nil {
			_document.Identifier = _document.Library + "/" + _name
			return nil
		} else {
			return _error
		}
	}
	
	if (_document.PathOriginal != "") && _perhapsUseFileName {
		if _name, _, _error := pathSplitFileNameAndExtension (_document.PathOriginal); _error == nil {
			_document.Identifier = _name
			return nil
		} else {
			return _error
		}
	}
	
	if (_document.Library != "") && (_document.PathInLibrary != "") {
		_fingerprint := fingerprintString ("/:/" + _document.Library + "/:/" + _document.PathInLibrary)
		_document.Identifier = _fingerprint[:32]
		return nil
	}
	
	if _document.PathOriginal != "" {
		_fingerprint := fingerprintString ("/_/" + _document.PathOriginal)
		_document.Identifier = _fingerprint[:32]
		return nil
	}
	
	return errorf (0x1c58da80, "identifier unresolvable")
}




func DocumentResolveFormat (_document *Document, _perhapsUseFileExtension bool) (*Error) {
	
	if _document.Format != "" {
		return nil
	}
	
	if (_document.PathOriginal != "") && _perhapsUseFileExtension {
		if _, _extension, _error := pathSplitFileNameAndExtension (_document.PathOriginal); _error == nil {
			_format := ""
			switch _extension {
				case "md" :
					_format = "commonmark"
				case "txt" :
					_format = "text"
			}
			if _format != "" {
				_document.Format = _format
				return nil
			}
		} else {
			return _error
		}
	}
	
	return errorf (0xe5e1dd0f, "format unresolvable")
}




func DocumentLoadFromPath (_path string) (*Document, *Error) {
	
	var _sourceBytes []byte
	if _bytes, _error := os.ReadFile (_path); _error == nil {
		_sourceBytes = _bytes
	} else {
		return nil, errorw (0x483c6b27, _error)
	}
	
	if ! utf8.Valid (_sourceBytes) {
		return nil, errorf (0xa24965ce, "invalid UTF-8 source")
	}
	_source := string (_sourceBytes)
	
	var _document *Document
	if _document_0, _error := DocumentLoadFromBuffer (_source); _error == nil {
		_document = _document_0
	} else {
		return nil, _error
	}
	
	_document.PathOriginal = _path
	
	return _document, nil
}



func DocumentLoadFromBuffer (_source string) (*Document, *Error) {
	
	var _identifier string
	var _format string
	var _title string
	var _titles []string
	
	_body := _source
	for {
		
		if _body == "" {
			break
		}
		_header, _rest, _ok := stringSplitLine (_body)
		if !_ok {
			break
		}
		
		if _header == "" {
			_body = _rest
			break
		}
		
		if ! strings.HasPrefix (_header, "## ") {
			break
		}
		
		_body = _rest
		
		_header = _header[3:]
		_header = stringTrimSpaces (_header)
		
		if _header == "" {
			return nil, errorf (0x8d4a068d, "header empty")
		}
		
		if strings.HasPrefix (_header, "-- ") {
			
			_header = _header[3:]
			_header = stringTrimSpaces (_header)
			
			if strings.HasPrefix (_header, "identifier:") {
				_identifier_0 := _header[11:]
				_identifier_0 = stringTrimSpaces (_identifier_0)
				_identifier = _identifier_0
			} else if strings.HasPrefix (_header, "format:") {
				_format_0 := _header[7:]
				_format_0 = stringTrimSpaces (_format_0)
				_format = _format_0
			} else {
				return nil, errorf (0xc5ccdc9e, "metadata invalid `%s`", _header)
			}
			
		} else {
			
			_titles = append (_titles, _header)
			if _title == "" {
				_title = _header
			}
		}
	}
	
	_sourceFingerprint := fingerprintString (_source)
	_bodyFingerprint := fingerprintString (_body)
	
	_document := & Document {
			Title : _title,
			TitleAlternatives : _titles,
			Identifier : _identifier,
			Format : _format,
			SourceFingerprint : _sourceFingerprint,
			Body : _body,
			BodyFingerprint : _bodyFingerprint,
		}
	
	return _document, nil
}




func DocumentDump (_stream io.Writer, _document *Document, _includeIdentifiers bool, _includeBody bool, _includeRender bool) (*Error) {
	
	_buffer := bytes.NewBuffer (nil)
	
	if _document.Title != "" {
		fmt.Fprintf (_buffer, "-- title (primary): `%s`\n", _document.Title)
	}
	for _, _title := range _document.TitleAlternatives {
		if _title == _document.Title {
			continue
		}
		fmt.Fprintf (_buffer, "-- title (alternative): `%s`\n", _title)
	}
	
	if _document.Library != "" {
		fmt.Fprintf (_buffer, "-- library: `%s`\n", _document.Library)
	}
	if _document.PathInLibrary != "" {
		fmt.Fprintf (_buffer, "-- path in library: `%s`\n", _document.PathInLibrary)
	}
	if _document.PathOriginal != "" {
		fmt.Fprintf (_buffer, "-- path in file-system: `%s`\n", _document.PathOriginal)
	}
	
	if _document.Format != "" {
		fmt.Fprintf (_buffer, "-- format: `%s`\n", _document.Format)
	}
	
	if _includeIdentifiers {
		if _document.Identifier != "" {
			fmt.Fprintf (_buffer, "-- identifier: `%s`\n", _document.Identifier)
		}
		fmt.Fprintf (_buffer, "-- source fingerprint: `%s`\n", _document.SourceFingerprint)
		fmt.Fprintf (_buffer, "-- body fingerprint: `%s`\n", _document.BodyFingerprint)
	}
	
	if _document.Body == "" {
		fmt.Fprintf (_buffer, "-- body: empty\n")
	} else if _includeBody {
		fmt.Fprintf (_buffer, "-- body:\n")
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
		_lines, _ := stringSplitLines (_document.Body)
		for _, _line := range _lines {
			fmt.Fprintf (_buffer, "    %s\n", _line)
		}
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
	}
	
	if _includeRender && (_document.RenderText != "") {
		fmt.Fprintf (_buffer, "-- render text:\n")
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
		_lines, _ := stringSplitLines (_document.RenderText)
		for _, _line := range _lines {
			fmt.Fprintf (_buffer, "    %s\n", _line)
		}
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
	}
	
	if _includeRender && (_document.RenderHtml != "") {
		fmt.Fprintf (_buffer, "-- render HTML:\n")
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
		_lines, _ := stringSplitLines (_document.RenderHtml)
		for _, _line := range _lines {
			fmt.Fprintf (_buffer, "    %s\n", _line)
		}
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
	}
	
	if _, _error := _buffer.WriteTo (_stream); _error != nil {
		return errorw (0x9fe2e5af, _error)
	}
	
	return nil
}

